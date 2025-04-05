package configurationx

import "github.com/spf13/viper"

// read a key to string value
func (c *Configuration) ReadString(key string) string {
	if c.viper == nil {
		return ""
	}
	return c.viper.GetString(key)
}

// read bool value
// if not exist,return false
func (c *Configuration) ReadBool(key string) bool {
	if c.viper == nil {
		return false
	}
	return c.viper.GetBool(key)
}

// read key value as any type
func (c *Configuration) Read(key string) any {
	if c.viper == nil {
		return false
	}
	return c.viper.Get(key)
}

// read key value as int value
func (c *Configuration) ReadInt(key string) int {
	if c.viper == nil {
		return 0
	}
	return c.viper.GetInt(key)
}

// read key value as int value
func (c *Configuration) ReadInt32(key string) int32 {
	if c.viper == nil {
		return 0
	}
	return c.viper.GetInt32(key)
}

// read key value as int value
func (c *Configuration) ReadInt64(key string) int64 {
	if c.viper == nil {
		return 0
	}
	return c.viper.GetInt64(key)
}

// 反序列化指定的key的值到一个对象中
func (c *Configuration) UnmarshFromKey(key string, v interface{}, opts ...viper.DecoderConfigOption) error {
	return c.viper.UnmarshalKey(key, v, opts...)
}
