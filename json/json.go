package json

import "encoding/json"

type HillReq struct {
	Type      json.Number `json:"type"`
	InputText string      `json:"input_text"`
	Key       string      `json:"key"`
	M         json.Number `json:"m"`
	Encrypt   json.Number `json:"encrypt"`
	// File      []byte      `json:"file"`
}
