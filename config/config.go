package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

func init()  {
	path, osErr := os.Getwd()
	if osErr != nil {
		panic(osErr)
	}
	viper.SetConfigFile(path+"/config/config.toml")
	err := viper.ReadInConfig()
	if err != nil{
		panic(fmt.Errorf("viper error:%s",err))
	}
}

func GetString(key string) string {
	chkKey(key)
	return viper.GetString(key)
}

func GetInt(key string) int  {
	chkKey(key)
	return viper.GetInt(key)
}

func GetBool(key string) bool  {
	chkKey(key)
	return viper.GetBool(key)
}

func chkKey(key string)  {
	if !viper.IsSet(key){
		panic(fmt.Errorf("viper error:key %s no found",key))
	}
}