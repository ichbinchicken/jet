package object

import (
	"github.com/jet/pkg/helper"
)

type Blob struct {
	fname string
	data  []byte
	otype string
	oid   []byte // SHA-1 bytes
}

func (o Blob) Name() string {
	return o.fname
}

func (o Blob) Odata() []byte {
	return o.data
}

func (o Blob) Otype() string {
	return o.otype
}

func (o Blob) Oid() []byte {
	return o.oid
}

func NewBlob(contents string, fname string) Blob {
	header := GenerateObjectHeader(helper.BlobType, len(contents))
	return Blob{
		fname: fname,
		data:  []byte(header + contents),
		otype: helper.BlobType,
		oid:   GenerateSHA1Hash(header + contents),
	}
}
