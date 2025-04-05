package configurationx

import (
	"testing"

	"github.com/go-viper/mapstructure/v2"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestSetupViperFromFilePath(t *testing.T) {
	configValue := `
app:
  abmpconsul:
    path: abmpio_dev
`
	t.Run("config should has value", func(t *testing.T) {
		v := viper.New()
		err := SetupViperFromString(v, configValue, ConfigType_Yaml)

		abmpconsulPathValue := v.GetString("app.abmpconsul.path")
		assert.Equal(t, "abmpio_dev", abmpconsulPathValue)
		assert.Equal(t, nil, err)
	})

	t.Run("config unmarshal should has value", func(t *testing.T) {
		v := viper.New()
		err := SetupViperFromString(v, configValue, ConfigType_Yaml)
		assert.Equal(t, nil, err)

		var app struct {
			AbmpConsul struct {
				Path string `json:"path"`
			} `json:"abmpconsul"`
		}
		err = v.UnmarshalKey("app", &app, func(dc *mapstructure.DecoderConfig) {
			dc.TagName = "json"
		})
		assert.Equal(t, "abmpio_dev", app.AbmpConsul.Path)
		assert.Equal(t, nil, err)
	})
}
