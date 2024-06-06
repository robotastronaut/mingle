package mpackage

import "github.com/spf13/viper"

func NewPackageViper(confPath string) *viper.Viper {
	v := viper.New()
	v.SetConfigName("mfile")
	v.SetConfigType("json")
	v.AddConfigPath(confPath)
	return v
}
