package auth

import (
	"errors"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name       string
		headers    http.Header
		wantAPIKey string
		wantErr    error
	}{
		{
			name: "Valid API key",
			headers: func() http.Header {
				h := http.Header{}
				h.Set("Authorization", "ApiKey 12345abcde")
				return h
			}(),
			wantAPIKey: "12345abcde",
			wantErr:    nil,
		},
		{
			name: "No Authorization header",
			headers: func() http.Header {
				return http.Header{}
			}(),
			wantAPIKey: "",
			wantErr:    errors.New("no authorization header included"),
		},
		{
			name: "Malformed Authorization header",
			headers: func() http.Header {
				h := http.Header{}
				h.Set("Authorization", "ApiKey")
				return h
			}(),
			wantAPIKey: "",
			wantErr:    errors.New("malformed authorization header"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			apiKey, err := GetAPIKey(tt.headers)
			if (err != nil && tt.wantErr == nil) || (err == nil && tt.wantErr != nil) || (err != nil && tt.wantErr != nil && err.Error() != tt.wantErr.Error()) {
				t.Errorf("Test %s failed: expected error %v, got %v", tt.name, tt.wantErr, err)
			}
			if apiKey != tt.wantAPIKey {
				t.Errorf("Test %s failed: expected API key %s, got %s", tt.name, tt.wantAPIKey, apiKey)
			}
		})
	}
}
