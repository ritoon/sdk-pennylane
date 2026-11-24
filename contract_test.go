package pennylane

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime"
	"net/http"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"testing"

	"github.com/oapi-codegen/runtime/types"
)

// Exercise the actual typed methods against the shipped OpenAPI contract. These
// synthetic bodies test HTTP wiring and decoding, not Pennylane business rules.
func TestOperationHTTPContracts(t *testing.T) {
	type payload struct{ Content map[string]json.RawMessage }
	type operation struct {
		OperationID string
		Parameters  []struct {
			Name, In string
			Required bool
		}
		RequestBody *payload
		Responses   map[string]payload
	}
	var spec struct {
		Paths      map[string]map[string]operation
		Components struct {
			Schemas map[string]struct{ Type string }
		}
	}
	data, err := os.ReadFile("internal/spec/company.json")
	if err != nil {
		t.Fatal(err)
	}
	if err := json.Unmarshal(data, &spec); err != nil {
		t.Fatal(err)
	}
	placeholder := regexp.MustCompile(`\{[^}]+\}`)
	count := 0
	for path, methods := range spec.Paths {
		for verb, op := range methods {
			count++
			t.Run(op.OperationID, func(t *testing.T) {
				parts := strings.Split(op.OperationID, "-")
				for i, p := range parts {
					parts[i] = strings.ToUpper(p[:1]) + p[1:]
				}
				name := strings.Join(parts, "") + "WithResponse"
				statuses := map[string]payload{"418": {}}
				for status, response := range op.Responses {
					statuses[status] = response
				}
				for status, response := range statuses {
					t.Run(status, func(t *testing.T) {
						code, err := strconv.Atoi(status)
						if err != nil {
							t.Fatalf("unsupported status %q", status)
						}
						body := "upstream error"
						if media, ok := response.Content["application/json"]; ok {
							body = "{}"
							var m struct {
								Schema struct {
									Ref string `json:"$ref"`
								}
							}
							if err := json.Unmarshal(media, &m); err != nil {
								t.Fatal(err)
							}
							if spec.Components.Schemas[strings.TrimPrefix(m.Schema.Ref, "#/components/schemas/")].Type == "array" {
								body = "[]"
							}
						}
						if code == 204 {
							body = ""
						}
						calls := 0
						expectedPath := path
						c := mockClient(t, code, body, func(r *http.Request) {
							calls++
							if r.Method != strings.ToUpper(verb) || r.URL.Path != expectedPath {
								t.Fatalf("got %s %s, want %s %s", r.Method, r.URL.Path, strings.ToUpper(verb), expectedPath)
							}
							company := "42"
							if strings.HasPrefix(path, "/api/external/v2/webhook_subscriptions") {
								company = ""
							}
							if r.Header.Get("X-Company-Id") != company {
								t.Fatal("incorrect company scope")
							}
							for _, p := range op.Parameters {
								if p.Required && p.In == "query" && !r.URL.Query().Has(p.Name) {
									t.Errorf("missing required query %s", p.Name)
								}
							}
							if op.RequestBody != nil {
								ct, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
								if err != nil {
									t.Fatal(err)
								}
								if _, ok := op.RequestBody.Content[ct]; !ok {
									t.Fatalf("unexpected content type %q", ct)
								}
								if ct == "application/json" {
									b, err := io.ReadAll(r.Body)
									if err != nil || !json.Valid(b) {
										t.Fatalf("invalid JSON request %q: %v", b, err)
									}
								} else if ct == "multipart/form-data" {
									if err := r.ParseMultipartForm(4096); err != nil {
										t.Fatal(err)
									}
									defer r.MultipartForm.RemoveAll()
									file, header, err := r.FormFile("file")
									if err != nil {
										t.Fatal(err)
									}
									defer file.Close()
									b, err := io.ReadAll(file)
									if err != nil || string(b) != "invoice content" || header.Filename != "invoice.pdf" {
										t.Fatal("multipart file corrupted")
									}
								}
							}
						})
						method := reflect.ValueOf(c).MethodByName(name)
						if !method.IsValid() {
							t.Fatalf("missing %s", name)
						}
						args := []reflect.Value{reflect.ValueOf(context.Background())}
						for i := 1; i < method.Type().NumIn()-1; i++ {
							typ := method.Type().In(i)
							v := reflect.New(typ).Elem()
							switch typ.Kind() {
							case reflect.Int, reflect.Int64:
								v.SetInt(73)
							case reflect.Float32, reflect.Float64:
								v.SetFloat(73)
							case reflect.Slice:
								v.Set(reflect.MakeSlice(typ, 0, 0))
							case reflect.String:
								v.SetString("sample +&")
							case reflect.Pointer:
								v.Set(reflect.New(typ.Elem()))
							case reflect.Struct:
								if f := v.FieldByName("Email"); f.IsValid() && f.Type() == reflect.TypeOf(types.Email("")) {
									f.SetString("test@example.org")
								}
								if f := v.FieldByName("File"); f.IsValid() {
									var file types.File
									file.InitFromBytes([]byte("invoice content"), "invoice.pdf")
									f.Set(reflect.ValueOf(file))
								}
							default:
								t.Fatalf("unsupported argument %s", typ)
							}
							if typ.Kind() == reflect.String || typ.Kind() == reflect.Int || typ.Kind() == reflect.Int64 || typ.Kind() == reflect.Float32 || typ.Kind() == reflect.Float64 {
								loc := placeholder.FindStringIndex(expectedPath)
								if loc == nil {
									t.Fatalf("unexpected path argument %s", typ)
								}
								expectedPath = expectedPath[:loc[0]] + fmt.Sprint(v.Interface()) + expectedPath[loc[1]:]
							}
							args = append(args, v)
						}
						result := method.Call(args)
						if !result[1].IsNil() {
							t.Fatal(result[1].Interface())
						}
						if calls != 1 {
							t.Fatalf("got %d requests", calls)
						}
						got := result[0].Interface().(interface{ StatusCode() int })
						if got.StatusCode() != code {
							t.Fatalf("status = %d, want %d", got.StatusCode(), code)
						}
						if string(result[0].Elem().FieldByName("Body").Bytes()) != body {
							t.Fatal("raw response lost")
						}
						if _, ok := response.Content["application/json"]; ok {
							field := result[0].Elem().FieldByName("JSON" + status)
							if !field.IsValid() || field.IsNil() {
								t.Fatalf("missing decoded JSON%s", status)
							}
						}
					})
				}
			})
		}
	}
	if count != 186 {
		t.Fatalf("expected 186 operations, got %d", count)
	}
}
