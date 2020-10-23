package cmd_errors

type IOFileError struct {
	MSG string
}

func (err IOFileError) Error() string {
	return err.MSG
}

type OSStdInError struct {
	MSG string
}

func (err OSStdInError) Error() string {
	return err.MSG
}

type ZeroVaultError struct {
	MSG string
}

func (err ZeroVaultError) Error() string {
	return err.MSG
}

type MapConversionError struct {
	MSG string
}

func (err MapConversionError) Error() string {
	return err.MSG
}

type InitNotDoneError struct {
	MSG string
}

func (err InitNotDoneError) Error() string {
	return err.MSG
}
