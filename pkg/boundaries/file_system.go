package boundaries

import (
	"os"
)

type FileSystem interface {
	Read(file string) (string, error)
	Write(file string, content string) error
	MakeDirs(path string) error
	ListFiles(path string) ([]os.DirEntry, error)
}

type OsFileSystem struct{}

func (fs *OsFileSystem) Read(file string) (string, error) {
	contents, err := os.ReadFile(file)
	if err != nil {
		return "", err
	}

	return string(contents), nil
}

func (fs *OsFileSystem) Write(file string, content string) error {
	return os.WriteFile(file, []byte(content), 0o644)
}

func (fs *OsFileSystem) MakeDirs(path string) error {
	return os.MkdirAll(path, os.ModePerm)
}

func (fs *OsFileSystem) ListFiles(path string) ([]os.DirEntry, error) {
	return os.ReadDir(path)
}
