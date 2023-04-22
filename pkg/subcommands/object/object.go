package object

import (
	"crypto/sha1"
	"encoding/hex"
	"io"
	"strconv"
)

type Type string

const (
	blobType Type = "blob"
	treeType Type = "tree"
)

type Object struct {
	objData []byte
	objType Type
	objId   string // SHA-1 hash string
}

func (o Object) ObjData() []byte {
	return o.objData
}

func (o Object) ObjType() Type {
	return o.objType
}

func (o Object) ObjId() string {
	return o.objId
}

func NewBlob(contents string) Object {
	header := string(blobType) + " " + strconv.Itoa(len(contents)) + "\x00"
	return Object{
		objData: []byte(header + contents),
		objType: blobType,
		objId:   GenerateSHA1Hash(header + contents),
	}
}

func GenerateSHA1Hash(input string) string {
	hash := sha1.New()
	_, err := io.WriteString(hash, input)
	if err != nil {
		panic(err)
	}
	return hex.EncodeToString(hash.Sum(nil))
}
