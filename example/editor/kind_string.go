// Code generated by "stringer -type Kind lex.go"; DO NOT EDIT

package main

import "fmt"

const _Kind_name = "kindOpkindStringkindSlashkindQuestkindRelkindCommakindDotkindEofkindColonkindSemikindHashkindErrkindRegexpkindRegexpBackkindByteOffsetkindLineOffsetkindCmdkindArg"

var _Kind_index = [...]uint8{0, 6, 16, 25, 34, 41, 50, 57, 64, 73, 81, 89, 96, 106, 120, 134, 148, 155, 162}

func (i Kind) String() string {
	if i < 0 || i >= Kind(len(_Kind_index)-1) {
		return fmt.Sprintf("Kind(%d)", i)
	}
	return _Kind_name[_Kind_index[i]:_Kind_index[i+1]]
}
