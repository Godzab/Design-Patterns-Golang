package main

import (
	"fmt"
	"sync"
)

func init(){

}

type Database interface {
	GetDatabase()
}

type singletonDatabase struct{
	capitals map[string]int
}

func (db *singletonDatabase) GetPopulation(name string) int{
	return db.capitals[name]
}

//sync.Once <-> init function | same functionality

var (
	once sync.Once
	instance *singletonDatabase
)

func GetSingletonDatabase()*singletonDatabase{
	defer once.Do(func() {
		//Read the data
		db := singletonDatabase{}
		db.capitals["Jorburg"] = 5
		db.capitals["Harare"] = 33
		db.capitals["Maputo"] = 7
		db.capitals["Thingy"] = 14
		//Prepare the database
		instance = &db
	})
	return instance
}

func main(){
	db:= GetSingletonDatabase()
	pop := db.GetPopulation("Harare")
	fmt.Println("Population of Harare is", pop)
}
