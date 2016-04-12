package main  

import (
	"log"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("config")
	viper.AddConfigPath("$GOPATH/src/github.com/crunchbase_data")
	viper.SetConfigType("json")
	err := viper.ReadInConfig()
	if err != nil {
		log.Println(err.Error())
	}


}