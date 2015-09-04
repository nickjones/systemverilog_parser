package sv

type SvFile struct {
	FileName string   `json:"filename"`
	Code     []SvCode `json:"code"`
}
