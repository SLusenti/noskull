package parser

import (
	"regexp"
)

const (
	COMMANDS_FST = []string{
		"use",
		"cluster",
		"select",
		"append",
		"prepend",
		"insert",
		"update",
		"delete",
		"drop",
		"list",
		"show",
		"conf",
		"func",
		"create",
		"help",
	}
	COMMANDS_MID = []string{
		"filter",
		"append",
		"prepend",
		"insert",
		"update",
		"delete",
		"drop",
		"show",
		"count",
		"return",
		"collect",
		"merge"
		"limit"
	}
)

const (
	TYPE_COMMAND int = iota
	TYPE_ARG
	TYPE_START_FUNC
	TYPE_END_FUNC
	TYPE_START_STR
	TYPE_END_STR
	TYPE_START_JSON
	TYOE_END_JSON
)

type Token struct {
	Value string
	Args  []string
	Type  int
}

func Tokenize(s string) []*Token {
	r := regexp.Compile("[[:print:]]")
	tvalue := ""
	for _, char := range s {
		if r.MatchString(string(char)) {
			tvalue += char
		}
	}
	tok := &Token{
		Value: tvalue,
	}
}
