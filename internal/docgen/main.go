// Command docgen rebuilds the Markdown reference from the shipped contract and SDK.
package main

import (
	"encoding/json"
	"fmt"
	"maps"
	"os"
	"path/filepath"
	"reflect"
	"regexp"

	"slices"
	"sort"
	"strings"

	pennylane "github.com/ritoon/sdk-pennylane"
)

type schema = map[string]any
type media struct {
	Schema   schema
	Encoding map[string]struct{ ContentType string }
}
type payload struct {
	Required    bool
	Description string
	Content     map[string]media
}
type parameter struct {
	Name, In, Description string
	Required              bool
	Example               any
	Schema                schema
}
type operation struct {
	OperationID, Summary, Description string
	Tags                              []string
	ExternalDocs                      struct{ URL string }
	Parameters                        []parameter
	RequestBody                       *payload
	Responses                         map[string]payload
}
type contract struct {
	Paths      map[string]map[string]operation
	Components struct{ Schemas map[string]schema }
}

func read(path string) []byte {
	b, e := os.ReadFile(path)
	if e != nil {
		panic(e)
	}
	return b
}
func save(path, text string) {
	if e := os.WriteFile(path, []byte(text), 0644); e != nil {
		panic(e)
	}
}
func keys[V any](m map[string]V) []string { return slices.Sorted(maps.Keys(m)) }
func literal(v any) string {
	b, e := json.Marshal(v)
	if e != nil {
		panic(e)
	}
	return "`" + string(b) + "`"
}
func cell(s string) string {
	return strings.NewReplacer("|", "&#124;", "\r", "", "\n", "<br>").Replace(strings.TrimSpace(s))
}
func presence(required bool) string {
	if required {
		return "obligatoire"
	}
	return "facultatif"
}
func format(s schema) string {
	t, _ := s["type"].(string)
	if t == "" {
		t = "non spécifié"
		for _, k := range []string{"oneOf", "anyOf", "allOf"} {
			if s[k] != nil {
				t = k
				break
			}
		}
	}
	if f, ok := s["format"].(string); ok {
		t += " (" + f + ")"
	}
	if s["nullable"] == true {
		t += " ou null"
	}
	return t
}
func details(s schema) string {
	var parts []string
	if d, ok := s["description"].(string); ok {
		parts = append(parts, d)
	}
	for _, k := range []string{"title", "enum", "default", "minimum", "maximum", "minLength", "maxLength", "pattern", "minItems", "maxItems", "uniqueItems", "deprecated", "additionalProperties", "example"} {
		if v, ok := s[k]; ok {
			parts = append(parts, k+" "+literal(v))
		}
	}
	return strings.Join(parts, "<br>")
}

const heading = "| Champ | Type | Présence | Description, contraintes et exemple |\n|---|---|---|---|\n"

func rows(b *strings.Builder, path, required string, s schema) {
	fmt.Fprintf(b, "| `%s` | %s | %s | %s |\n", cell(path), format(s), required, cell(details(s)))
	props, _ := s["properties"].(map[string]any)
	req, _ := s["required"].([]any)
	for _, key := range keys(props) {
		rows(b, path+"."+key, presence(slices.Contains(req, any(key))), props[key].(map[string]any))
	}
	if item, ok := s["items"].(map[string]any); ok {
		rows(b, path+"[]", "par élément", item)
	}
	for _, k := range []string{"oneOf", "anyOf", "allOf"} {
		branches, _ := s[k].([]any)
		for i, v := range branches {
			rows(b, fmt.Sprintf("%s (%s %d)", path, k, i+1), "selon variante", v.(map[string]any))
		}
	}
}
func fields(b *strings.Builder, s schema) { b.WriteString(heading); rows(b, "$", "racine", s) }

func requestSchema(s schema, definitions map[string]schema) schema {
	if ref, ok := s["$ref"].(string); ok {
		name, local := strings.CutPrefix(ref, "#/components/schemas/")
		resolved, found := definitions[name]
		if !local || !found {
			panic("unknown request schema: " + ref)
		}
		return resolved
	}
	return s
}

// example returns only values explicitly illustrated by the provider.
func example(s schema) (any, bool) {
	if v, ok := s["example"]; ok {
		return v, true
	}
	if child, ok := s["items"].(map[string]any); ok {
		if v, ok := example(child); ok {
			return []any{v}, true
		}
	}
	obj := map[string]any{}
	props, _ := s["properties"].(map[string]any)
	for key, child := range props {
		if v, ok := example(child.(map[string]any)); ok {
			obj[key] = v
		}
	}
	return obj, len(obj) > 0
}
func illustration(b *strings.Builder, s schema) {
	if v, ok := example(s); ok {
		data, err := json.MarshalIndent(v, "", "  ")
		if err != nil {
			panic(err)
		}
		fmt.Fprintf(b, "\n<details><summary>Extrait JSON assemblé à partir des exemples du contrat</summary>\n\n```json\n%s\n```\n\n</details>\n", data)
	}
}
func argumentNames(sig string, method reflect.Method) map[string]string {
	names := map[string]string{}
	args := strings.Split(strings.SplitN(sig, "(", 2)[1], ")")[0]
	for _, arg := range strings.Split(args, ", ") {
		name := strings.Fields(arg)[0]
		names[strings.ToLower(name)] = name
	}
	for i := 1; i < method.Type.NumIn(); i++ {
		typ := method.Type.In(i)
		if typ.Kind() != reflect.Pointer || !strings.HasSuffix(typ.Elem().Name(), "Params") {
			continue
		}
		for _, field := range reflect.VisibleFields(typ.Elem()) {
			key, _, _ := strings.Cut(field.Tag.Get("json"), ",")
			names[key] = "params." + field.Name
		}
	}
	return names
}
func main() {
	var doc contract
	if err := json.Unmarshal(read("internal/spec/company.json"), &doc); err != nil {
		panic(err)
	}
	files, err := filepath.Glob("api_*.gen.go")
	if err != nil {
		panic(err)
	}
	src := string(read("multipart.go"))
	for _, file := range files {
		src += "\n" + string(read(file))
	}
	sigs := map[string]string{}
	for _, m := range regexp.MustCompile(`(?m)^func \(c \*Client(?:WithResponses)?\) ((\w+WithResponse)\(.*?\) \(\*\w+, error\)) \{`).FindAllStringSubmatch(src, -1) {
		sigs[m[2]] = m[1]
	}
	dir := "docs/endpoints"
	if err := os.MkdirAll(dir, 0755); err != nil {
		panic(err)
	}
	groups := map[string][]string{}
	errors := map[string]schema{}
	for _, path := range keys(doc.Paths) {
		for _, verb := range keys(doc.Paths[path]) {
			op := doc.Paths[path][verb]
			parts := strings.Split(op.OperationID, "-")
			for i, p := range parts {
				parts[i] = strings.ToUpper(p[:1]) + p[1:]
			}
			name := strings.Join(parts, "") + "WithResponse"
			sig := sigs[name]
			if sig == "" {
				panic("missing method " + name)
			}
			method, ok := reflect.TypeOf((*pennylane.Client)(nil)).MethodByName(name)
			if !ok {
				panic(name)
			}
			names := argumentNames(sig, method)
			filename := op.OperationID + ".md"
			route := strings.ToUpper(verb) + " " + path
			groups[op.Tags[0]] = append(groups[op.Tags[0]], fmt.Sprintf("- [`%s`](endpoints/%s) — %s · `%s`\n", route, filename, op.Summary, name))
			var b strings.Builder
			fmt.Fprintf(&b, "# %s\n\n`%s`\n\n%s\n\n[Source Pennylane](%s) · [Index du SDK](../operations.md)\n\n## Appel Go\n\n```go\n%s\n```\n\n", op.Summary, route, strings.TrimSpace(op.Description), op.ExternalDocs.URL, sig)
			b.WriteString("## Entrée\n\n")
			if len(op.Parameters) > 0 {
				b.WriteString("| Paramètre HTTP | Argument Go | Emplacement | Type | Présence | Utilisation |\n|---|---|---|---|---|---|\n")
				for _, p := range op.Parameters {
					desc := strings.TrimSpace(p.Description)
					extra := details(p.Schema)
					if extra != "" {
						desc += "<br>" + extra
					}
					if p.Example != nil {
						desc += "<br>example " + literal(p.Example)
					}
					arg := names[p.Name]
					if p.In == "path" {
						arg = names[strings.ToLower(strings.ReplaceAll(p.Name, "_", ""))]
					}
					if arg == "" {
						panic(name + " " + p.Name)
					}
					fmt.Fprintf(&b, "| `%s` | `%s` | %s | %s | %s | %s |\n", p.Name, arg, p.In, format(p.Schema), presence(p.Required), cell(desc))
				}
				b.WriteString("\n")
			}
			if op.RequestBody == nil {
				b.WriteString("Aucun corps de requête.\n")
			} else {
				for _, ct := range keys(op.RequestBody.Content) {
					m := op.RequestBody.Content[ct]
					s := requestSchema(m.Schema, doc.Components.Schemas)
					fmt.Fprintf(&b, "**Corps %s** · `%s`\n\n", presence(op.RequestBody.Required), ct)
					fields(&b, s)
					if ct == "application/json" {
						illustration(&b, s)
					}
					for _, k := range keys(m.Encoding) {
						fmt.Fprintf(&b, "\nPartie `%s` encodée en `%s`.\n", k, m.Encoding[k].ContentType)
					}
				}
			}
			b.WriteString("\n## Sortie\n")
			for _, status := range keys(op.Responses) {
				r := op.Responses[status]
				fmt.Fprintf(&b, "\n### HTTP %s\n\n%s\n", status, r.Description)
				if len(r.Content) == 0 {
					b.WriteString("\nAucun schéma de corps documenté.\n")
					continue
				}
				for _, ct := range keys(r.Content) {
					ref := r.Content[ct].Schema["$ref"].(string)
					key := strings.TrimPrefix(ref, "#/components/schemas/")
					s, ok := doc.Components.Schemas[key]
					if !ok {
						panic(ref)
					}
					field, ok := method.Type.Out(0).Elem().FieldByName("JSON" + status)
					if !ok {
						panic(name + status)
					}
					fmt.Fprintf(&b, "\n`%s` → `réponse.JSON%s` (type `*%s`).\n\n", ct, status, field.Type.Elem().Name())
					if status[0] == '2' {
						fields(&b, s)
						illustration(&b, s)
					} else {
						errors[key] = s
						fmt.Fprintf(&b, "[Champs du payload d’erreur](erreurs.md#%s)\n", strings.ToLower(key))
					}
				}
			}
			save(filepath.Join(dir, filename), b.String())
		}
	}
	var index strings.Builder
	index.WriteString("# Référence du client Go\n\nChaque fiche décrit l’usage, la signature Go, les paramètres HTTP, le corps de requête et les réponses par statut. Les descriptions métier et les permissions reprennent la langue du contrat Pennylane.\n\nLes paramètres de chemin sont des arguments Go. Les paramètres de requête et les en-têtes se renseignent dans `params`. Pour les champs de type `nullable.Nullable[T]`, une valeur absente, une valeur nulle et une valeur renseignée sont distinctes. Dans les tableaux, `$` désigne le payload entier et `[]` un élément de liste. La présence d’un champ s’applique lorsque son objet parent est présent. Pour les variantes, `oneOf` impose un seul schéma, `anyOf` au moins un, et `allOf` tous.\n\n[Exemple Go de lecture paginée](utilisation.md) · [Payloads d’erreur](endpoints/erreurs.md)\n")
	for _, group := range keys(groups) {
		fmt.Fprintf(&index, "\n## %s\n\n", group)
		sort.Strings(groups[group])
		for _, line := range groups[group] {
			index.WriteString(line)
		}
	}
	save("docs/operations.md", index.String())
	var eb strings.Builder
	eb.WriteString("# Payloads d’erreur\n\n[Retour aux opérations](../operations.md)\n")
	for _, key := range keys(errors) {
		fmt.Fprintf(&eb, "\n## %s\n\n", key)
		fields(&eb, errors[key])
	}
	save(filepath.Join(dir, "erreurs.md"), eb.String())
}
