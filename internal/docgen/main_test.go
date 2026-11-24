package main

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestNestedFieldsAndExamples(t *testing.T) {
	var s schema
	if err := json.Unmarshal([]byte(`{"type":"object","required":["items"],"properties":{"items":{"type":"array","items":{"type":"object","properties":{"paid":{"type":"boolean","example":false},"amount":{"type":"string","nullable":true}}}}}}`), &s); err != nil {
		t.Fatal(err)
	}
	var b strings.Builder
	fields(&b, s)
	for _, row := range []string{"`$.items` | array | obligatoire", "`$.items[].amount` | string ou null | facultatif"} {
		if !strings.Contains(b.String(), row) {
			t.Errorf("missing %s", row)
		}
	}
	value, ok := example(s)
	got, err := json.Marshal(value)
	if err != nil || !ok || string(got) != `{"items":[{"paid":false}]}` {
		t.Fatalf("unexpected example %s (%v)", got, err)
	}
}

func TestRequestSchemaReferences(t *testing.T) {
	var body schema
	if err := json.Unmarshal([]byte(`{"type":"object","properties":{"name":{"type":"string","example":"Audit"}}}`), &body); err != nil {
		t.Fatal(err)
	}
	defs := map[string]schema{"PostCustomerJSONBody": body}
	for _, input := range []schema{body, {"$ref": "#/components/schemas/PostCustomerJSONBody"}} {
		resolved := requestSchema(input, defs)
		var b strings.Builder
		fields(&b, resolved)
		illustration(&b, resolved)
		if !strings.Contains(b.String(), "`$.name`") || !strings.Contains(b.String(), `"name": "Audit"`) {
			t.Fatal("referenced request fields or example missing")
		}
	}
	for _, ref := range []string{"#/components/schemas/Unknown", "https://example.org/schema"} {
		t.Run(ref, func(t *testing.T) {
			defer func() {
				if recover() == nil {
					t.Fatal("unknown reference accepted")
				}
			}()
			requestSchema(schema{"$ref": ref}, defs)
		})
	}
}
