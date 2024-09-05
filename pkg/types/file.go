package types

import (
	"go/build"
	"strings"
)

type GoFile struct {
	Package         string
	Path            string
	ImportPath      string
	GlobalConstants []*GoType
	GlobalVariables []*GoType
	Structs         []*GoStruct
	Interfaces      []*GoInterface
	Imports         []*GoImport
	StructMethods   []*GoStructMethod
}

func isInGoPackages(path string) bool {
	goPath := strings.Replace(build.Default.GOPATH, "\\", "/", -1)
	return strings.Contains(path, goPath)
}
