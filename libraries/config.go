package libraries

import (
	"fmt"
	"github.com/pelletier/go-toml"
	"os"
)

const CONF_DIR = "../config/"
const FILE_NAME = "app.toml"

type Config struct {
	app *toml.TomlTree
}

var sharedConfig *Config

func GetConfig() (*Config, error) {
	if sharedConfig == nil {
		sharedConfig = &Config{app: nil}

		defaultConf, err := toml.Load(defaultFile())

		if err != nil {
			return nil, err
		}

		conf, err := toml.Load(fileByEnv())

		if err != nil {
			return nil, err
		}

		for _, key := range conf.Keys() {
			defaultConf.Set(key, conf.Get(key))
		}

		sharedConfig.app = defaultConf
	}

	return sharedConfig, nil
}

func defaultFile() string {
	return CONF_DIR + FILE_NAME
}

func fileByEnv() string {
	return fmt.Sprintf(CONF_DIR+"%s/"+FILE_NAME, os.Getenv("APP_ENV"))
}

func (c *Config) Get(key string) interface{} {
	return c.app.Get(key)
}
