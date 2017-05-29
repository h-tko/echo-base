package libraries

import (
	"fmt"
	"github.com/pelletier/go-toml"
	"os"
)

const confDir = "./config/"
const fileName = "app.toml"

// Config ...
//
// app.tomlに設定されているデータにアクセスするための構造体.
type Config struct {
	// app.tomlの保存変数
	app *toml.Tree
}

// singletonインスタンス.
var sharedConfig *Config

// GetConfig ...
//
// Config構造体のsingletonインスタンスを得る.
func GetConfig() (*Config, error) {
	if sharedConfig == nil {
		sharedConfig = &Config{app: nil}

		// まずデフォルト読み込み
		defaultConf, err := toml.LoadFile(defaultFile())

		if err != nil {
			fmt.Printf("%#v", err)
			return nil, err
		}

		// 環境別読み込み（errorしてたらdefaultのみ使う）
		conf, err := toml.LoadFile(fileByEnv())

		if err != nil {
			for _, key := range conf.Keys() {
				defaultConf.Set(key, conf.Get(key))
			}
		}

		sharedConfig.app = defaultConf
	}

	return sharedConfig, nil
}

// defaultConf ...
//
// デフォルトのapp.tomlのフルパス
func defaultFile() string {
	return confDir + fileName
}

// fileByEnv ...
//
// 環境別のapp.tomlのフルパス
func fileByEnv() string {
	return fmt.Sprintf(confDir+"%s/"+fileName, os.Getenv("APP_ENV"))
}

// Get ...
//
// configデータを取得し、interface{}として返却.
func (c *Config) Get(key string) interface{} {
	return c.app.Get(key)
}

// GetString ...
//
// configデータを取得し、stringとして返却.
func (c *Config) GetString(key string) string {
	return c.app.Get(key).(string)
}

// GetInt ...
//
// configデータを取得し、intとして返却.
func (c *Config) GetInt(key string) int {
	return c.app.Get(key).(int)
}

// GetBool ...
//
// configデータを取得し、boolとして返却.
func (c *Config) GetBool(key string) bool {
	return c.app.Get(key).(bool)
}
