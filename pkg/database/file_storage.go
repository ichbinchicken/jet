package database

import (
	"encoding/hex"
	"github.com/jet/pkg/helper"
	"github.com/jet/pkg/subcommands/object"
	"os"
	"path/filepath"
)

type FileStorage interface {
	Read(file string) (string, error)
	Write(file string, content string) error
	MakeDirs(path string) error
	ListFiles(path string) ([]os.DirEntry, error)
	// git(jet) specific operations:
	WriteJetObject(objectsPath string, obj object.Object) error
}

type OsFileStorage struct{}

func (fs *OsFileStorage) WriteJetObject(objectsPath string, obj object.Object) error {
	// create new dir
	hexStr := hex.EncodeToString(obj.Oid())
	blobFileParentPath := filepath.Join(objectsPath, hexStr[:2])
	err := fs.MakeDirs(blobFileParentPath)
	if err != nil {
		return err
	}

	// create a temp file under blobFileParentPath
	tmpFile, err := os.CreateTemp(blobFileParentPath, "jetTempBlob-*")
	if err != nil {
		return err
	}
	defer os.Remove(tmpFile.Name())

	// compress the content
	compressed, err := helper.Compress(obj.Odata())
	if err != nil {
		return err
	}

	// write compressed content into the temp file
	_, err = tmpFile.Write(compressed)
	if err != nil {
		return err
	}

	// rename
	blobFilePath := filepath.Join(blobFileParentPath, hexStr[2:])
	err = os.Rename(tmpFile.Name(), blobFilePath)
	return err // err can be nil
}

func (fs *OsFileStorage) Read(file string) (string, error) {
	contents, err := os.ReadFile(file)
	if err != nil {
		return "", err
	}

	return string(contents), nil
}

func (fs *OsFileStorage) Write(file string, content string) error {
	return os.WriteFile(file, []byte(content), 0o644)
}

func (fs *OsFileStorage) MakeDirs(path string) error {
	return os.MkdirAll(path, os.ModePerm)
}

func (fs *OsFileStorage) ListFiles(path string) ([]os.DirEntry, error) {
	return os.ReadDir(path)
}
