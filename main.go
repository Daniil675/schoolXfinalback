package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"schoolXfinalback/server"
	"schoolXfinalback/storage/datastore"
	"sync"

	//"./server"
	//"./storage/datastore"
	//"./telegram"
	//_ "./utility"
	"github.com/sirupsen/logrus"
	"github.com/upper/db/v4/adapter/mysql"
)

//const (
//	listeningPort = "8080"
//)
type Config struct {
	//Token    string `json:"token"`
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Database string `json:"database"`
}

func main() {
	logrus.Info("Wake up, Neo...")

	config := getConfig()
	settings := mysql.ConnectionURL{
		User:     config.User,
		Password: config.Password,
		Host:     config.Host,
		Database: config.Database,
	}
	datastore.New(settings)
	//go analytics.Start()

	var wg sync.WaitGroup
	//
	wg.Add(1)
	s := &server.Server{}
	go s.Start("80")

	wg.Wait()
}

func getConfig() Config {
	// Open our jsonFile
	jsonFile, err := os.Open("resources/config.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		logrus.Error("Error in getConfig.", err)
	}

	fmt.Println("Successfully Opened config.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we initialize our Users array
	var config Config

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	err = json.Unmarshal(byteValue, &config)
	if err != nil {
		logrus.Error("Error in getConfig.", err)
	}
	return config
}
