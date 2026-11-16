package pennylane

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

//go:generate go tool oapi-codegen -config internal/spec/config.yaml internal/spec/company.json

// Config contient un token Company, Firm ou un access token OAuth déjà obtenu.
type Config struct {
	Token string
	// CompanyID cible la société avec un token Firm. Zéro omet l'en-tête.
	CompanyID  int64
	HTTPClient *http.Client
}

// Client fournit les opérations HTTP brutes et les variantes WithResponse typées.
type Client struct{ *ClientWithResponses }

// New crée un client authentifié pour la Company API v2 de Pennylane.
func New(cfg Config) (*Client, error) {
	token := strings.TrimSpace(cfg.Token)
	if token == "" || strings.ContainsAny(token, " \t\r\n") {
		return nil, fmt.Errorf("invalid Pennylane bearer token")
	}
	if cfg.CompanyID < 0 {
		return nil, fmt.Errorf("invalid Pennylane company ID")
	}
	httpClient := cfg.HTTPClient
	if httpClient == nil {
		httpClient = &http.Client{Timeout: 60 * time.Second}
	}
	api, err := NewClientWithResponses("https://app.pennylane.com", WithHTTPClient(httpClient), WithRequestEditorFn(func(_ context.Context, r *http.Request) error {
		r.Header.Set("Authorization", "Bearer "+token)
		if cfg.CompanyID != 0 && r.Header.Get("X-Company-Id") == "" && !strings.HasPrefix(r.URL.Path, "/api/external/v2/webhook_subscriptions") {
			r.Header.Set("X-Company-Id", strconv.FormatInt(cfg.CompanyID, 10))
		}
		return nil
	}))
	if err != nil {
		return nil, err
	}
	return &Client{api}, nil
}

// Ping vérifie que /me accepte le token et retourne une société identifiable.
func (c *Client) Ping(ctx context.Context) error {
	r, err := c.GetMe(ctx, nil)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	if r.StatusCode != http.StatusOK {
		return fmt.Errorf("Pennylane GET /me: HTTP %d", r.StatusCode)
	}
	var profile GetMeStatus200
	if err := json.NewDecoder(r.Body).Decode(&profile); err != nil {
		return fmt.Errorf("Pennylane GET /me: %w", err)
	}
	if profile.Company.Id <= 0 {
		return fmt.Errorf("Pennylane GET /me: missing company ID")
	}
	return nil
}
