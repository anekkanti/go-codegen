package types

type GoInterface struct {
	File     *GoFile
	Name     string
	Comments string
	Methods  []*GoMethod
}
