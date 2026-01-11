package utils

import (
	"core/models"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
)

func GenerateNodeKey(p models.Parsed) string {
	key := fmt.Sprintf(
		"vless|%s|%d|%s|%s",
		strings.ToLower(p.Address),
		p.Port,
		p.UUID,
		string(p.Transport),
	)

	if p.Security == models.SecurityReality {
		key += fmt.Sprintf(
			"|security=reality|sni=%s|fp=%s|pbk=%s|sid=%s|flow=%s",
			strings.ToLower(p.Sni),
			p.Fp,
			p.Pbk,
			strings.ToLower(p.Sid),
			p.Flow,
		)
	}

	return key
}

func DeterministicID(p models.Parsed) string {
	sum := sha256.Sum256([]byte(GenerateNodeKey(p)))
	return "vless_" + hex.EncodeToString(sum[:16])
}