package binding_struct

type HillReq struct {
	Type      int    `json:"type" form:"type"`
	InputText string `json:"input_text" form:"input_text"`
	Key       string `json:"key" form:"key"`
	M         int    `json:"m" form:"m"`
	Encrypt   int    `json:"encrypt" form:"encrypt"`
}

type PlayfairReq struct {
	Type      int    `json:"type" form:"type"`
	InputText string `json:"input_text" form:"input_text"`
	Key       string `json:"key" form:"key"`
	Encrypt   int    `json:"encrypt" form:"encrypt"`
}

type VigenereReq struct {
	Type      int    `json:"type" form:"type"`
	InputText string `json:"input_text" form:"input_text"`
	Key       string `json:"key" form:"key"`
	Encrypt   int    `json:"encrypt" form:"encrypt"`
}

type AffineReq struct {
	Type      int    `json:"type" form:"type"`
	InputText string `json:"input_text" form:"input_text"`
	M         int    `json:"m" form:"m"`
	B         int    `json:"b" form:"b"`
	Encrypt   int    `json:"encrypt" form:"encrypt"`
}
