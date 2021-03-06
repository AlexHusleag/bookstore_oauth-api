package access_token

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestAccessTokenConstants(t *testing.T){
	assert.EqualValues(t, 24, expirationTime, "Expiration time should be 24 hours")
}

func Test(t *testing.T) {
	at := GetNewAccessToken()

	assert.False(t, at.IsExpired(), "Brand new access token should not be expired")
	assert.EqualValues(t, "", at.AccessToken, "New access token should not gave defined access token id")
	assert.True(t, at.UserId == 0, "New access token should not have ana associated user id")

}

func TestAccessTokenIsExpired(t *testing.T) {
	at := AccessToken{}
	assert.True(t, at.IsExpired(), "Empty access token should be expired by default")

	at.Expires = time.Now().UTC().Add(3 * time.Hour).Unix()

	assert.False(t, at.IsExpired(), "Access token expiring three hours from now should NOT be expired")

}
