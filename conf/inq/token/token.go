// generated by gocc; DO NOT EDIT.

package token

import (
	"fmt"
)

type Token struct {
	Type
	Lit []byte
	Pos
}

type Type int

const (
	INVALID Type = iota
	EOF
)

type Pos struct {
	Offset int
	Line   int
	Column int
}

func (this Pos) String() string {
	return fmt.Sprintf("Pos(offset=%d, line=%d, column=%d)", this.Offset, this.Line, this.Column)
}

type TokenMap struct {
	typeMap []string
	idMap   map[string]Type
}

func (this TokenMap) Id(tok Type) string {
	if int(tok) < len(this.typeMap) {
		return this.typeMap[tok]
	}
	return "unknown"
}

func (this TokenMap) Type(tok string) Type {
	if typ, exist := this.idMap[tok]; exist {
		return typ
	}
	return INVALID
}

func (this TokenMap) TokenString(tok *Token) string {
	//TODO: refactor to print pos & token string properly
	return fmt.Sprintf("%s(%d,%s)", this.Id(tok.Type), tok.Type, tok.Lit)
}

func (this TokenMap) StringType(typ Type) string {
	return fmt.Sprintf("%s(%d)", this.Id(typ), typ)
}

var TokMap = TokenMap{
	typeMap: []string{
		"INVALID",
		"$",
		"check",
		"service",
		"name",
		"host",
		"expose",
		",",
		"restart",
		"reload",
		"alert",
		"with",
		":",
		"(",
		")",
		"if",
		"operator",
		"then",
		"for",
		"cycles",
	},

	idMap: map[string]Type{
		"INVALID":  0,
		"$":        1,
		"check":    2,
		"service":  3,
		"name":     4,
		"host":     5,
		"expose":   6,
		",":        7,
		"restart":  8,
		"reload":   9,
		"alert":    10,
		"with":     11,
		":":        12,
		"(":        13,
		")":        14,
		"if":       15,
		"operator": 16,
		"then":     17,
		"for":      18,
		"cycles":   19,
	},
}
