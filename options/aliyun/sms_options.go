package aliyun

type AliyunSmsOptions struct {
	// aksk
	AccessKeyId     string `json:"accessKeyId" mapstructure:"accessKeyId"`
	AccessKeySecret string `json:"accessKeySecret" mapstructure:"accessKeySecret"`
	// 默认的签名
	DefaultSign string `json:"defaultSign" mapstructure:"defaultSign"`

	// 不用于实际逻辑，仅仅只是用来表示这个aksk所对应的用户信息
	UserInfo string `json:"userInfo" mapstructure:"userInfo"`
}
