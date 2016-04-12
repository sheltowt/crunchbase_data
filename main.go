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

	dataLoad := data_load.NewDataLoad(viper.GetString("mongo_lab.connection_string"), viper.GetString("crunchbase.api_user_token"))

	err = dataLoad.PullFromCrunchbase()
	if err != nil {
		log.Println(err.Error())
	} else {
		log.Println("executed flawlessly")
	}
}