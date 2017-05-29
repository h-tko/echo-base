package libraries

import (
	"fmt"
	"github.com/pelletier/go-toml"
	"os"
)

const confDir = "./config/"
const fileName = "app.toml"

// Config ...
//   configへのアクセス用構造体
type Config struct {
	// app.tomlの保存変数
	app *toml.Tree
}

// singletonインスタンス
var sharedConfig *Config

// GetConfig ...
//   Configをロードして構造体に詰めて返却
//
//   scope: public
//
//   return: *Config
//
//   return: error
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

// デフォルトのapp.tomlのフルパス
//
// scope: private
//
// return: string
func defaultFile() string {
	return confDir + fileName
}

// 環境別のapp.tomlのフルパス
//
// scope: private
//
// return: string
func fileByEnv() string {
	return fmt.Sprintf(confDir+"%s/"+fileName, os.Getenv("APP_ENV"))
}

// Get ...
//   configデータ取得
//
//   scope: public
//
//   param: key string キー
//
//   return: interface{} 値
func (c *Config) Get(key string) interface{} {
	return c.app.Get(key)
}

// GetString ...
//   configデータ取得
//
//   scope: public
//
//   param: key string キー
//
//   return: string 値
func (c *Config) GetString(key string) string {
	return c.app.Get(key).(string)
}

// GetInt ...
//   configデータ取得
//
//   scope: public
//
//   param: key string キー
//
//   return: int 値
func (c *Config) GetInt(key string) int {
	return c.app.Get(key).(int)
}

// GetBool ...
//   configデータ取得
//
//   scope: public
//
//   param: key string キー
//
//   return: bool 値
func (c *Config) GetBool(key string) bool {
	return c.app.Get(key).(bool)
}
