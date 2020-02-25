package config

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"os"
	"strings"
)

func InitConfig(configDir string) error {
	viper.SetConfigType("yaml")
	viper.SetConfigName("production")
	viper.AddConfigPath(configDir)
	if err := viper.ReadInConfig(); err != nil {
		return errors.Wrap(err, "can't load config file: production")
	}

	if env := os.Getenv("ENV"); len(env) != 0 {
		viper.SetConfigName(env)
		if err := viper.MergeInConfig(); err != nil {
			return errors.Wrap(err, fmt.Sprintf("can't load config file: %s", configDir))
		}
	}
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	return nil
}
