package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	headers := http.Header{
		"Authorization": []string{"ApiKey api-key"},
	}

	apiKey, err := GetAPIKey(headers)
	if err != nil {
		t.Fatalf("err: %s", err.Error())
	}

	if apiKey != "api-key" {
		t.Fatalf("expected: %s, got %s", "api-key", apiKey)
	}
}
