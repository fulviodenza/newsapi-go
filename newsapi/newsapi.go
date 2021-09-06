package newsapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/fatih/color"
)

type configuration struct {
	Argument                    string `json:"argument"`
	Date                        string `json:"date"`
	News_number                 int    `json:"news_number"`
	Sorted_by                   string `json:"sorted_by"`
	Country                     string `json:"country"`
	Category                    string `json:"category"`
	Everything_or_top_headlines string `json:"everything_or_top_headlines"`
	ApiKey                      string `json:"apiKey"`
	Language                    string `json:"language"`
}

type Articles struct {
	Articles []News `json:"articles"`
}
type News struct {
	Author      string `json:"author"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func ParseConfig(filename string) (configuration, error) {
	configFile, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer configFile.Close()

	byteValue, _ := ioutil.ReadAll(configFile)
	config := configuration{}
	if err := json.Unmarshal(byteValue, &config); err != nil {
		log.Fatal(err)
	}
	return config, err
}

// Return the Json file
func GetJson(config configuration, dateString string, url string) Articles {

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	var newsList Articles
	err = json.NewDecoder(resp.Body).Decode(&newsList)

	if err != nil {
		log.Fatal(err)
	}

	return newsList
}

// Print news submitted from the user in a pre-set style
func PrintNews(newsList Articles) {
	for i := range newsList.Articles {
		color.Blue("#%d:", i)
		color.HiGreen("%+v\n\n", newsList.Articles[i])
	}
	fmt.Println("EOF")
}

// Compose the complete url base on the given configuration
func ComposeURL(config configuration, dateString string) string {
	url := "https://newsapi.org/v2/"
	urlEverythingOrTopHeadlines := config.Everything_or_top_headlines + "?"
	urlQuery := "q=" + config.Argument
	urlDate := "&from=" + dateString
	urlSort := "&sortBy=" + config.Sorted_by
	urlApi := "&apiKey=" + config.ApiKey
	urlLanguage := "&language=" + config.Language

	completeUrl := url + urlEverythingOrTopHeadlines + urlQuery + urlDate + urlSort + urlLanguage + urlApi

	return completeUrl
}
