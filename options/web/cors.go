package web

type CorsMode string

const (
	CorsMode_AllowAll  = "allow-all"
	CorsMode_Whitelist = "whitelist"
)

type CORS struct {
	Mode      CorsMode        `mapstructure:"mode" json:"mode" yaml:"mode"`
	Whitelist []CORSWhitelist `mapstructure:"whitelist" json:"whitelist" yaml:"whitelist"`

	AllowedMethods []string `mapstructure:"allowedMethods" json:"allowedMethods" yaml:"allowedMethods"`
	AllowedHeaders []string `mapstructure:"allowedHeaders" json:"allowedHeaders" yaml:"allowedHeaders"`
	ExposedHeaders []string `mapstructure:"exposedHeaders" json:"exposedHeaders" yaml:"exposedHeaders"`
	MaxAge         *int     `mapstructure:"maxAge" json:"maxAge" yaml:"maxAge"`
}

type CORSWhitelist struct {
	AllowedOrigins string `mapstructure:"allow-origin" json:"allow-origin" yaml:"allow-origin"`
}

func (c *CORS) GetAllowedOrigins() []string {
	if len(c.Whitelist) <= 0 {
		return make([]string, 0)
	}
	list := make([]string, 0)
	for _, eachWhitelist := range c.Whitelist {
		list = append(list, eachWhitelist.AllowedOrigins)
	}
	return list
}
