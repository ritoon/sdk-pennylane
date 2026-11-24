package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
)

func TestNormalize(t *testing.T) {
	for _, tc := range []struct{ name, input, want string }{
		{"primitive", `{"allOf":[{"type":"string","nullable":true,"default":"x"}]}`, `{"type":"string","nullable":true,"default":"x"}`},
		{"parent wins", `{"default":"parent","allOf":[{"type":"string","default":"child"}]}`, `{"type":"string","default":"parent"}`},
		{"nested", `{"items":[{"allOf":[{"type":"string"},{"enum":["a"]}]}]}`, `{"items":[{"type":"string","enum":["a"]}]}`},
		{"object intersection", `{"allOf":[{"type":"object"},{"type":"object"}]}`, `{"allOf":[{"type":"object"},{"type":"object"}]}`},
		{"two patterns", `{"allOf":[{"type":"string","pattern":"a"},{"pattern":"b"}]}`, `{"allOf":[{"type":"string","pattern":"a"},{"pattern":"b"}]}`},
	} {
		t.Run(tc.name, func(t *testing.T) {
			var input, want any
			if err := json.Unmarshal([]byte(tc.input), &input); err != nil {
				t.Fatal(err)
			}
			if err := json.Unmarshal([]byte(tc.want), &want); err != nil {
				t.Fatal(err)
			}
			normalize(input)
			if !reflect.DeepEqual(input, want) {
				t.Fatalf("got %#v, want %#v", input, want)
			}
		})
	}
}

func TestLiftRequestUnions(t *testing.T) {
	if containsUnion([]any{"plain", object{"type": "string"}}) {
		t.Fatal("non-union classified as a union")
	}
	for _, body := range []string{
		`{"oneOf":[{"type":"string"},{"type":"integer"}]}`,
		`{"type":"object","properties":{"items":{"type":"array","items":{"anyOf":[{"type":"string"},{"type":"array","items":{"type":"string"}}]}}}}`,
	} {
		var schema object
		check(json.Unmarshal([]byte(body), &schema))
		plain := object{"type": "object", "properties": object{"name": object{"type": "string"}}}
		media := object{"schema": schema}
		plainMedia := object{"schema": plain}
		spec := object{
			"components": object{"schemas": object{}},
			"paths": object{
				"/union": object{"post": object{"operationId": "post-sample", "requestBody": object{"content": object{"application/json": media}}}},
				"/plain": object{"post": object{"operationId": "postPlain", "requestBody": object{"content": object{"application/json": plainMedia}}}},
				"/empty": object{"get": object{"operationId": "getEmpty"}},
			},
		}
		liftRequestUnions(spec)
		defs := spec["components"].(object)["schemas"].(object)
		if len(defs) != 1 || !reflect.DeepEqual(defs["PostSampleJSONBody"], schema) {
			t.Fatal("union schema changed while lifting")
		}
		if media["schema"].(object)["$ref"] != "#/components/schemas/PostSampleJSONBody" || !reflect.DeepEqual(plainMedia["schema"], plain) {
			t.Fatal("wrong request reference or modified ordinary request")
		}
		liftRequestUnions(spec)
		if len(defs) != 1 {
			t.Fatal("lifting must be idempotent")
		}
	}
}

func TestLiftRequestUnionNameCollision(t *testing.T) {
	defer func() {
		if r := recover(); r == nil || !strings.Contains(r.(string), "collision") {
			t.Fatal("expected schema collision")
		}
	}()
	spec := object{
		"components": object{"schemas": object{"PostSampleJSONBody": object{"type": "string"}}},
		"paths": object{"/sample": object{"post": object{
			"operationId": "postSample", "requestBody": object{"content": object{"application/json": object{
				"schema": object{"oneOf": []any{object{"type": "string"}}},
			}}},
		}}},
	}
	liftRequestUnions(spec)
}

func TestMergeReferencePages(t *testing.T) {
	dir := t.TempDir()
	write := func(name, body string) {
		t.Helper()
		if err := os.WriteFile(filepath.Join(dir, name), []byte(body), 0600); err != nil {
			t.Fatal(err)
		}
	}
	write("index", "https://pennylane.readme.io/reference/first.md\nhttps://pennylane.readme.io/reference/first.md\nhttps://pennylane.readme.io/reference/second.md\nhttps://pennylane.readme.io/reference/dms_filecreated.md")
	write("dms_filecreated.md", "Webhook notification")
	for _, name := range []string{"first", "second"} {
		write(name+".md", "# OpenAPI definition\n```json\n"+`{"openapi":"3.0.3","info":{"title":"fixture","version":"1"},"servers":[],"components":{"schemas":{}},"paths":{"/`+name+`":{"get":{"operationId":"`+name+`","responses":{"200":{"content":{"application/json":{"schema":{"type":"object","properties":{"name":{"allOf":[{"type":"string","nullable":true}]}}}}}}}}}}}`+"\n```\n")
	}
	output, err := os.CreateTemp(dir, "output")
	if err != nil {
		t.Fatal(err)
	}
	oldArgs, oldStdout := os.Args, os.Stdout
	t.Cleanup(func() { os.Args, os.Stdout = oldArgs, oldStdout; output.Close() })
	os.Args = []string{"spec", filepath.Join(dir, "index"), dir}
	os.Stdout = output
	main()
	data, err := os.ReadFile(output.Name())
	if err != nil {
		t.Fatal(err)
	}
	var got struct {
		Paths map[string]map[string]struct {
			ExternalDocs struct{ URL string }
			Responses    map[string]struct {
				Content map[string]struct{ Schema map[string]any }
			}
		}
		Components struct{ Schemas map[string]any }
	}
	if err := json.Unmarshal(data, &got); err != nil {
		t.Fatal(err)
	}
	if len(got.Paths) != 2 || len(got.Components.Schemas) != 1 {
		t.Fatalf("paths=%d schemas=%d", len(got.Paths), len(got.Components.Schemas))
	}
	for _, name := range []string{"first", "second"} {
		op := got.Paths["/"+name]["get"]
		if op.ExternalDocs.URL != "https://pennylane.readme.io/reference/"+name {
			t.Fatal("source URL missing")
		}
		if op.Responses["200"].Content["application/json"].Schema["$ref"] != "#/components/schemas/FirstStatus200" {
			t.Fatal("shared response not deduplicated")
		}
	}
}
