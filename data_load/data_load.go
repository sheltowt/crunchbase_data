package data_load

import(
	"log"
	"strconv"
	"strings"
	"encoding/json"
	"io/ioutil"
	"net/http"
	mgo"gopkg.in/mgo.v2"
)

type DataLoad struct{
	MongoConnection *mgo.Collection
	CrunchBaseKey string
}

type CrunchBaseResponse struct{
	MetaData interface{} `json:"metadata"`
	Data interface{} `json:"data"`
	Items []interface{} `json:"items"`
}

func NewDataLoad(mongoConnection string, crunchBaseKey string) *DataLoad {
	session, _ := mgo.Dial(mongoConnection)
	c := session.DB("crunchbase_data").C("practice")
	return &DataLoad{c, crunchBaseKey}
}

func(dataLoad DataLoad) PullFromCrunchbase() (err error) {

	for i := 1; i <= 100; i++ {

		components := []string{"https://api.crunchbase.com/v/3/organizations?page=", strconv.Itoa(i),"&user_key=", dataLoad.CrunchBaseKey}
		url := strings.Join(components, "")

		resp, err := http.Get(url)
		if err != nil {
			return err
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		var dataResponse CrunchBaseResponse
		err = json.Unmarshal(body, &dataResponse)
		if err != nil {
			log.Println(err.Error())
			return err
		}

		for _, doc := range dataResponse.Items {
			log.Println(doc)
			dataLoad.MongoConnection.Insert(doc)
		}
	}

	return nil
}