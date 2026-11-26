package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestGenerateDocumentation(t *testing.T) {
	root, err := filepath.Abs("../..")
	if err != nil {
		t.Fatal(err)
	}
	dir := t.TempDir()
	files, err := filepath.Glob(filepath.Join(root, "api_*.gen.go"))
	if err != nil {
		t.Fatal(err)
	}
	names := []string{"internal/spec/company.json", "multipart.go"}
	for _, file := range files {
		names = append(names, filepath.Base(file))
	}
	for _, name := range names {
		data, err := os.ReadFile(filepath.Join(root, name))
		if err != nil {
			t.Fatal(err)
		}
		target := filepath.Join(dir, name)
		if err := os.MkdirAll(filepath.Dir(target), 0755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(target, data, 0600); err != nil {
			t.Fatal(err)
		}
	}
	t.Chdir(dir)
	main()
	entries, err := os.ReadDir("docs/endpoints")
	if err != nil {
		t.Fatal(err)
	}
	if len(entries) != 187 {
		t.Fatalf("got %d pages, want 186 endpoints plus errors", len(entries))
	}
	for name, needles := range map[string][]string{
		"docs/operations.md":                                    {"GetJournalsWithResponse", "endpoints/getJournals.md"},
		"docs/endpoints/getJournals.md":                         {"params.Cursor", "réponse.JSON200", "next_cursor"},
		"docs/endpoints/postCustomer.md":                        {"billing_address", "customer_type", "oneOf"},
		"docs/endpoints/importCustomerInvoices.md":              {"sale_type", "performance_period", "file_attachment_id"},
		"docs/endpoints/createCustomerInvoiceEInvoiceImport.md": {"multipart/form-data", "application/json", "send_to_pa"},
		"docs/endpoints/erreurs.md":                             {"Payloads d’erreur"},
	} {
		data, err := os.ReadFile(name)
		if err != nil {
			t.Fatal(err)
		}
		for _, needle := range needles {
			if !strings.Contains(string(data), needle) {
				t.Errorf("%s missing %q", name, needle)
			}
		}
	}
	// A second generation must produce byte-for-byte identical documentation.
	before := string(read("docs/operations.md"))
	main()
	if string(read("docs/operations.md")) != before {
		t.Fatal("unstable operation index")
	}
}
