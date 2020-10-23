package internal

import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"syscall"

	"golang.org/x/crypto/ssh/terminal"
)

// DecryptVault takes a path to a vault and the key it has bin encrypted with.
// It opens and reads the contend returns the AES decrypted content in form
// of map[string]string
func DecryptVault(path, vaultkey string) (map[string]string, error) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	hash, err := hashkey(vaultkey)
	if err != nil {
		return nil, err
	}

	decrypted, err := decrypt(hash, file)
	if err != nil {
		return nil, err
	}

	return sliceToMap(decrypted)
}

// EncryptVault takes the new vault does AES encryption with the users
// password. Before rewriting the data in the file it takes a backup of the
// current file (in case os fails vault is not lost) after success it writes the new
// and deletes the old file
func EncryptVault(path, vaultkey string, vault []byte) error {

	hash, err := hashkey(vaultkey)
	if err != nil {
		return err
	}
	encrypted, err := encrypt(hash, vault)
	if err != nil {
		return err
	}

	// backup current vault in case of drama
	backuppath, err := renameFile(path)
	if err != nil {
		// rollback happens new sh*t hit the fan
		return err
	}

	if err := writeFile(path, encrypted); err != nil {
		return err
	}

	if err := removeFile(backuppath); err != nil {
		return err
	}
	return nil
}

// InputPassword handels users password input
func InputPassword() (string, error) {

	fmt.Print("🔒: ")
	bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return "", err
	}
	password := string(bytePassword)
	fmt.Print("\n")
	return strings.TrimSpace(password), nil
}

// InputCredentials handels user/account name and password input from the user
func InputCredentials() (string, string, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("👽: ")
	username, _ := reader.ReadString('\n')

	fmt.Print("🔏: ")
	bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return "", "", err
	}
	password := string(bytePassword)

	return strings.TrimSpace(username), strings.TrimSpace(password), nil
}

func hashkey(key string) ([]byte, error) {
	hash := sha256.Sum256([]byte(key))

	hexHash := hex.EncodeToString(hash[:])
	return []byte(hexHash), nil
}

// sliceToMap unmarshales a byte slice to a map[string]string
// if the length of slice is zero it returns an empty map[string]string
func sliceToMap(data []byte) (map[string]string, error) {
	var v map[string]string
	if len(data) <= 0 {
		v = make(map[string]string)
		return v, nil
	}

	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return v, nil
}
