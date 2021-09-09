package main

import (
	"fmt"
	"time"

	"feedme-go/newsapi"
)

func main() {

	config, _ := newsapi.ParseConfig("./config-file.json")

	var date time.Time
	layout := "2006-01-02"

	if config.Date == "today" {
		date = time.Now()
	}

	dateString := date.Format(layout)
	fmt.Println(dateString)

	// compose the url, send the request and print news get
	completeUrl := newsapi.ComposeURL(config, dateString)
	newsList := newsapi.SendRequest(completeUrl)
	newsapi.PrintNews(*newsList)
}
