package weixin

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/viper"
)

const (
	ConfigurationKey string = "weixin"
)

// 公众号配置
type OffiAccount struct {
	AppID         string `json:"appId" mapstructure:"appId"`
	AppSecret     string `json:"appSecret" mapstructure:"appSecret"`
	MessageToken  string `json:"messageToken" mapstructure:"messageToken"`
	MessageAesKey string `json:"messageAesKey" mapstructure:"messageAesKey"`
}

func (o *OffiAccount) Validate() error {
	if o.AppID == "" {
		return fmt.Errorf("appId参数不能为空")
	}
	if o.AppSecret == "" {
		return fmt.Errorf("appSecret参数不能为空")
	}
	return nil
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
