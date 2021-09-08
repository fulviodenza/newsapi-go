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
	Author      string `json:"author"`      //done
	Title       string `json:"title"`       //done
	Description string `json:"description"` //done
	Url         string `json:"url"`         //done
	UrlToImage  string `json:"urlToImage"`  //done
	PublishedAt string `json:"publishedAt"` //done
	Content     string `json:"content"`     //done
	Source      Source `json:"source"`      //done
}

type Source struct {
	Id   string `json:"id"`
	Name string `json:"name"`
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
	urlApi := "&apiKey=" + os.Getenv("APIKEY")
	urlLanguage := "&language=" + config.Language

	completeUrl := url + urlEverythingOrTopHeadlines + urlQuery + urlDate + urlSort + urlLanguage + urlApi

	return completeUrl
}

//GETTER Functions

func (a Articles) GetAuthor(i int) string {
	return a.Articles[i].Author
}

func (a Articles) GetAllAuthors() []string {
	authors := make([]string, 0)
	for i := range a.Articles {
		authors = append(authors, a.Articles[i].Author)
	}

	return authors
}

func (a Articles) GetTitle(i int) string {
	return a.Articles[i].Title
}

func (a Articles) GetAllTitles(i int) []string {
	titles := make([]string, 0)
	for i := range a.Articles {
		titles = append(titles, a.Articles[i].Title)
	}

	return titles
}

func (a Articles) GetContent(i int) string {
	return a.Articles[i].Content
}

func (a Articles) GetAllContents() []string {
	contents := make([]string, 0)
	for i := range a.Articles {
		contents = append(contents, a.Articles[i].Content)
	}

	return contents
}

func (a Articles) GetDescription(i int) string {
	return a.Articles[i].Description
}

func (a Articles) GetAllDescriptions() []string {
	descriptions := make([]string, 0)
	for i := range a.Articles {
		descriptions = append(descriptions, a.Articles[i].Description)
	}

	return descriptions
}

func (a Articles) GetURL(i int) string {
	return a.Articles[i].Url
}

func (a Articles) GetAllUrls() []string {
	urls := make([]string, 0)
	for i := range a.Articles {
		urls = append(urls, a.Articles[i].Url)
	}

	return urls
}

func (a Articles) GetUrlToImage(i int) string {
	return a.Articles[i].UrlToImage
}

func (a Articles) GetAllUrlsToImage() []string {
	urlsToImage := make([]string, 0)
	for i := range a.Articles {
		urlsToImage = append(urlsToImage, a.Articles[i].UrlToImage)
	}

	return urlsToImage
}

func (a Articles) GetPublishedAt(i int) string {
	return a.Articles[i].PublishedAt
}

func (a Articles) GetAllPublishedAt() []string {
	publishedAt := make([]string, 0)
	for i := range a.Articles {
		publishedAt = append(publishedAt, a.Articles[i].PublishedAt)
	}

	return publishedAt
}

func (a Articles) GetSource(i int) Source {
	return a.Articles[i].Source
}

func (a Articles) GetAllSources() []Source {
	sources := make([]Source, 0)
	for i := range a.Articles {
		sources = append(sources, a.Articles[i].Source)
	}

	return sources
}
