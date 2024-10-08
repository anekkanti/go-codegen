package types

import "strings"

type GoImport struct {
	File *GoFile
	Name string
	Path string
}

// For an import - guess what prefix will be used
// in type declarations.  For examples:
//
//	"strings" -> "strings"
//	"net/http/httptest" -> "httptest"
//
// Libraries where the package name does not match
// will be mis-identified.
func (g *GoImport) Prefix() string {
	if g.Name != "" {
		return g.Name
	}

	path := strings.Trim(g.Path, "\"")
	lastSlash := strings.LastIndex(path, "/")
	if lastSlash == -1 {
		return path
	}

	return path[lastSlash+1:]
}
