package internal

// PasswordManager is the interface to act with the vault
type PasswordManager interface {
	IsInit() (bool, error)
	Read() ([]byte, error)
	Write(data []byte) error
	Backup(after func() error) (func() error, error)
	Serialize(data map[string]string) ([]byte, error)
	EvaluatePassword(password string) int
	Encrypt(key string, vault []byte) ([]byte, error)
	Decrypt(key string, vault []byte) (map[string]string, error)
	GetPath() string
}

func NewPasswordManager(path string) PasswordManager {
	return &Vault{
		Path: path,
	}
}

type IO interface {
	Credentials() (string, string, error)
	Password() (string, error)
	SimpleText(txt string) (string, error)
}

func NewIO() IO {
	return &cmdIO{}
}
