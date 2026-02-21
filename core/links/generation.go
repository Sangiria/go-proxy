package links

import (
	"core/models"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
)

func ParsedToKey(p models.Parsed) string {
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

func GenerateID(s string) string {
	sum := sha256.Sum256([]byte(s))
	return hex.EncodeToString((sum)[:16])
}