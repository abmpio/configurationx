package consulv

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"strings"

	"github.com/abmpio/configurationx"
	"github.com/abmpio/configurationx/options/consul"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
)

func ReadFromConsul(consulOptions consul.ConsulOptions, consulPathList []string) *configurationx.Configuration {
	c := configurationx.New()
	if len(consulPathList) <= 0 {
		return c
	}
	c.BaseConsulPathList = consulPathList
	if len(consulOptions.Host) <= 0 {
		c.Logger.Info("没有配置好consul.host参数,不从consul中读取配置")
		return c
	}

	if consulOptions.Disabled {
		c.Logger.Info("consul.disabled参数值为true,将不从consul中读取配置")
		return c
	}
	endpoint := fmt.Sprintf("%s:%d", consulOptions.Host, consulOptions.Port)
	err := initConfigManager(c, []string{endpoint})
	if err != nil {
		return c
	}

	childKeyList, err := getChildKvPairs(c, c.BaseConsulPathList, true)
	if err != nil {
		panic(err)
	}
	if len(childKeyList) <= 0 {
		//没有子节点，则直接返回
		return c
	}

	for _, eachChildKey := range childKeyList {
		allSettings, err := getDataFromConfigManager(c, c.ConfigManager, eachChildKey)
		if err != nil {
			continue
		}
		//合并配置
		c.GetViper().MergeConfigMap(allSettings)
	}
	return c
}

func getChildKvPairs(c *configurationx.Configuration, keyList []string, containSubPath bool) (childKeyList []string, err error) {
	if len(keyList) <= 0 {
		return childKeyList, nil
	}
	pathList := make([]string, 0)
	for _, eachKey := range keyList {
		//增加path后缀
		if strings.HasSuffix(eachKey, "/") {
			pathList = append(pathList, eachKey)
		} else {
			pathList = append(pathList, eachKey+"/")
		}
	}
	for _, eachKey := range pathList {
		thisChildKvPairs, err := c.ConfigManager.List(eachKey)
		if err != nil {
			c.Logger.Error(fmt.Sprintf("获取%s的子key时出现异常,详细异常信息:%s", eachKey, err.Error()))
			return childKeyList, err
		}
		if len(thisChildKvPairs) <= 0 {
			continue
		}
		for _, eachChildKey := range thisChildKvPairs {
			if eachChildKey == nil {
				continue
			}
			if eachChildKey.Key == eachKey {
				//自身，直接循环
				continue
			}
			if strings.HasSuffix(eachChildKey.Key, "/") && containSubPath {
				//递归获取子节点
				subChildKeyList, err := getChildKvPairs(c, []string{eachChildKey.Key}, containSubPath)
				if err != nil {
					return childKeyList, err
				}
				if len(subChildKeyList) <= 0 {
					continue
				}
				childKeyList = append(childKeyList, subChildKeyList...)
			} else {
				childKeyList = append(childKeyList, eachChildKey.Key)
			}
		}
	}
	return childKeyList, nil
}

func initConfigManager(c *configurationx.Configuration, endPoint []string) error {
	configManager, err := NewStandardConsulConfigManager(endPoint)
	if err != nil {
		c.Logger.Error(fmt.Sprintf("连接到consul时出现异常,err:%s", err.Error()))
		return err
	}
	c.ConfigManager = configManager
	return nil
}

func getDataFromConfigManager(c *configurationx.Configuration, cm configurationx.ConfigManager, path string) (map[string]interface{}, error) {
	b, err := cm.Get(path)
	if err != nil {
		return nil, err
	}
	// 检测字符串格式
	format := detectFormat(b)
	if format == "Unknown" {
		err = fmt.Errorf("unknow data format,path:%s", path)
		c.Logger.Error(err.Error())
		return nil, err
	}
	v := viper.New()
	v.SetConfigType(format)

	// 尝试加载字符串
	err = v.ReadConfig(strings.NewReader(string(b)))
	if err != nil {
		c.Logger.Error("error reading config from string:", err)
		return nil, err
	}
	// result := make(map[string]interface{})
	// err = json.Unmarshal(b, &result)
	// if err != nil {
	// 	c.Logger.Error("cann't unmarshal map from data, path: %s,err:%s", path,err.Error())
	// 	return nil, err
	// }
	return v.AllSettings(), err
}

// check string format（JSON or YAML）
func detectFormat(input []byte) string {
	if isValidJSON(input) {
		return "JSON"
	}

	if isValidYAML(input) {
		return "YAML"
	}

	return "Unknown"
}

// check string is  JSON format
func isValidJSON(input []byte) bool {
	var js json.RawMessage
	return json.Unmarshal(input, &js) == nil
}

// check string is YAML format
func isValidYAML(input []byte) bool {
	var y interface{}
	err := yaml.Unmarshal(input, &y)
	if err != nil {
		slog.Default().Warn(fmt.Sprintf("无效的yaml格式,err:%s", err.Error()))
	}
	return err == nil
}
