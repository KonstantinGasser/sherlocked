package internal

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"regexp"

	"github.com/KonstantinGasser/sherlocked/cmd_errors"
)

type Vault struct {
	Path string
}

// Decrypt takes a key and a byte slice and decrypts the AES encrypted byte slice
func (v *Vault) Decrypt(key string, file []byte) (map[string]string, error) {

	aeskey, _ := v.hash(key)

	block, err := aes.NewCipher(aeskey[:16])
	if err != nil {
		panic(err)
	}

	if len(file) < aes.BlockSize {
		return make(map[string]string), nil
	}

	iv := file[:aes.BlockSize]

	ciphervault := file[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)

	stream.XORKeyStream(ciphervault, ciphervault)

	return v.deserialize(ciphervault)
}

// Encrypt takes a key and a byte slice and performs a AES encryption on the slcie
func (v *Vault) Encrypt(key string, vault []byte) ([]byte, error) {

	aeskey, _ := v.hash(key)

	block, err := aes.NewCipher(aeskey[:16])
	if err != nil {
		return nil, err
	}
	ciphertext := make([]byte, aes.BlockSize+len(vault))

	iv := ciphertext[:aes.BlockSize]

	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	stream := cipher.NewCFBEncrypter(block, iv)

	stream.XORKeyStream(ciphertext[aes.BlockSize:], vault)

	return ciphertext, err

}

// Read reads the content from the vault file
func (v Vault) Read() ([]byte, error) {
	return ioutil.ReadFile(v.Path)
}

// Write wrties the byte slice to in the vault file
func (v Vault) Write(data []byte) error {

	f, err := os.OpenFile(v.Path, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return cmd_errors.IOFileError{
			MSG: `😱 Unable to open file '.sherlocked'. This file should have been created
				by in the make command. If it is missing execute 'touch $HOME/.sherlocked'.`,
		}
	}

	if _, err := f.Write(data); err != nil {
		return cmd_errors.IOFileError{
			MSG: `😅 could not write the changed vault to file. Don't worry we
			have a plan B - if the '.sherlocked' is corrupted execute
			'mv $HOME/.sherlocked-{some-unix-time} $HOME/.sherlocked'`,
		}
	}
	return nil

}

// TmpBackup takes a backup of the current file
// then executes the passed function. If function returns nil
// the backup file gets deleted
func (v Vault) TmpBackup() (func() error, error) {
	// take a backup of the current vault file
	backup, err := renameFile(v.Path)
	if err != nil {
		return nil, err
	}

	// return clean-up function to remove tmp backup
	return func() error {
		return removeFile(backup)
	}, nil
}

// EvaluatePassword evaluates the strength of a password (range between 0-100)
func (v Vault) EvaluatePassword(password string) int {
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

// IsInit controls if there is already a vault file created
func (v Vault) IsInit() (bool, error) {
	fi, err := os.Stat(v.Path)
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

// Init initializes the vault file for the first time
func (v Vault) Init(clIO IO) error {
	fmt.Printf("Looks like your not yet set-up. Set a vault password to decrypt\nencrpyt your vault\n")
	var vault = make(map[string]string)

	password, err := clIO.SetNewPassword(v.EvaluatePassword)
	if err != nil {
		return err
	}
	// write changed vault
	b, err := v.Serialize(vault)
	if err != nil {
		return err
	}
	encrypted, err := v.Encrypt(password, b)
	if err != nil {
		return err
	}
	if err := v.Write(encrypted); err != nil {
		return err
	}
	fmt.Printf("✌🏼 You're all set!\n")
	return nil
}

// GetPath returns the path of the vault file
func (v Vault) GetPath() string {
	return v.Path
}

// hash hashes the vault key
func (v Vault) hash(key string) ([]byte, error) {
	b := sha256.Sum256([]byte(key))
	hexB := hex.EncodeToString(b[:])
	return []byte(hexB), nil
}

// Serialize marshals the vault to a byte slice in order for it be written to a
// file
func (v Vault) Serialize(data map[string]string) ([]byte, error) {
	b, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return b, nil
}

// deserialize takes the vault byte slice and unmarshales it to a map[string]string
func (v Vault) deserialize(vault []byte) (map[string]string, error) {

	var decrypted map[string]string
	if err := json.Unmarshal(vault, &decrypted); err != nil {
		return nil, fmt.Errorf("🧐 Ups looks like your password does not work for this vault!")
	}
	return decrypted, nil
}
