package consul

import "os"

var _options ConsulOptions

const (
	ConfigurationKey string = "consul"
)

// 获取consul配置
func GetConsulOptions() *ConsulOptions {
	return &_options
}

func SetConsul(options *ConsulOptions) {
	if options != nil {
		_options = *options
		(&_options).Normalize()
		if len(_options.AclToken) > 0 {
			os.Setenv("CONSUL_HTTP_TOKEN", _options.AclToken)
		}
	}
}
