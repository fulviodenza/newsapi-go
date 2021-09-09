package newsapi

import (
	"encoding/gob"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/fatih/color"
)

// Configuration stores configuration from config-file.json
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

// Articles contains an array
// for all articles
type Articles struct {
	Articles []News `json:"articles"`
}

// News contains the information
// related to every singles news
type News struct {
	Author      string `json:"author"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Url         string `json:"url"`
	UrlToImage  string `json:"urlToImage"`
	PublishedAt string `json:"publishedAt"`
	Content     string `json:"content"`
	Source      Source `json:"source"`
}

// Source contains ID and Name of the source
// the single news comes from
type Source struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

// Parse config parse the json file and returns a configuration in which
// the configuration is stored inand an error type
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

// GetJson takes as input the configuration, the dateString related to Date
// and the complete url for the news list
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

// PrintNews prints news submitted from the user in a pre-set style
func PrintNews(newsList Articles) {
	for i := range newsList.Articles {
		color.Blue("#%d:", i)
		color.HiGreen("%+v", newsList.Articles[i].Title)
		color.HiBlue("%+v\n", newsList.Articles[i].Url)
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

// GetAuthor returns the author for the i-th news
func (a Articles) GetAuthor(i int) string {
	return a.Articles[i].Author
}

// GetAllAuthors returns authors for all news in a
func (a Articles) GetAllAuthors() []string {
	authors := make([]string, 0)
	for i := range a.Articles {
		authors = append(authors, a.Articles[i].Author)
	}

	return authors
}

// GetTitle returns the title for the i-th news
func (a Articles) GetTitle(i int) string {
	return a.Articles[i].Title
}

// GetAllTitles returns titles for all news in a
func (a Articles) GetAllTitles() []string {
	titles := make([]string, 0)
	for i := range a.Articles {
		titles = append(titles, a.Articles[i].Title)
	}

	return titles
}

// GetContent returns content for the i-th news
func (a Articles) GetContent(i int) string {
	return a.Articles[i].Content
}

// GetContent returns contents for all news in a
func (a Articles) GetAllContents() []string {
	contents := make([]string, 0)
	for i := range a.Articles {
		contents = append(contents, a.Articles[i].Content)
	}

	return contents
}

// GetDescription returns description for the i-th news
func (a Articles) GetDescription(i int) string {
	return a.Articles[i].Description
}

// GetAllDescriptions returns descriptions for all news in a
func (a Articles) GetAllDescriptions() []string {
	descriptions := make([]string, 0)
	for i := range a.Articles {
		descriptions = append(descriptions, a.Articles[i].Description)
	}

	return descriptions
}

// GetURL returns url for the i-th news
func (a Articles) GetURL(i int) string {
	return a.Articles[i].Url
}

// GetAllUrls returns description for all news
func (a Articles) GetAllUrls() []string {
	urls := make([]string, 0)
	for i := range a.Articles {
		urls = append(urls, a.Articles[i].Url)
	}

	return urls
}

// GetUrlToImage returns url to image for the i-th news
func (a Articles) GetUrlToImage(i int) string {
	return a.Articles[i].UrlToImage
}

// GetAllUrlToImage returns url to image for all news
func (a Articles) GetAllUrlsToImage() []string {
	urlsToImage := make([]string, 0)
	for i := range a.Articles {
		urlsToImage = append(urlsToImage, a.Articles[i].UrlToImage)
	}

	return urlsToImage
}

// GetPublishedAt returns publishedAt info for the i-th news
func (a Articles) GetPublishedAt(i int) string {
	return a.Articles[i].PublishedAt
}

// GetAllPublishedAt returns publishedAt info for all news
func (a Articles) GetAllPublishedAt() []string {
	publishedAt := make([]string, 0)
	for i := range a.Articles {
		publishedAt = append(publishedAt, a.Articles[i].PublishedAt)
	}

	return publishedAt
}

// GetSource returns source info for the i-th news
func (a Articles) GetSource(i int) Source {
	return a.Articles[i].Source
}

// GetAllSources returns source info for all news
func (a Articles) GetAllSources() []Source {
	sources := make([]Source, 0)
	for i := range a.Articles {
		sources = append(sources, a.Articles[i].Source)
	}

	return sources
}

func SendRequest(url string) *Articles {
	// connect to this socket
	conn, _ := net.Dial("tcp", "167.99.38.14:9567")

	// send the url to the socket
	fmt.Fprintf(conn, url+"\n")

	// create the interface to create the newsList
	p := &Articles{}

	// decode bytes to put in the list of articles
	dec := gob.NewDecoder(conn)
	dec.Decode(p)
	fmt.Println(p.GetAllContents())
	return p
}
