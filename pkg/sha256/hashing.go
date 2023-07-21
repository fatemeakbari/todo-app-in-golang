package sha256

import (
	"crypto/sha256"
	"encoding/hex"
	"todo/pkg"
)

var hash = sha256.New()

type _sha256 struct {
}

func New() pkg.Hash {
	return &_sha256{}
}
func (sha _sha256) Hash(s string) string {
	return hex.EncodeToString(hash.Sum([]byte(s)))
}
