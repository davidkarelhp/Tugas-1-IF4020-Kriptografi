package json

type HillReq struct {
	Type      int    `json:"type" form:"type"`
	InputText string `json:"input_text" form:"input_text"`
	Key       string `json:"key" form:"key"`
	M         int    `json:"m" form:"m"`
	Encrypt   int    `json:"encrypt" form:"encrypt"`
	// File      []byte      `json:"file"`
}
