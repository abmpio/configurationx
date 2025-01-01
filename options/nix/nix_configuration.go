package nix

import "os"

var _options NixOptions

const (
	ConfigurationKey string = "nix"
)

// 获取nix配置
func GetNixOptions() *NixOptions {
	return &_options
}

func SetNix(options *NixOptions) {
	_options = *options
	(&_options).Normalize()
	if len(_options.AclToken) > 0 {
		os.Setenv("NIX_HTTP_TOKEN", _options.AclToken)
	}
}
