package refresh_token

import (
	"crypto/rand"
	"encoding/hex"
)

func GenerateRefreshToken() string {
	byte := make([]byte, 18)

	_, err := rand.Read(byte)
	if err != nil {
		return ""
	}

	return hex.EncodeToString(byte)
}
