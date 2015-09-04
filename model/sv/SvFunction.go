package sv

type SvFunction struct {
	Name      string          `json:"name"`
	Arguments []SvDeclaration `json:"arguments"`
}
