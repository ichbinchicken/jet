package Refs

import (
	"os"
)

type Refs struct{}

func (refs *Refs) UpdateHEAD(path string, commitHexStr string) {
	_, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	err = os.WriteFile(path, []byte(commitHexStr), 0o644)
	if err != nil {
		panic(err)
	}
}

func (refs *Refs) ReadHEAD(path string) (string, bool) {
	content, err := os.ReadFile(path)
	// Here, we only care about that file not existed.
	// So we only want to pop up this single error.
	// However, we might probably hide other types of error.
	// If there is an error, ReadHEAD will return false.
	// If there is no error, ReadHEAD will return true.
	if err != nil {
		return "", false
	}
	return string(content), true
}
