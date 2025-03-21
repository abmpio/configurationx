package weixin

import (
	"encoding/json"
	"fmt"
)

// 企业微信配置
type Work struct {
	// 是否启用企业微信
	Enabled bool `json:"enabled" mapstructure:"enabled"`

	CorpID          string `json:"corpId" mapstructure:"corpId"`
	AgentID         int    `json:"agentId" mapstructure:"agentId"`
	Secret          string `json:"secret" mapstructure:"secret"`
	MessageToken    string `json:"messageToken" mapstructure:"messageToken"`
	MessageAesKey   string `json:"messageAesKey" mapstructure:"messageAesKey"`
	MessageCallback string `json:"messageCallback" mapstructure:"messageCallback"`
	OAuthCallback   string `json:"oauthCallback" mapstructure:"oauthCallback"`

	// contacts app
	Contacts *ContactsAppOptions `json:"contacts" mapstructure:"contacts"`
}

// contacts app
type ContactsAppOptions struct {
	Secret   string `json:"secret" mapstructure:"secret"`
	Token    string `json:"token" mapstructure:"token"`
	AESKey   string `json:"aesKey" mapstructure:"aesKey"`
	Callback string `json:"callback" mapstructure:"callback"`
}

func (o *Work) Validate() error {
	if !o.Enabled {
		// disabled, return
		return nil
	}
	if o.CorpID == "" {
		return fmt.Errorf("corpId参数不能为空")
	}
	return nil
}

// 序列化为json字符串
func (c *Work) ToJsonString() []byte {
	jsonValue, _ := json.Marshal(c)
	return jsonValue
}
