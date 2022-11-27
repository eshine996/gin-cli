package global

import (
	"github.com/spf13/viper"
	"os"
	"strings"
)

const (
	DefaultConfigDir  = "configs"
	DefaultConfigType = "yaml"
)

func NewDefaultConfig() *ViperConfig {
	vp := viper.New()
	vp.SetConfigType(DefaultConfigType)
	vp.AddConfigPath(DefaultConfigDir)
	vp.AddConfigPath(".")

	return &ViperConfig{
		vp,
	}
}

type ViperConfig struct {
	*viper.Viper
}

func FormatEnvKey(s string) string {
	return strings.ToUpper(strings.Replace(s, ".", "_", -1))
}

func (c *ViperConfig) ReadData() error {
	return c.ReadInConfig()
}

func (c *ViperConfig) GetWithEnv(key string) interface{} {
	v := os.Getenv(FormatEnvKey(key))
	if len(v) > 0 {
		return v
	} else {
		return c.Get(key)
	}
}
