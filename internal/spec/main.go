// Command spec merges the endpoint definitions from an official ReadMe index
// and a directory containing the corresponding downloaded Markdown pages.
package main

import (
	"bytes"
	"encoding/json"
	"maps"
	"os"
	"path/filepath"
	"regexp"
	"slices"
	"strings"
)

type object = map[string]any

func main() {
	if len(os.Args) != 3 {
		panic("usage: go run ./internal/spec INDEX PAGES_DIRECTORY")
	}
	index, err := os.ReadFile(os.Args[1])
	check(err)
	links := regexp.MustCompile(`https://pennylane\.readme\.io/reference/[\w.-]+\.md`).FindAllString(string(index), -1)
	if len(links) == 0 {
		panic("no reference pages")
	}
	paths, schemas := object{}, object{}
	shared := map[string]string{}
	seen := map[string]bool{}
	var spec object
	for _, link := range links {
		if seen[link] {
			continue
		}
		seen[link] = true
		page, err := os.ReadFile(filepath.Join(os.Args[2], filepath.Base(link)))
		check(err)
		_, section, ok := bytes.Cut(page, []byte("# OpenAPI definition"))
		if !ok {
			switch filepath.Base(link) {
			case "supplier_invoicee_invoicing_received.md", "customer_invoicee_invoicing_status_updated.md", "dms_filecreated.md":
				continue
			}
			panic("missing OpenAPI definition: " + link)
		}
		_, block, ok := bytes.Cut(section, []byte("```json"))
		if !ok {
			panic(link)
		}
		block, _, ok = bytes.Cut(block, []byte("\n```"))
		if !ok {
			panic(link)
		}
		var doc object
		check(json.Unmarshal(block, &doc))
		if spec == nil {
			spec = object{"openapi": doc["openapi"], "info": doc["info"], "servers": doc["servers"], "components": doc["components"]}
			spec["paths"] = paths
			spec["components"].(object)["schemas"] = schemas
		}
		for path, raw := range doc["paths"].(object) {
			if paths[path] == nil {
				paths[path] = object{}
			}
			for method, value := range raw.(object) {
				op := value.(object)
				id := op["operationId"].(string)
				if paths[path].(object)[method] != nil {
					panic("duplicate " + method + " " + path)
				}
				paths[path].(object)[method] = op
				op["externalDocs"] = object{"url": strings.TrimSuffix(link, ".md")}
				responses := op["responses"].(object)
				for _, status := range slices.Sorted(maps.Keys(responses)) {
					r := responses[status]
					content, _ := r.(object)["content"].(object)
					for _, v := range content {
						media := v.(object)
						schema := media["schema"]
						key, err := json.Marshal(schema)
						check(err)
						name := shared[string(key)]
						if name == "" {
							name = strings.ToUpper(id[:1]) + id[1:] + "Status" + status
							schemas[name] = schema
							shared[string(key)] = name
						}
						media["schema"] = object{"$ref": "#/components/schemas/" + name}
					}
				}
			}
		}
	}
	normalize(spec)
	encoder := json.NewEncoder(os.Stdout)
	encoder.SetEscapeHTML(false)
	encoder.SetIndent("", "  ")
	check(encoder.Encode(spec))
}
func check(err error) {
	if err != nil {
		panic(err)
	}
}

// Inline primitive allOf members so codegen retains field types, defaults and nullability.
func normalize(value any) {
	switch node := value.(type) {
	case []any:
		for _, item := range node {
			normalize(item)
		}
	case map[string]any:
		for _, item := range node {
			normalize(item)
		}
		parts, _ := node["allOf"].([]any)
		if len(parts) == 0 {
			return
		}
		first := parts[0].(object)
		if len(parts) > 1 && first["type"] != "string" {
			return
		}
		if len(parts) == 2 && first["pattern"] != nil && parts[1].(object)["pattern"] != nil {
			return
		}
		for _, item := range parts {
			for k, v := range item.(object) {
				if _, exists := node[k]; !exists {
					node[k] = v
				}
			}
		}
		delete(node, "allOf")
	}
}
