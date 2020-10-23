package internal

import (
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"
)

func openFile(path string) (*os.File, error) {

	f, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}

	return f, nil
}

func renameFile(from string) (string, error) {
	home, _ := os.UserHomeDir()
	backupname := strconv.FormatInt(time.Now().UnixNano(), 10)
	backuppath := strings.Join([]string{home, backupname}, "/")

	return backuppath, os.Rename(from, backuppath)
}

func removeFile(path string) error {
	if err := os.Remove(path); err != nil {
		return err
	}
	return nil
}

func readFile(path string) ([]byte, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return content, nil
}

func writeFile(path string, vault []byte) error {

	f, err := openFile(path)
	if err != nil {
		return err
	}

	if _, err := f.Write(vault); err != nil {
		return err
	}
	return nil
}
