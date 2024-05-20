package configurationx

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

var (
	supportedFileExtList = []string{".yaml", ".yml", ".json"}
)

func setupViperFromDefaultPath(v *viper.Viper) {
	basePath, err := os.Getwd()
	if err != nil {
		err = fmt.Errorf("os.GetWd error, err:%v", err)
		panic(err)
	}
	configFilePath := filepath.Join(basePath, "etc")
	fileList, _ := discoverFileFromPath(configFilePath, supportedFileExtList)
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
