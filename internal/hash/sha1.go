package hash

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
)

// Sha1 builds sha1 hash with proper salt
func Sha1(value, salt string) string {
	value = fmt.Sprintf(salt, value)
	hash := sha1.New()
	hash.Write([]byte(value))
	sha1Hash := hex.EncodeToString(hash.Sum(nil))
	return sha1Hash
}
