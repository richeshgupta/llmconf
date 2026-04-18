package cli

import "testing"

func TestCredentialNameForProvider(t *testing.T) {
	tests := []struct {
		provider string
		expected string
	}{
		{"litellm", "ANTHROPIC_AUTH_TOKEN"},
		{"fireworks", "ANTHROPIC_API_KEY"},
		{"anthropic", "ANTHROPIC_API_KEY"},
		{"bedrock", "ANTHROPIC_API_KEY"},
		{"vertex", "ANTHROPIC_API_KEY"},
		{"unknown", "ANTHROPIC_API_KEY"},
	}

	for _, tt := range tests {
		t.Run(tt.provider, func(t *testing.T) {
			got := credentialNameForProvider(tt.provider)
			if got != tt.expected {
				t.Errorf("credentialNameForProvider(%q) = %q, want %q", tt.provider, got, tt.expected)
			}
		})
	}
}
