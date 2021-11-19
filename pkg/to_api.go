package api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/urfave/cli/v2"
)

type Book struct {
	ID   string `json:"_id"`
	Name string `json:"name"`
}

type listBookRespBody struct {
	Docs []Book `json:"docs"`
}

// ListBooks returns books.
func ListBooks() []Book {
	var url = "https://the-one-api.dev/v2/book"
	resp, err := http.Get(url)

	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	// var data map[string]interface{}

	var respBody listBookRespBody

	// if err := json.Unmarshal(resp.Body, &data); err != nill {
	err = json.NewDecoder(resp.Body).Decode(&respBody)

	if err != nil {
		log.Fatalln(err)
	}
	return respBody.Docs
}

type character struct {
	ID    string `json:"_id"`
	Name  string `json:"name"`
	Race  string `json:"race"`
	Realm string `json:"realm"`
}

type listCharacterRespBody struct {
	Docs []character `json:"docs"`
}

// func (v Vertex) Abs() float64 {

func (l listCharacterRespBody) Documents() []interface{} {
	// convert from list of struct to generic interfaces
	interfaces := make([]interface{}, len(l.Docs))
	for i, v := range l.Docs {
		interfaces[i] = v
	}
	return interfaces
}

type movie struct {
	ID               string `json:"_id"`
	Name             string `json:"name"`
	RuntimeInMinutes int    `json:"runtimeInMinutes"`
	BudgetInMillions int    `json:"budgetInMillions"`
}

type listMovieRespBody struct {
	Docs []movie `json:"docs"`
}

func ListMovies() {

	var url = "https://the-one-api.dev/v2/movie"
	resp, err := http.Get(url)
	validateResp(resp)

	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	var respBody listMovieRespBody
	err = json.NewDecoder(resp.Body).Decode(&respBody)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(respBody.Docs)
}

// ListCharacters returns list of characters.
func ListCharacters(c *cli.Context) {
	fmt.Println("attempt to list characters")
	var url = "https://the-one-api.dev/v2/character"
	var respBody listCharacterRespBody
	var response *http.Response = genericGet(c, url)

	defer response.Body.Close()
	err := json.NewDecoder(response.Body).Decode(&respBody)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(respBody.Documents())
}

type chapter struct {
	ID          string `json:"_id"`
	ChapterName string `json:"chapterName"`
	BookID      string `json:"book"`
}

type listChapterRespBody struct {
	Docs []chapter `json:"docs"`
}

func (l listChapterRespBody) Names() []string {
	var names []string = []string{}
	for _, v := range l.Docs {
		names = append(names, v.ChapterName)
	}
	return names
}

// ListChapters returns list of chapters
func ListChapters(c *cli.Context) {
	var url = "https://the-one-api.dev/v2/chapter"
	var respBody listChapterRespBody
	var response *http.Response = genericGet(c, url)

	defer response.Body.Close()
	err := json.NewDecoder(response.Body).Decode(&respBody)
	if err != nil {
		log.Fatalln(err)
	}

	var names []string = respBody.Names()

	for _, v := range names {
		fmt.Println(v)
	}
}

type Quote struct {
	ID        string `json:"_id"`
	Dialog    string `json:"dialog"`
	MovieID   string `json:"movie"`
	Character string `json:"character"`
}

type listQuoteRespBody struct {
	Docs []Quote `json:"docs"`
}

// ListQuotes returns list of quotess
func ListQuotes(c *cli.Context) []Quote {
	var url = "https://the-one-api.dev/v2/quote"
	var respBody listQuoteRespBody
	var response *http.Response = genericGet(c, url)

	defer response.Body.Close()
	err := json.NewDecoder(response.Body).Decode(&respBody)
	if err != nil {
		log.Fatalln(err)
	}

	var quotes []Quote = respBody.Docs

	return quotes
}

type responseBody interface {
	Documents() []interface{}
}

// genericGet
// c cli.Context, used for authentication token
// url string, the url to call get on
func genericGet(c *cli.Context, url string) *http.Response {

	var client *http.Client = &http.Client{
		Timeout: time.Second * 20,
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalln(err)
	}

	var apiToken string = c.String("api-token")
	if apiToken != "None" {
		var bearerString string = fmt.Sprintf("Bearer %s", apiToken)
		req.Header.Add("Authorization", bearerString)
	}

	resp, err := client.Do(req)

	// resp, err := http.Get(url)
	//replace with validation function
	fmt.Println(resp.Status)
	if err != nil {
		log.Fatalln(err)
	}

	return resp
}

func validateResp(resp *http.Response) {
	if resp.StatusCode != 200 {
		log.Fatalln(resp.Status)
	}
}

func respToString(resp http.Response) string {
	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	return string(bytes)
}
