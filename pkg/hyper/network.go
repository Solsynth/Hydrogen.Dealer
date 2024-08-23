package hyper

import jsoniter "github.com/json-iterator/go"

type NetworkPackage struct {
	Action   string `json:"w"`
	Endpoint string `json:"e,omitempty"`
	Message  string `json:"m"`
	Payload  any    `json:"p"`
}

func (v NetworkPackage) Marshal() []byte {
	data, _ := jsoniter.Marshal(v)
	return data
}

func (v NetworkPackage) RawPayload() []byte {
	out, _ := jsoniter.Marshal(v.Payload)
	return out
}
