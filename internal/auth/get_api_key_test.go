package auth

import (
	"errors"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	// Check if the header does not exist
	header := make(http.Header)
	_, err := GetAPIKey(header)
	if err == nil {
		t.Fatalf("expected: %v, got: %v", ErrNoAuthHeaderIncluded, err)
	}

	// Check if the header is correctly formatted
	header.Set("APIAuthorization", "nonesense")
	_, err = GetAPIKey(header)
	if err == nil {
		t.Fatalf("expected: %v, got: %v", errors.New("malformed authorization header"), err)
	}

	// Check if the return value is correct for the correct headers
	header.Set("Authorization", "ApiKey 1234")
	res, err := GetAPIKey(header)
	if res != "1234" || err != nil {
		t.Fatalf("expected: (%v, %v), got: (%v, %v)", "1234", nil, res, err)
	}
}
