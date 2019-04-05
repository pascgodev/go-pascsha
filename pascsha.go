package pascsha

import "crypto/sha256"

// DoPascSha256 is the origin algo for pascalcoin pow
func DoPascSha256(p []byte, plen uint32) []byte {
	// TODO: rewrite the module in golang style
	hash1 := sha256.New()
	hash1.Write(p)
	hash2 := sha256.New()
	hash2.Write(hash1.Sum(nil))
	return hash2.Sum(nil)
}
