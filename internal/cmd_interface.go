package internal

// PasswordManager is the interface to act with the vault
type PasswordManager interface {
	// required pre-condition: is false new vault needs be created
	Init(clIO IO) error
	IsInit() (bool, error)
	// wrapper to read and write to the file
	Read() (content []byte, err error)
	Write(data []byte) error
	Serialize(data map[string]string) (marshaled []byte, err error)
	// wrapper function to perform a backup of the vault
	// befor proceeding with a function (like writing a new vault)
	TmpBackup() (cleanup func() error, err error)
	EvaluatePassword(password string) (strength int)
	Encrypt(key string, vault []byte) (encrypted []byte, err error)
	Decrypt(key string, vault []byte) (decrypted map[string]string, err error)
	GetPath() string
}

// NewPasswordManager returns a pointer to a Vault implementing
// the PasswordManager interface
func NewPasswordManager(path string) PasswordManager {
	return &Vault{
		Path: path,
	}
}

type IO interface {
	Credentials() (uname string, pass string, err error)
	Password() (string, error)
	SimpleText(txt string) (string, error)
	SetNewPassword(eval func(pass string) int) (newpassword string, err error)
}

// NewIO returns a pointer to a cmdIO implemnting the IO interface
func NewIO() IO {
	return &cmdIO{}
}
