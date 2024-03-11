package conf

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/techoc/fanqie-novel-api/pkg/global"
)

func Setup() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("conf")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	initGlobal()

}

func initGlobal() {
	viper.AutomaticEnv()

	// 初始化http服务配置
	global.ServerConf.HttpPort = viper.GetInt("server.HttpPort")
	global.ServerConf.ReadTimeout = viper.GetDuration("server.ReadTimeout")
	global.ServerConf.WriteTimeout = viper.GetDuration("server.WriteTimeout")
	global.ServerConf.RunMode = viper.GetString("server.RunMode")

	// 设置数据库默认配置
	viper.SetDefault("database.Type", "sqlite")
	viper.SetDefault("database.Name", "novel")
	viper.SetDefault("database.TablePrefix", "novel")

	// 初始化数据库配置
	global.DatabaseConf.Host = viper.GetString("database.Host")
	global.DatabaseConf.Name = viper.GetString("database.Name")
	global.DatabaseConf.Password = viper.GetString("database.Password")
	global.DatabaseConf.TablePrefix = viper.GetString("database.TablePrefix")
	global.DatabaseConf.Type = viper.GetString("database.Type")
	global.DatabaseConf.User = viper.GetString("database.User")

	// 从环境变量中获取数据库配置
	global.DatabaseConf.DSN = viper.GetString("DSN")
	global.DatabaseConf.Type = viper.GetString("DATABASE_TYPE")
}
