package types

type GoMethod struct {
	Name     string
	Params   []*GoType
	Comments string
	Results  []*GoType
}

type GoStructMethod struct {
	GoMethod
	Receivers []string
}
