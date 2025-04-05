package configurationx

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

var (
	supportedFileExtList = []string{
		fmt.Sprintf(".%s", ConfigType_Yaml),
		fmt.Sprintf(".%s", ConfigType_Yml),
		fmt.Sprintf(".%s", ConfigType_Json),
	}
)

const (
	ConfigType_Yaml string = "yaml"
	ConfigType_Yml  string = "yml"
	ConfigType_Json string = "json"
)

func setupViperFromDefaultPath(v *viper.Viper) {
	basePath, err := os.Getwd()
	if err != nil {
		err = fmt.Errorf("os.GetWd error, err:%v", err)
		panic(err)
	}
	configFilePath := filepath.Join(basePath, "etc")
	SetupViperFromPath(v, configFilePath)
}

// setup viper from full filepath
func SetupViperFromFilePath(v *viper.Viper, fullFilePath string) error {
	v.SetConfigFile(fullFilePath)
	err := v.ReadInConfig()
	return err
}

// 从一段字符串中读取
func SetupViperFromString(v *viper.Viper, configValue string, configType string) error {
	v.SetConfigType(configType) // 必须指定类型
	err := v.ReadConfig(strings.NewReader(configValue))
	return err
}

// setup viper from specified filepath and specified ext names
// only for .yaml,.yml,.json
func SetupViperFromPath(v *viper.Viper, configFilePath string) {
	SetupViperFromPathAndFileExt(v, configFilePath, supportedFileExtList)
}

// setup viper from specified filepath and specified ext names
func SetupViperFromPathAndFileExt(v *viper.Viper, configFilePath string, fileExtList []string) {
	if len(fileExtList) <= 0 {
		return
	}
	fileList, _ := discoverFileFromPath(configFilePath, fileExtList)
	if len(fileList) <= 0 {
		// empty folder
		return
	}
	for _, eachFile := range fileList {
		fileName := filepath.Base(eachFile)
		i := strings.LastIndex(fileName, ".")
		if i == -1 {
			continue
		}
		configName := fileName[:i]
		configType := fileName[i+1:]
		if configName == "" || configType == "" {
			continue
		}
		v.SetConfigType(configType)
		v.SetConfigName(configName)
		v.AddConfigPath(configFilePath)
		err := v.ReadInConfig()
		if err != nil {
			err = fmt.Errorf("读取配置文件时出现异常,文件名:%s,异常信息:%s", eachFile, err.Error())
			panic(err)
		}
	}
}

func (c *Configuration) readFromDefaultPath() {
	defaultViper := viper.New()
	setupViperFromDefaultPath(defaultViper)
	//合并
	c.viper.MergeConfigMap(defaultViper.AllSettings())
}

// ConfigurationReadOption that read from ./etc path, belown file type is searched,
// 1. *.yml
// 2. *.yaml
// 3. *.json
func ReadFromDefaultPath() ConfigurationReadOption {
	return func(c *Configuration) {
		c.readFromDefaultPath()
	}
}
