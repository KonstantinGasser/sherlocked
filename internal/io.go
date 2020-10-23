package internal

// import (
// 	"bufio"
// 	"fmt"
// 	"io/ioutil"
// 	"log"
// 	"os"
// 	"strings"
// 	"syscall"

// 	"golang.org/x/crypto/ssh/terminal"
// )

// func openFile(path string) (*os.File, error) {
// 	pwd, err := os.UserHomeDir()
// 	filepath := []string{pwd, path}
// 	if err != nil {
// 		return nil, err
// 	}
// 	f, err := os.OpenFile(strings.Join(filepath, "/"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return f, nil
// }

// func FileToString(path string) (string, error) {

// 	pwd, err := os.UserHomeDir()
// 	filepath := []string{pwd, path}
// 	if err != nil {
// 		return "", err
// 	}

// 	content, err := ioutil.ReadFile(strings.Join(filepath, "/"))
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	text := string(content)
// 	return text, nil
// }

// func WriteCipher(vault []byte) error {
// 	f, err := openFile(vaultfile)
// 	if err != nil {
// 		return err
// 	}
// 	defer f.Close()
// 	if _, err := f.Write(vault); err != nil {
// 		return err
// 	}
// 	return nil
// }

// func RemoveFile(path string) error {
// 	pwd, _ := os.UserHomeDir()
// 	filepath := []string{pwd, path}
// 	err := os.Remove(strings.Join(filepath, "/"))
// 	return err
// }

// func InputPassword() (string, error) {

// 	fmt.Print("Password: ")
// 	bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
// 	if err != nil {
// 		return "", err
// 	}
// 	password := string(bytePassword)
// 	fmt.Print("\n")
// 	return strings.TrimSpace(password), nil
// }

// func InputCredentials() (string, string, error) {
// 	reader := bufio.NewReader(os.Stdin)

// 	fmt.Print("Enter Username: ")
// 	username, _ := reader.ReadString('\n')

// 	fmt.Print("Enter Password: ")
// 	bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
// 	if err != nil {
// 		return "", "", err
// 	}
// 	password := string(bytePassword)

// 	return strings.TrimSpace(username), strings.TrimSpace(password), nil
// }
