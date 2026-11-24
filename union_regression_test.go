package pennylane

import (
	"encoding/json"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/oapi-codegen/runtime/types"
)

// Exercise the public From/As/Merge helpers for every request union variant,
// including nested unions reachable only through another union's As method.
func TestRequestUnionVariants(t *testing.T) {
	client := reflect.TypeOf((*Client)(nil))
	seen := map[reflect.Type]bool{}
	var visit func(reflect.Type)
	visit = func(typ reflect.Type) {
		if typ.Kind() == reflect.Pointer || typ.Kind() == reflect.Slice || typ.Kind() == reflect.Map {
			visit(typ.Elem())
			return
		}
		if typ.Kind() != reflect.Struct || seen[typ] {
			return
		}
		seen[typ] = true
		if _, union := typ.FieldByName("union"); union {
			if !typ.Implements(reflect.TypeOf((*json.Marshaler)(nil)).Elem()) || !reflect.PointerTo(typ).Implements(reflect.TypeOf((*json.Unmarshaler)(nil)).Elem()) {
				t.Errorf("%s lacks JSON union methods", typ)
				return
			}
			ptr := reflect.PointerTo(typ)
			for i := 0; i < ptr.NumMethod(); i++ {
				method := ptr.Method(i)
				if !strings.HasPrefix(method.Name, "From") {
					continue
				}
				variant := method.Type.In(1)
				t.Run(typ.Name()+"/"+variant.Name(), func(t *testing.T) {
					value := unionExample(t, variant)
					want, err := json.Marshal(value.Interface())
					if err != nil {
						t.Fatal(err)
					}
					u := reflect.New(typ)
					if r := u.MethodByName(method.Name).Call([]reflect.Value{value}); !r[0].IsNil() {
						t.Fatal(r[0].Interface())
					}
					for _, action := range []string{"As", "Merge", "As"} {
						name := action + strings.TrimPrefix(method.Name, "From")
						var args []reflect.Value
						if action == "Merge" {
							args = []reflect.Value{value}
						}
						result := u.MethodByName(name).Call(args)
						if !result[len(result)-1].IsNil() {
							t.Fatal(result[len(result)-1].Interface())
						}
						if action == "As" {
							got, err := json.Marshal(result[0].Interface())
							if err != nil {
								t.Fatal(err)
							}
							assertJSONFields(t, string(want), got)
						}
					}
					encoded, err := json.Marshal(u.Interface())
					if err != nil {
						t.Fatal(err)
					}
					decoded := reflect.New(typ)
					if err := json.Unmarshal(encoded, decoded.Interface()); err != nil {
						t.Fatal(err)
					}
					got, err := json.Marshal(decoded.Interface())
					if err != nil {
						t.Fatal(err)
					}
					assertJSONFields(t, string(want), got)
				})
				visit(variant)
			}
		}
		for i := 0; i < typ.NumField(); i++ {
			if field := typ.Field(i); field.IsExported() {
				visit(field.Type)
			}
		}
	}
	for i := 0; i < client.NumMethod(); i++ {
		method := client.Method(i)
		if !strings.HasSuffix(method.Name, "WithResponse") {
			continue
		}
		for j := 2; j < method.Type.NumIn()-1; j++ {
			visit(method.Type.In(j))
		}
	}
}

// Nonzero data catches dropped fields without depending on any live account.
// Enum membership and business validation belong to the API, not these codec tests.
func unionExample(t *testing.T, typ reflect.Type) reflect.Value {
	t.Helper()
	v := reflect.New(typ).Elem()
	date := time.Date(2026, 11, 24, 0, 0, 0, 0, time.UTC)
	if typ == reflect.TypeOf(time.Time{}) {
		return reflect.ValueOf(date)
	}
	if typ == reflect.TypeOf(types.Date{}) {
		return reflect.ValueOf(types.Date{Time: date})
	}
	if typ == reflect.TypeOf(types.Email("")) {
		return reflect.ValueOf(types.Email("audit@example.org"))
	}
	switch typ.Kind() {
	case reflect.Pointer:
		v.Set(reflect.New(typ.Elem()))
		v.Elem().Set(unionExample(t, typ.Elem()))
	case reflect.Struct:
		if _, union := typ.FieldByName("union"); union {
			ptr := reflect.New(typ)
			for i := 0; i < ptr.NumMethod(); i++ {
				method := ptr.Type().Method(i)
				if strings.HasPrefix(method.Name, "From") {
					arg := unionExample(t, method.Type.In(1))
					result := ptr.MethodByName(method.Name).Call([]reflect.Value{arg})
					if !result[0].IsNil() {
						t.Fatal(result[0].Interface())
					}
					return ptr.Elem()
				}
			}
			t.Fatalf("no constructor for %s", typ)
		}
		for i := 0; i < typ.NumField(); i++ {
			if v.Field(i).CanSet() {
				v.Field(i).Set(unionExample(t, typ.Field(i).Type))
			}
		}
	case reflect.Slice:
		v.Set(reflect.MakeSlice(typ, 1, 1))
		v.Index(0).Set(unionExample(t, typ.Elem()))
	case reflect.Map:
		v.Set(reflect.MakeMap(typ))
		v.SetMapIndex(unionExample(t, typ.Key()), unionExample(t, typ.Elem()))
	case reflect.String:
		v.SetString("audit")
	case reflect.Bool:
		v.SetBool(true)
	case reflect.Int, reflect.Int64:
		v.SetInt(73)
	case reflect.Float32, reflect.Float64:
		v.SetFloat(1.25)
	}
	return v
}

func TestOptionalUnionNullAndOmission(t *testing.T) {
	var body ImportCustomerInvoicesJSONRequestBody
	data, err := json.Marshal(body)
	if err != nil {
		t.Fatal(err)
	}
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		t.Fatal(err)
	}
	if _, exists := fields["sale_type"]; exists {
		t.Fatal("unset sale_type must be omitted")
	}
	// Keep an explicit pointer to the null union when asking the API to clear it.
	var saleType ImportCustomerInvoicesJSONBody_SaleType
	if err := json.Unmarshal([]byte("null"), &saleType); err != nil {
		t.Fatal(err)
	}
	body.SaleType = &saleType
	data, err = json.Marshal(body)
	if err != nil {
		t.Fatal(err)
	}
	assertJSONFields(t, `{"sale_type":null}`, data)
}
