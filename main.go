package main  

import (
	"log"
	"github.com/spf13/viper"
	"github.com/crunchbase_data/data_load"
)

func main() {
	viper.SetConfigName("config")
	viper.AddConfigPath("$GOPATH/src/github.com/crunchbase_data")
	viper.SetConfigType("json")
	err := viper.ReadInConfig()
	if err != nil {
		log.Println(err.Error())
	}

	dataLoad := NewDataLoad(viper.Get("crunchbase.api_user_token"), viper.Get("crunchbase.mongo_lab"))

	dataLoad.
}