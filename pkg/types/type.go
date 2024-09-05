package types

type GoType struct {
	Name       string
	Type       string
	Underlying string
	Inner      []*GoType
}
