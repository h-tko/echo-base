package libraries

import (
	"fmt"
	"github.com/pelletier/go-toml"
	"os"
)

const CONF_DIR = "./config/"
const FILE_NAME = "app.toml"

type Config struct {
	app *toml.Tree
}

var sharedConfig *Config

func GetConfig() (*Config, error) {
	if sharedConfig == nil {
		sharedConfig = &Config{app: nil}

		defaultConf, err := toml.LoadFile(defaultFile())

		if err != nil {
			fmt.Printf("%#v", err)
			return nil, err
		}

		// errorしてたらdefaultのみ使う
		conf, _ := toml.LoadFile(fileByEnv())

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
