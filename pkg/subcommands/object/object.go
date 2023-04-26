package object

import (
	"crypto/sha1"
	"io"
	"strconv"
)

type Object interface {
	Oid() []byte
	Otype() string
	// Odata always contains header
	Odata() []byte
}

func GenerateSHA1Hash(input string) []byte {
	hash := sha1.New()
	_, err := io.WriteString(hash, input)
	if err != nil {
		panic(err)
	}
	return hash.Sum(nil)
}

func GenerateObjectHeader(t string, l int) string {
	return t + " " + strconv.Itoa(l) + "\x00"
}
