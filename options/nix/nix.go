package nix

import "encoding/json"

type NixOptions struct {
	Host       string `json:"host,omitempty"`
	Port       int32  `json:"port,omitempty"`
	AclToken   string `json:"aclToken,omitempty"`
	Disabled   bool   `json:"disabled,omitempty"`
	DefaultApp string `json:"defaultApp,omitempty"`
}

// new default nix options
// Disabled: true
func NewDefaultNixOptions() *NixOptions {
	options := &NixOptions{}
	options.Normalize()
	options.Disabled = true
	return options
}

func (o *NixOptions) Normalize() {
	if len(o.Host) <= 0 {
		o.Host = "127.0.0.1"
	}
	if o.Port <= 0 {
		o.Port = 9028
	}
}

// 序列化为json字符串
func (c *NixOptions) ToJsonString() []byte {
	jsonValue, _ := json.Marshal(c)
	return jsonValue
}
