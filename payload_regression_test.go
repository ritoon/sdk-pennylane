package pennylane

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"testing"
)

// These payload fragments isolate transport regressions. They are not complete
// business documents and are never sent to a live Pennylane company.
func TestRequestPayloadFields(t *testing.T) {
	cases := []struct{ method, payload string }{
		{"ImportCustomerInvoices", `{"sale_type":"services","invoice_lines":[{"performance_period":{"date_type":"service_period","start_date":"2026-11-01","end_date":"2026-11-24"}}]}`},
		{"PostBillingSubscriptions", `{"mode":{"type":"email","email_settings":{"recipients":["audit@example.org"]}}}`},
		{"PutBillingSubscriptions", `{"mode":{"type":"email","email_settings":{"recipients":["audit@example.org"]}}}`},
		{"PostCategories", `{"direction":"cash_in"}`},
		{"PostCommercialDocuments", `{"invoice_lines":[{"label":"Service","quantity":2,"raw_currency_unit_price":"12.50","custom_fields":[{"name":"project","value":"alpha"}]}]}`},
		{"PostCustomer", `{"customer_type":"company","name":"Audit","billing_address":{"address":"1 rue du Test","city":"Paris","postal_code":"75001","country_alpha2":"FR"}}`},
		{"PostCustomer", `{"customer_type":"individual","first_name":"Ada","last_name":"Test","billing_address":{"address":"1 rue du Test","city":"Paris","postal_code":"75001","country_alpha2":"FR"}}`},
		{"PostCustomerInvoices", `{"customer_id":42,"date":"2026-11-24","deadline":"2026-12-24","invoice_lines":[{"product_id":7,"quantity":2}]}`},
		{"PostProducts", `{"ean_code":"12345678","custom_fields":[{"name":"project","value":"alpha"},{"name":"teams","value":["A","B"]}]}`},
		{"PutProduct", `{"ean_code":"1234567890123","custom_fields":[{"name":"teams","value":["A","B"]}]}`},
		{"PostQuotes", `{"invoice_lines":[{"product_id":7,"quantity":2}]}`},
		{"PutSupplierInvoice", `{"invoice_lines":{"update":[{"id":7,"ledger_entry_line":{"label":"Service"}}]}}`},
		{"PutSupplierInvoiceEInvoiceStatus", `{"status":"disputed","reason":"incorrect_vat_rate"}`},
		{"PutSupplierInvoiceEInvoiceStatus", `{"status":"refused","reason":"duplicate_invoice"}`},
		{"PutSupplierInvoiceEInvoiceStatus", `{"status":"approved"}`},
		{"UpdateCommercialDocument", `{"invoice_lines":{"create":[{"label":"Service","custom_fields":[{"name":"teams","value":["A","B"]}]}],"update":[{"id":7,"custom_fields":[{"name":"project","value":"alpha"}]}]}}`},
		{"UpdateCustomerInvoice", `{"label":"Updated","invoice_lines":{"create":[{"product_id":7,"quantity":2}]}}`},
		{"UpdateImportedCustomerInvoice", `{"invoice_lines":{"create":[{"performance_period":{"date_type":"service_date","end_date":"2026-11-24"}}],"update":[{"id":7,"performance_period":{"date_type":"service_period","start_date":"2026-11-01","end_date":"2026-11-24"}}]}}`},
		{"UpdateQuote", `{"invoice_lines":{"create":[{"product_id":7,"quantity":2}],"update":[{"id":7,"custom_fields":[{"name":"project","value":"alpha"}]}]}}`},
		{"UpdateTransaction", `{"customer_id":42}`},
		{"UpdateTransaction", `{"supplier_id":73}`},
		{"UpdateTransaction", `{"customer_id":null}`},
	}
	for _, tc := range cases {
		t.Run(tc.method, func(t *testing.T) {
			c := mockClient(t, 400, `{}`, func(r *http.Request) {
				body, err := io.ReadAll(r.Body)
				if err != nil {
					t.Fatal(err)
				}
				assertJSONFields(t, tc.payload, body)
			})
			invokePayload(t, c, tc.method, tc.payload)
		})
	}
}

func TestResponsePayloadFields(t *testing.T) {
	cases := []struct {
		method, payload string
		status          int
	}{
		{"GetCommercialDocumentInvoiceLines", `{"items":[{"custom_fields":[{"name":"project","value":"alpha"},{"name":"teams","value":["A","B"]}]}]}`, 200},
		{"GetCustomerInvoiceInvoiceLines", `{"items":[{"ean_code":"12345678","ledger_account":{"id":7,"number":"706","url":"https://app.pennylane.com/api/external/v2/ledger_accounts/7"}}]}`, 200},
		{"GetCustomerInvoices", `{"included":{"invoice_lines":{"42":{"items":[{"ean_code":"12345678","ledger_account":{"id":7,"number":"706","url":"https://app.pennylane.com/api/external/v2/ledger_accounts/7"}}]}}}}`, 200},
		{"GetSupplierInvoiceLines", `{"items":[{"ledger_account":{"id":7,"number":"606","url":"https://app.pennylane.com/api/external/v2/ledger_accounts/7"}}]}`, 200},
		{"GetSupplier", `{"bic":"BNPAFRPP"}`, 200},
		{"GetSuppliers", `{"items":[{"bic":"BNPAFRPP"}]}`, 200},
		{"PostSupplier", `{"bic":"BNPAFRPP"}`, 201},
		{"PutSupplier", `{"bic":"BNPAFRPP"}`, 200},
		{"GetSupplier", `{"bic":null}`, 200},
		{"GetCustomerInvoiceInvoiceLines", `{"items":[{"ean_code":null}]}`, 200},
	}
	for _, tc := range cases {
		t.Run(tc.method, func(t *testing.T) {
			c := mockClient(t, tc.status, tc.payload, func(*http.Request) {})
			response := invokePayload(t, c, tc.method, "")
			field := response.Elem().FieldByName(fmt.Sprintf("JSON%d", tc.status))
			if !field.IsValid() || field.IsNil() {
				t.Fatal("missing decoded response")
			}
			data, err := json.Marshal(field.Interface())
			if err != nil {
				t.Fatal(err)
			}
			assertJSONFields(t, tc.payload, data)
		})
	}
}

func invokePayload(t *testing.T, c *Client, name, payload string) reflect.Value {
	t.Helper()
	method := reflect.ValueOf(c).MethodByName(name + "WithResponse")
	if !method.IsValid() {
		t.Fatalf("missing method %s", name)
	}
	args := []reflect.Value{reflect.ValueOf(context.Background())}
	for i := 1; i < method.Type().NumIn()-1; i++ {
		typ := method.Type().In(i)
		v := reflect.New(typ).Elem()
		if i == method.Type().NumIn()-2 && payload != "" {
			if err := json.Unmarshal([]byte(payload), v.Addr().Interface()); err != nil {
				t.Fatal(err)
			}
		} else {
			switch typ.Kind() {
			case reflect.Pointer:
				v.Set(reflect.New(typ.Elem()))
			case reflect.Int, reflect.Int64:
				v.SetInt(73)
			}
		}
		args = append(args, v)
	}
	results := method.Call(args)
	if !results[1].IsNil() {
		t.Fatal(results[1].Interface())
	}
	return results[0]
}

func assertJSONFields(t *testing.T, expected string, encoded []byte) {
	t.Helper()
	var want, got any
	if err := json.Unmarshal([]byte(expected), &want); err != nil {
		t.Fatal(err)
	}
	if err := json.Unmarshal(encoded, &got); err != nil {
		t.Fatal(err)
	}
	var check func(any, any, string)
	check = func(w, g any, path string) {
		t.Helper()
		switch w := w.(type) {
		case map[string]any:
			m, ok := g.(map[string]any)
			if !ok {
				t.Fatalf("%s: expected object, got %#v", path, g)
			}
			for key, value := range w {
				actual, exists := m[key]
				if !exists {
					t.Errorf("%s/%s: field lost", path, key)
					continue
				}
				check(value, actual, path+"/"+key)
			}
		case []any:
			a, ok := g.([]any)
			if !ok || len(a) != len(w) {
				t.Fatalf("%s: array changed: %#v", path, g)
			}
			for i, value := range w {
				check(value, a[i], fmt.Sprintf("%s/%d", path, i))
			}
		default:
			if !reflect.DeepEqual(w, g) {
				t.Errorf("%s: got %#v, want %#v", path, g, w)
			}
		}
	}
	check(want, got, "")
}
