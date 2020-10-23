package internal

import (
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/KonstantinGasser/sherlocked/cmd_errors"
)

func openFile(path string) (*os.File, error) {

	f, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, cmd_errors.IOFileError{
			MSG: `😱 Unable to open file '.sherlocked'. This file should have been created
				by in the make command. If it is missing execute 'touch $HOME/.sherlocked'.`,
		}
	}

	return f, nil
}

func renameFile(from string) (string, error) {
	home, _ := os.UserHomeDir()
	backupname := strconv.FormatInt(time.Now().UnixNano(), 10)
	backuppath := strings.Join([]string{home, ".sherlocked-" + backupname}, "/")

	return backuppath, os.Rename(from, backuppath)
}

func removeFile(path string) error {
	if err := os.Remove(path); err != nil {
		return cmd_errors.IOFileError{
			MSG: `🧐 the backup file of the vault could not be removed.
			If this is bordering you delete '$HOME/.sherlocked-{some-unix-time}'`,
		}
	}
	return nil
}

func readFile(path string) ([]byte, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, cmd_errors.IOFileError{
			MSG: `🧐 Could not read from the vault file. Verify that '.sherlocked'
			exists under $HOME.`,
		}
	}
	return content, nil
}

func writeFile(path string, vault []byte) error {

	f, err := openFile(path)
	if err != nil {
		return err
	}

	if _, err := f.Write(vault); err != nil {
		return cmd_errors.IOFileError{
			MSG: `😅 could not write the changed vault to file. Don't worry we
			have a plan B - if the '.sherlocked' is corrupted execute
			'mv $HOME/.sherlocked-{some-unix-time} .sherlocked'`,
		}
	}
	return nil
}
