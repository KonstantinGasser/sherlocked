package internal

import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
	"syscall"

	"github.com/KonstantinGasser/sherlocked/cmd_errors"
	"golang.org/x/crypto/ssh/terminal"
)

func CheckVaultInit(path string) (bool, error) {
	fi, err := os.Stat(path)
	if err != nil {
		return false, err
	}
	if fi.Size() == 0 {
		return false, cmd_errors.InitNotDoneError{
			MSG: "🏁 two steps bofore you can start:\n1️⃣ run 'lock password' (this will set the password to encrypt/decrypt your vault)\n2️⃣ run 'lock add' to add your first password\nthat's it - you're ready to go 🎉🥳",
		}
	}
	return true, nil
}

// DecryptVault takes a path to a vault and the key it has bin encrypted with.
// It opens and reads the contend returns the AES decrypted content in form
// of map[string]string
func DecryptVault(path, vaultkey string) (map[string]string, error) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	hash := hashkey(vaultkey)
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

	hash := hashkey(vaultkey)
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
		return "", cmd_errors.OSStdInError{
			MSG: `🤨 Sorry ~ somehow we failed to read your password from os.Stdin..`,
		}
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
		return "", "", cmd_errors.OSStdInError{
			MSG: `🤨 Sorry ~ somehow we failed to read your password from os.Stdin..`,
		}
	}
	fmt.Print("\n")

	password := string(bytePassword)

	return strings.TrimSpace(username), strings.TrimSpace(password), nil
}

// InputNewPassword is used to collect the new password set by a user
func InputNewPassword(txt string) (string, error) {

	fmt.Print(txt)
	bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return "", cmd_errors.OSStdInError{
			MSG: `🤨 Sorry ~ somehow we failed to read your password from os.Stdin..`,
		}
	}
	password := string(bytePassword)
	fmt.Print("\n")
	return strings.TrimSpace(password), nil
}

func hashkey(key string) []byte {
	hash := sha256.Sum256([]byte(key))
	hexHash := hex.EncodeToString(hash[:])
	return []byte(hexHash)
}

// sliceToMap unmarshales a byte slice to a map[string]string
// if the length of slice is zero it returns an empty map[string]string
func sliceToMap(data []byte) (map[string]string, error) {
	var v map[string]string
	if len(data) <= 0 {
		v = make(map[string]string)
		return v, cmd_errors.ZeroVaultError{
			MSG: `🤷🏼‍♀️: Ups looks like you have not stored any password in the vault yet.
Use 'lock password' to set a vault password and then use
'lock add' to add a new password to the vault`,
		}
	}

	if err := json.Unmarshal(data, &v); err != nil {
		return nil, cmd_errors.MapConversionError{
			MSG: `🙅🏼‍♀️ hey! have you just misstyped the password
      or are you a bad bad boy? 🤨`,
		}
	}
	return v, nil
}

// EvalPassword evaluates the strength of a password (range between 0-100)
func EvaluatePassword(password string) int {
	var strength = 0

	regN := regexp.MustCompile(`[0-9]`)
	numbers := regN.FindAllString(password, -1)
	strength += len(numbers) * 4

	regC := regexp.MustCompile(`[A-Z]`)
	caper := regC.FindAllString(password, -1)
	strength += (len(password) - len(caper)) * 2

	regL := regexp.MustCompile(`[a-z]`)
	lower := regL.FindAllString(password, -1)
	strength += (len(password) - len(lower)) * 2

	regS := regexp.MustCompile(`[$#_-]`)
	specials := regS.FindAllString(password, -1)
	strength += len(specials) * 6

	if strength > 100 {
		return 100
	}
	return strength
}
