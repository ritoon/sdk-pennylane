package main

import (
	"bytes"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"testing"
)

// Compare all declarations (including method bodies), not just exported names.
func declarations(t *testing.T, files map[string][]byte) map[string]int {
	t.Helper()
	result := map[string]int{}
	for name, src := range files {
		fs := token.NewFileSet()
		f, err := parser.ParseFile(fs, name, src, 0)
		if err != nil {
			t.Fatal(err)
		}
		for _, d := range f.Decls {
			if g, ok := d.(*ast.GenDecl); ok && g.Tok == token.IMPORT {
				continue
			}
			var b bytes.Buffer
			if err := format.Node(&b, fs, d); err != nil {
				t.Fatal(err)
			}
			result[b.String()]++
		}
	}
	return result
}

func TestGenerationPreservesDeclarations(t *testing.T) {
	root, err := filepath.Abs("../..")
	if err != nil {
		t.Fatal(err)
	}
	dir := t.TempDir()
	for _, name := range []string{"go.mod", "go.sum", "internal/spec/config.yaml", "internal/spec/company.json"} {
		data, err := os.ReadFile(filepath.Join(root, name))
		if err != nil {
			t.Fatal(err)
		}
		dest := filepath.Join(dir, name)
		if err := os.MkdirAll(filepath.Dir(dest), 0755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(dest, data, 0600); err != nil {
			t.Fatal(err)
		}
	}
	t.Chdir(dir)
	cmd := exec.Command("go", "tool", "oapi-codegen", "-config", "internal/spec/config.yaml", "-o", "api.gen.go", "internal/spec/company.json")
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("%v: %s", err, out)
	}
	original, err := os.ReadFile("api.gen.go")
	if err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile("api_obsolete.gen.go", []byte(header), 0600); err != nil {
		t.Fatal(err)
	}
	if err := run(); err != nil {
		t.Fatal(err)
	}
	readFiles := func() map[string][]byte {
		t.Helper()
		names, err := filepath.Glob("api_*.gen.go")
		if err != nil {
			t.Fatal(err)
		}
		files := map[string][]byte{}
		for _, name := range names {
			data, err := os.ReadFile(name)
			if err != nil {
				t.Fatal(err)
			}
			files[name] = data
		}
		return files
	}
	files := readFiles()
	if len(files) != 17 {
		t.Fatalf("got %d domain files, want 17", len(files))
	}
	if _, err := os.Stat("api.gen.go"); !os.IsNotExist(err) {
		t.Fatal("monolithic output was not removed")
	}
	if _, ok := files["api_obsolete.gen.go"]; ok {
		t.Fatal("stale output was not removed")
	}
	if !reflect.DeepEqual(declarations(t, map[string][]byte{"original.go": original}), declarations(t, files)) {
		t.Fatal("declarations changed during splitting")
	}
	for name, data := range files {
		want, err := os.ReadFile(filepath.Join(root, name))
		if err != nil {
			t.Fatal(err)
		}
		if !bytes.Equal(data, want) {
			t.Errorf("%s needs regeneration", name)
		}
	}
	if err := run(); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(files, readFiles()) {
		t.Fatal("generation is not reproducible")
	}
}

func TestInvalidInputs(t *testing.T) {
	for _, spec := range []string{`broken`, `{"paths":{"/api/external/v2/unknown":{"get":{"operationId":"getUnknown"}}}}`} {
		if _, err := operationDomains([]byte(spec)); err == nil {
			t.Fatal("accepted invalid contract")
		}
	}
	if _, err := split([]byte("not Go"), []byte(`{}`)); err == nil {
		t.Fatal("accepted invalid Go")
	}
	if _, err := split(nil, []byte("broken")); err == nil {
		t.Fatal("accepted invalid JSON")
	}
}
