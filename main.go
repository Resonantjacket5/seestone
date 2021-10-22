package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/urfave/cli/v2"
)

type book struct {
	ID   string `json:"_id"`
	Name string `json:"name"`
}

type listBookRespBody struct {
	Docs []book `json:"docs"`
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

func main() {

	var defaultFlags *[]cli.Flag = &[]cli.Flag{
		&cli.StringFlag{
			Name:  "api-token",
			Value: "None",
			Usage: "the-one-api api token",
		},
	}

	var app *cli.App = &cli.App{
		Name:  "seestone",
		Usage: "The seeing stone",
		Action: func(c *cli.Context) error {
			fmt.Println("no subcommand called")
			return nil
		},
		Commands: []*cli.Command{
			{
				Flags: *defaultFlags,
				Name:  "books",
				Action: func(c *cli.Context) error {
					listBooks()
					return nil
				},
			},
			{
				Flags: *defaultFlags,
				Name:  "movies",
				Action: func(c *cli.Context) error {
					listMovies()
					return nil
				},
			},
			{
				Flags: *defaultFlags,
				Name:  "characters",
				Action: func(c *cli.Context) error {
					listCharacters(c)
					return nil
				},
			},
		},
	}
	// (&cli.App{}).Run(os.Args)
	err := app.Run(os.Args)
	if err != nil {
		log.Fatalln(err)
	}
}

func listBooks() {
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
	fmt.Println(respBody.Docs)
}

func listMovies() {

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

type responseBody interface {
	Documents() []interface{}
}

func listCharacters(c *cli.Context) {
	fmt.Println("attempt to list characters")
	var url = "https://the-one-api.dev/v2/character"
	var respBody listCharacterRespBody
	var response *http.Response = genericGet(c, url, respBody)

	defer response.Body.Close()
	err := json.NewDecoder(response.Body).Decode(&respBody)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(respBody.Documents())
}

func genericGet(c *cli.Context, url string, respBody responseBody) *http.Response {

	// var header string = "Authorization: Bearer "

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
