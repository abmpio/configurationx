package weixin

type WeixinConfiguration struct {
	// 公众号配置
	OffiAccount *OffiAccount `json:"offiAccount" mapstructure:"offiAccount"`
	Work        *Work        `json:"work" mapstructure:"work"`
}
