package aliyun

import (
	"encoding/json"

	"github.com/spf13/viper"
)

const (
	AliasName_Default string = "default"
	ConfigurationKey  string = "aliyun"
)

type AliyunConfiguration struct {
	// sms列表
	SmsList map[string]*AliyunSmsOptions `mapstructure:"smsList" json:"smsList" yaml:"smsList"`
}

// new default configuration
func NewDefaultConfiguration() *AliyunConfiguration {
	options := &AliyunConfiguration{}
	options.SmsList = make(map[string]*AliyunSmsOptions)
	return options
}

// 获取默认的sms配置
func (c *AliyunConfiguration) GetDefaultSmsOptions() *AliyunSmsOptions {
	result := c.GetSmsOptions("")
	if result == nil {
		result = c.GetSmsOptions(AliasName_Default)
	}
	return result
}

// 获取指定别名的项
func (c *AliyunConfiguration) GetSmsOptions(aliasName string) *AliyunSmsOptions {
	if len(c.SmsList) <= 0 {
		return nil
	}
	item, ok := c.SmsList[aliasName]
	if !ok {
		return nil
	}
	return item
}

// 序列化为json字符串
func (c *AliyunConfiguration) ToJsonString() []byte {
	jsonValue, _ := json.Marshal(c)
	return jsonValue
}

// 从中读取配置
func ReadFrom(v *viper.Viper) (AliyunConfiguration, error) {
	var aliyunConfiguration AliyunConfiguration
	err := v.UnmarshalKey(ConfigurationKey, &aliyunConfiguration)
	return aliyunConfiguration, err
}
