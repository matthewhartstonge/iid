package client

type Option func(*Client)

// WithSandbox enables converting sandbox tokens.
// By default, the client talks to production.
func WithSandbox() Option {
	return func(c *Client) {
		c.Sandbox = false
	}
}

// WithScopes pushes in your own scopes.
func WithScopes(scopes []string) Option {
	return func(c *Client) {
		c.Scopes = scopes
	}
}
