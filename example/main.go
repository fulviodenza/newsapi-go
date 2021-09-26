package main

import (
	"fmt"
	"log"
	"time"

	"github.com/fulviodenza/newsapi-go/newsapi"
)

func main() {

	config, err := newsapi.ParseConfig("../config-file.json")
	if err != nil {
		log.Fatal(err)
	}
	var date time.Time
	layout := "2006-01-02"

	if config.Date == "today" {
		date = time.Now()
	}

	dateString := date.Format(layout)
	fmt.Println(dateString)

	// compose the url, send the request and print news get
	completeUrl := newsapi.ComposeURL(config, dateString)

	newsList, err := newsapi.SendRequest(completeUrl)
	if err != nil {
		log.Panic(err)
	}

	newsapi.PrintNews(*newsList)

}
