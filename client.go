package iid

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"

	"golang.org/x/oauth2/google"
	"golang.org/x/oauth2/jwt"
)

// New returns a client for importing APNs tokens to Firebase Cloud Messaging (FCM)
func New(bundleID string, credentials []byte, opts ...Option) (*Client, error) {
	client := &Client{
		// inject defaults
		credentials: credentials,

		Application: bundleID,
		Sandbox:     false,
		Scopes:      firebaseScopes,
	}

	for _, opt := range opts {
		opt(client)
	}

	cfg, err := google.JWTConfigFromJSON(client.credentials, client.Scopes...)
	if err != nil {
		return nil, err
	}

	client.jwt = cfg

	return client, nil
}

// Client provides the structure required to communicate with Google's IID
// endpoint.
type Client struct {
	credentials []byte
	jwt         *jwt.Config

	Application string
	Sandbox     bool
	Scopes      []string
}

type importTokensRequest struct {
	Application string   `json:"application"`
	Sandbox     bool     `json:"sandbox"`
	ApnsTokens  []string `json:"apns_tokens"`
}

type BatchImportResponse struct {
	Results []*BatchImportResults `json:"results"`
}

type BatchImportResults []struct {
	ApnsToken         string `json:"apns_token"`
	Status            string `json:"status"`
	RegistrationToken string `json:"registration_token,omitempty"`
}

// BatchImport imports APNs tokens returning an
func (c *Client) BatchImport(ctx context.Context, tokens []string) ([]*BatchImportResults, error) {
	if len(tokens) > maxTokens {
		return nil, ErrTooManyTokens
	}

	// Build and send the request
	payload := &importTokensRequest{
		Application: c.Application,
		Sandbox:     c.Sandbox,
		ApnsTokens:  tokens,
	}
	b, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, uriBatchImport, bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}

	tkn, err := c.jwt.TokenSource(ctx).Token()
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+tkn.AccessToken)
	req.Header.Set("Content-Type", "application/json; UTF-8")
	req.Header.Set("access_token_auth", "true")

	rawRes, err := c.jwt.Client(ctx).Do(req)
	if err != nil {
		return nil, err
	}

	// Process and return the response
	b, err = io.ReadAll(rawRes.Body)
	if err != nil {
		return nil, err
	}

	res := &BatchImportResponse{}
	err = json.Unmarshal(b, res)
	if err != nil {
		return nil, err
	}

	return res.Results, nil
}
