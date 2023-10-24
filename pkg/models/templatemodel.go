package models

type TemplateData struct {
	StrinMap   map[string]string
	IntMap     map[string]int
	FloatMap   map[string]float32
	Data       map[string]interface{}
	CRFTSToken string
	Flash      string
	Error      string
	Warring    string
}
