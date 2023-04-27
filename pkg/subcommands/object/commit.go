package object

import (
	"encoding/hex"
	"fmt"
	"github.com/jet/pkg/helper"
	"strings"
	"time"
)

type Commit struct {
	otype string
	odata []byte
	oid   []byte
}

func (t Commit) Odata() []byte {
	return t.odata
}

func (t Commit) Otype() string {
	return t.otype
}

func (t Commit) Oid() []byte {
	return t.oid
}

func NewCommit(author Author, msg string, treeOid []byte) Commit {
	contents := generateCommitConents(author, msg, hex.EncodeToString(treeOid))
	header := GenerateObjectHeader(helper.CommitType, len(contents))
	return Commit{
		odata: []byte(header + contents),
		otype: helper.CommitType,
		oid:   GenerateSHA1Hash(header + contents),
	}
}

func generateCommitConents(author Author, msg string, oidHexStr string) string {
	lines := []string{
		fmt.Sprintf("tree %s", oidHexStr),
		fmt.Sprintf("author %s", author.ToString()),
		fmt.Sprintf("committer %s", author.ToString()),
		"",
		msg,
	}
	return strings.Join(lines, "\n")
}

type Author struct {
	name  string
	email string
	ti    time.Time
}

func (a Author) ToString() string {
	formattedTime := a.ti.Format(time.RFC822)
	return fmt.Sprintf("%s <%s> %s", a.name, a.email, formattedTime)
}

func NewAuthor(name string, email string, ti time.Time) Author {
	return Author{
		name:  name,
		email: email,
		ti:    ti,
	}
}
