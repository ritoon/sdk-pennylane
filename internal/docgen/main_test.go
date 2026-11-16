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
