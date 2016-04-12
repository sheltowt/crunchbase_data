package data_load

import(
)

type DataLoad struct{
	MongoConnection string
	CrunchBaseKey string
}

func NewDataLoad(mongoConnection string, crunchBaseKey string) &DataLoad {
	return &DataLoad{mongoConnection, crunchBaseKey}
}

func(dataLoad DataLoad) PullFromCrunchbase() {
	
}