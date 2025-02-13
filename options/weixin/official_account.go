package weixin

import (
	"encoding/json"

	"github.com/spf13/viper"
)

const (
	ConfigurationKey string = "weixin"
)

type OffiAccount struct {
	AppID         string `json:"appId" mapstructure:"appId"`
	AppSecret     string `json:"appSecret" mapstructure:"appSecret"`
	MessageToken  string `json:"messageToken" mapstructure:"messageToken"`
	MessageAesKey string `json:"messageAesKey" mapstructure:"messageAesKey"`
}

type WeixinConfiguration struct {
	// 公众号配置
	OffiAccount *OffiAccount `json:"offiAccount" mapstructure:"offiAccount"`
}

// 序列化为json字符串
func (c *WeixinConfiguration) ToJsonString() []byte {
	jsonValue, _ := json.Marshal(c)
	return jsonValue
}

// 从中读取配置
func ReadFrom(v *viper.Viper) (*WeixinConfiguration, error) {
	var weixin WeixinConfiguration
	err := v.UnmarshalKey(ConfigurationKey, &weixin)
	return &weixin, err
}
