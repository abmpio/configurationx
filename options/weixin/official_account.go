package weixin

import (
	"encoding/json"
	"fmt"

	"github.com/go-viper/mapstructure/v2"
	"github.com/spf13/viper"
)

const (
	ConfigurationKey string = "weixin"
)

// 公众号配置
type OffiAccount struct {
	// 是否启用公众号配置
	Enabled bool `json:"enabled" mapstructure:"enabled"`

	AppID         string `json:"appId" mapstructure:"appId"`
	AppSecret     string `json:"appSecret" mapstructure:"appSecret"`
	MessageToken  string `json:"messageToken" mapstructure:"messageToken"`
	MessageAesKey string `json:"messageAesKey" mapstructure:"messageAesKey"`
}

func (o *OffiAccount) Validate() error {
	if !o.Enabled {
		return nil
	}

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
	err := v.UnmarshalKey(ConfigurationKey, &weixin, func(dc *mapstructure.DecoderConfig) {
		dc.TagName = "json"
	})
	return &weixin, err
}
