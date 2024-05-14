package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

func CreateHash(key string) string {
	// iniciando o modulo de sha256
	hasher := sha256.New()
	// transformando a string para byte e escrevendo o hash
	hasher.Write([]byte(key))
	// retornando o hash em sha256
	return hex.EncodeToString(hasher.Sum(nil))
}
