package lib

import (
    "crypto/sha256"
    "encoding/hex"
)

// Hash the refresh token using SHA-256
func HashToken(token string) string {
    hash := sha256.New()
    hash.Write([]byte(token))

    return hex.EncodeToString(hash.Sum(nil))
}
