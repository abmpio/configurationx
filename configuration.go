package configurationx

import (
	"errors"
	"log/slog"

	"github.com/abmpio/configurationx/options"
	"github.com/spf13/viper"
)

var (
	_instance *Configuration
)

type Configuration struct {
	Logger        slog.Logger
	viper         *viper.Viper
	ConfigManager ConfigManager

	BaseConsulPathList []string `json:"-"`
	EtcSubPath         string   `json:"-"`
	options.Options
}

type ConfigurationReadOption func(c *Configuration)

// 注册额外的配置信息
func RegistExtraProperties(key string, value interface{}) {
	options.RegistExtraProperties(key, value)
}

// 获取配置值
func GetOption(key string) interface{} {
	return GetInstance().Options.GetExtraProperties(key)
}

// 构建配置信息,将设置GetInstance()方法的返回值
func Use(c *Configuration, opts ...ConfigurationReadOption) (*Configuration, error) {
	for _, eachOpt := range opts {
		eachOpt(c)
	}

	_instance = c

	// reinitialize and read
	_instance.Options = options.NewOptions()
	err := _instance.ReadFrom(_instance.viper)
	if err != nil {
		return nil, err
	}
	return _instance, nil
}

// 获取实例
func GetInstance() *Configuration {
	return _instance
}

// Load出一个Configuration对象
func Load(etcSubPath string, opts ...ConfigurationReadOption) *Configuration {
	configuration := New()
	// configuration.Logger.Debug("准备初始化应用配置信息...")
	configuration.EtcSubPath = etcSubPath

	for _, eachOpt := range opts {
		if eachOpt == nil {
			continue
		}
		eachOpt(configuration)
	}

	//读取数据
	configuration.ReadFrom(configuration.viper)
	return configuration
}

func New() *Configuration {
	return NewConfiguration(viper.New())
}

func NewConfiguration(v *viper.Viper) *Configuration {
	if v == nil {
		panic(errors.New("viper参数不能为nil"))
	}
	configuration := new(Configuration)

	configuration.Logger = *slog.Default()
	configuration.viper = v
	configuration.BaseConsulPathList = []string{"abmpio/"}
	configuration.Options = options.NewOptions()
	return configuration
}

// Reset all configuration
func (c *Configuration) Reset() {
	c.viper = viper.New()
	(&c.Options).Reset()
}

// Merge others to c
func (c *Configuration) Merge(opts ...ConfigurationReadOption) *Configuration {
	for _, eachOpt := range opts {
		if eachOpt == nil {
			continue
		}
		eachOpt(c)
	}
	return c
}

func (c *Configuration) GetViper() *viper.Viper {
	return c.viper
}

// Merge from other Configuration
func (c *Configuration) MergeFrom(source *Configuration) *Configuration {
	c.viper.MergeConfigMap(source.viper.AllSettings())
	return c
}

func ReadFromConfiguration(source *Configuration) ConfigurationReadOption {
	return func(c *Configuration) {
		c.MergeFrom(source)
	}
}
