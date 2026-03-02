package links

import (
	"crypto/sha256"
	"encoding/hex"
)

func GenerateID(s string) string {
	sum := sha256.Sum256([]byte(s))
	return hex.EncodeToString((sum)[:16])
}