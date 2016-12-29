package cloudshare

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerateToken(t *testing.T) {
	token := generateToken()
	assert.Equal(t, len(token), 10, "expecting token to be 10 chars long")
}

func TestAuthToken(t *testing.T) {
	actual := authToken("api_key", "api_id", "url")
	expectedPattern := "userapiid:api_id;timestamp:\\d+;token:[a-zA-Z0-9]{10};hmac:[0-9a-f]+"
	assert.Regexp(t, expectedPattern, actual)
}
