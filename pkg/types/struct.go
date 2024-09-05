package types

import (
	"reflect"
	"strings"
)

type GoStruct struct {
	File     *GoFile
	Name     string
	Comments string
	Fields   []*GoField
}

type GoField struct {
	Struct *GoStruct
	Name   string
	Type   string
	Tag    *GoTag
}

type GoTag struct {
	Field *GoField
	Value string
}

func (g *GoTag) Get(key string) string {
	tag := strings.ReplaceAll(g.Value, "`", "")
	return reflect.StructTag(tag).Get(key)
}
