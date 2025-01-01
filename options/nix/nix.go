package nix

type NixOptions struct {
	Host       string `json:"host,omitempty"`
	Port       int32  `json:"port,omitempty"`
	AclToken   string `json:"aclToken,omitempty"`
	Disabled   bool   `json:"disabled,omitempty"`
	DefaultApp string `json:"defaultApp,omitempty"`
}

func (o *NixOptions) Normalize() {
	if len(o.Host) <= 0 {
		o.Host = "127.0.0.1"
	}
	if o.Port <= 0 {
		o.Port = 9028
	}
}
