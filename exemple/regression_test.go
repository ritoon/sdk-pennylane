package main

import (
	"context"
	"errors"
	"io"
	"net/http"
	"os"
	"strings"
	"testing"

	pennylane "github.com/ritoon/sdk-pennylane"
)

func TestInvalidInputDoesNotSendRequests(t *testing.T) {
	c, err := pennylane.New(pennylane.Config{Token: "test", HTTPClient: &http.Client{Transport: transport(func(*http.Request) (*http.Response, error) {
		t.Fatal("invalid input reached HTTP transport")
		return nil, nil
	})}})
	if err != nil {
		t.Fatal(err)
	}
	for _, tc := range []struct {
		name string
		args []string
		body string
	}{
		{"unknown command", []string{"unknown"}, ""},
		{"missing args", []string{"balance"}, ""},
		{"invalid start", []string{"balance", "invalid", "2026-01-01"}, ""},
		{"invalid end", []string{"balance", "2026-01-01", "invalid"}, ""},
		{"reversed period", []string{"balance", "2026-02-01", "2026-01-01"}, ""},
		{"invalid invoice", []string{"envoyer-facture", "0", "test@example.org"}, ""},
		{"invalid email", []string{"envoyer-facture", "18", "bad"}, ""},
		{"bad JSON", []string{"creer-client"}, "{"},
		{"unknown field", []string{"creer-client"}, `{"typo":true}`},
		{"multiple documents", []string{"creer-client"}, `{} {}`},
		{"trailing garbage", []string{"creer-client"}, `{} invalid`},
		{"incomplete client", []string{"creer-client"}, `{"customer_type":"company"}`},
		{"invalid customer", []string{"creer-mandat", "-1"}, `{}`},
		{"invalid mandate JSON", []string{"creer-mandat", "42"}, `{`},
		{"incomplete mandate", []string{"creer-mandat", "42"}, `{}`},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if err := run(context.Background(), c, tc.args, strings.NewReader(tc.body), io.Discard); err == nil {
				t.Fatal("accepted invalid input")
			}
		})
	}
}

func TestCommandFailures(t *testing.T) {
	failure := errors.New("transport unavailable")
	for _, command := range []struct {
		name string
		args []string
		file string
	}{
		{"ping", nil, ""}, {"balance", []string{"balance", "2026-01-01", "2026-02-01"}, ""},
		{"invoice", []string{"envoyer-facture", "18", "test@example.org"}, ""},
		{"customer", []string{"creer-client"}, "client.json"}, {"mandate", []string{"creer-mandat", "42"}, "mandat.json"},
	} {
		t.Run(command.name, func(t *testing.T) {
			var input []byte
			if command.file != "" {
				var err error
				input, err = os.ReadFile("donnees/" + command.file)
				if err != nil {
					t.Fatal(err)
				}
			}
			for _, tc := range []struct {
				name   string
				status int
				body   string
				err    error
			}{
				{"transport", 0, "", failure}, {"unauthorized", 401, `{}`, nil}, {"missing identity", 201, `{}`, nil},
			} {
				t.Run(tc.name, func(t *testing.T) {
					calls := 0
					c, err := pennylane.New(pennylane.Config{Token: "test", HTTPClient: &http.Client{Transport: transport(func(r *http.Request) (*http.Response, error) {
						calls++
						if tc.err != nil {
							return nil, tc.err
						}
						return &http.Response{StatusCode: tc.status, Header: http.Header{"Content-Type": {"application/json"}}, Body: io.NopCloser(strings.NewReader(tc.body)), Request: r}, nil
					})}})
					if err != nil {
						t.Fatal(err)
					}
					err = run(context.Background(), c, command.args, strings.NewReader(string(input)), io.Discard)
					if err == nil {
						t.Fatal("failure reported as success")
					}
					if calls != 1 {
						t.Fatalf("got %d HTTP calls", calls)
					}
					if tc.err != nil && !errors.Is(err, failure) {
						t.Fatalf("transport error lost: %v", err)
					}
				})
			}
		})
	}
}

func TestBalanceRejectsBrokenPagination(t *testing.T) {
	for _, cursor := range []string{`null`, `""`, `"repeat"`} {
		t.Run(cursor, func(t *testing.T) {
			calls := 0
			c, err := pennylane.New(pennylane.Config{Token: "test", HTTPClient: &http.Client{Transport: transport(func(r *http.Request) (*http.Response, error) {
				calls++
				if calls > 2 {
					t.Fatal("pagination did not stop")
				}
				if calls == 2 && r.URL.Query().Get("cursor") != "repeat" {
					t.Fatal("cursor not forwarded")
				}
				return &http.Response{StatusCode: 200, Header: http.Header{"Content-Type": {"application/json"}}, Body: io.NopCloser(strings.NewReader(`{"items":[],"has_more":true,"next_cursor":` + cursor + `}`)), Request: r}, nil
			})}})
			if err != nil {
				t.Fatal(err)
			}
			err = balance(context.Background(), c, "2026-01-01", "2026-02-01", io.Discard)
			if err == nil || !strings.Contains(err.Error(), "curseur") {
				t.Fatalf("error = %v", err)
			}
			want := 1
			if cursor == `"repeat"` {
				want = 2
			}
			if calls != want {
				t.Fatalf("calls = %d, want %d", calls, want)
			}
		})
	}
}
