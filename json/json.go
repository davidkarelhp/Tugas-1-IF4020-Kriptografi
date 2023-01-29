package json

import "encoding/json"

type HillReq struct {
	PlaintextType string      `json:"plaintext_type"`
	POrCText      string      `json:"p_or_c_text"`
	Key           string      `json:"key"`
	M             json.Number `json:"m"`
	Encrypt       json.Number `json:"encrypt"`
}
