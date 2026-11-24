package pennylane

import (
	"context"
	"errors"
	"io"
	"net/http"
	"strings"
	"testing"
	"time"
)

type trackedBody struct {
	io.Reader
	closed bool
}

func (b *trackedBody) Close() error { b.closed = true; return nil }

type brokenReader struct{ err error }

func (r brokenReader) Read([]byte) (int, error) { return 0, r.err }

func TestDefaultClientTimeout(t *testing.T) {
	c, err := New(Config{Token: "test"})
	if err != nil {
		t.Fatal(err)
	}
	raw := c.ClientInterface.(*APIClient).Client.(*http.Client)
	if raw.Timeout != 60*time.Second {
		t.Fatalf("timeout = %s", raw.Timeout)
	}
}

func TestResponseLifecycle(t *testing.T) {
	failure := errors.New("read failed")
	for _, tc := range []struct {
		name        string
		status      int
		contentType string
		reader      io.Reader
		wantErr     bool
	}{
		{"typed JSON", 200, "application/json; charset=utf-8", strings.NewReader(`{"items":[{"id":7,"code":"VE","label":"Ventes","type":"sales"}],"has_more":false,"next_cursor":null}`), false},
		{"malformed JSON", 200, "application/json", strings.NewReader(`{"items":`), true},
		{"read failure", 200, "application/json", brokenReader{failure}, true},
		{"HTML gateway error", 502, "text/html", strings.NewReader("<h1>Bad gateway</h1>"), false},
	} {
		t.Run(tc.name, func(t *testing.T) {
			body := &trackedBody{Reader: tc.reader}
			c, err := New(Config{Token: "test", HTTPClient: &http.Client{Transport: transportFunc(func(r *http.Request) (*http.Response, error) {
				return &http.Response{StatusCode: tc.status, Header: http.Header{"Content-Type": {tc.contentType}}, Body: body, Request: r}, nil
			})}})
			if err != nil {
				t.Fatal(err)
			}
			r, err := c.GetJournalsWithResponse(context.Background(), nil)
			if (err != nil) != tc.wantErr {
				t.Fatalf("error = %v", err)
			}
			if !body.closed {
				t.Fatal("response body not closed")
			}
			if tc.name == "read failure" && !errors.Is(err, failure) {
				t.Fatal("read error lost")
			}
			if tc.name == "typed JSON" && (r.JSON200 == nil || r.JSON200.Items[0].Code != "VE") {
				t.Fatal("typed payload lost")
			}
			if tc.status == 502 && (r.StatusCode() != 502 || string(r.Body) != "<h1>Bad gateway</h1>" || r.JSON200 != nil) {
				t.Fatal("gateway response lost")
			}
		})
	}
}

func TestRequestEditorErrorStopsTransport(t *testing.T) {
	c := mockClient(t, 200, "{}", func(*http.Request) { t.Fatal("editor failure reached transport") })
	failure := errors.New("editor failed")
	_, err := c.GetJournalsWithResponse(context.Background(), nil, func(context.Context, *http.Request) error { return failure })
	if !errors.Is(err, failure) {
		t.Fatalf("error = %v", err)
	}
}

func TestCanceledContext(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	c, err := New(Config{Token: "test", HTTPClient: &http.Client{Transport: transportFunc(func(r *http.Request) (*http.Response, error) {
		if r.Context() != ctx {
			t.Fatal("context not forwarded")
		}
		return nil, r.Context().Err()
	})}})
	if err != nil {
		t.Fatal(err)
	}
	if err := c.Ping(ctx); !errors.Is(err, context.Canceled) {
		t.Fatalf("error = %v", err)
	}
}
