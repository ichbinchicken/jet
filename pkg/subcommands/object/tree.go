package object

import (
	"fmt"
	"github.com/jet/pkg/helper"
)

const (
	MODE = "100644"
)

type Tree struct {
	otype string
	odata []byte
	oid   []byte
}

func (t Tree) Odata() []byte {
	return t.odata
}

func (t Tree) Otype() string {
	return t.otype
}

func (t Tree) Oid() []byte {
	return t.oid
}

type entry struct {
	name string
	oid  []byte
}

func NewTree(blobs helper.GenericSlice[Blob]) Tree {
	entries := helper.MapSlice(blobs, func(b Blob) entry {
		return entry{
			name: b.Name(),
			oid:  b.Oid(),
		}
	})
	contents := generateTreeContents(entries)
	header := GenerateObjectHeader(helper.TreeType, len(contents))
	odata := append([]byte(header), contents...)
	return Tree{
		odata: odata,
		otype: helper.TreeType,
		oid:   GenerateSHA1Hash(string(odata)),
	}
}

func generateTreeContents(entries []entry) []byte {
	var result []byte
	for _, en := range entries {
		prefix := fmt.Sprintf("%s %s\x00", MODE, en.name)
		result = append(result, []byte(prefix)...)
		result = append(result, en.oid...)
	}
	return result
}
