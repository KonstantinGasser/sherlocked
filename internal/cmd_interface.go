package internal

// PasswordManager is the interface to act with the vault
type PasswordManager interface {
	// required pre-condition: is false new vault needs be created
	Init(clIO IO) error
	IsInit() (bool, error)
	// wrapper to read and write to the file
	Read() ([]byte, error)
	Write(data []byte) error
	Serialize(data map[string]string) ([]byte, error)
	// wrapper function to perform a backup of the vault
	// befor proceeding with a function (like writing a new vault)
	Backup(after func() error) (func() error, error)
	EvaluatePassword(password string) int
	Encrypt(key string, vault []byte) ([]byte, error)
	Decrypt(key string, vault []byte) (map[string]string, error)
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
	Credentials() (string, string, error)
	Password() (string, error)
	SimpleText(txt string) (string, error)
	SetNewPassword(eval func(pass string) int) (string, error)
}

// NewIO returns a pointer to a cmdIO implemnting the IO interface
func NewIO() IO {
	return &cmdIO{}
}
