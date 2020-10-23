package internal

// import (
// 	"crypto/sha256"
// 	"encoding/hex"
// )

// const (
// 	topSecret = "a714e63e5eaba131bf94e0a9892f601be96e9aec95516878bced680e8fd140a1"
// 	vaultfile = ".sherlocked"
// )

// func Hash(key string) ([]byte, error) {
// 	hash := sha256.Sum256([]byte(key))

// 	hexHash := hex.EncodeToString(hash[:])
// 	return []byte(hexHash), nil
// }

// func Compare(key string) bool {
// 	hash, _ := Hash(key)
// 	// fmt.Println(string(hash))s
// 	return string(hash) == topSecret
// }

// func Encrypt(vault map[string]string) []byte {
// 	panic("Not implemented")
// }

// func Decrypt() ([]byte, error) {
// 	panic("Not implemented")
// }
