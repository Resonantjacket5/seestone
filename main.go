package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

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
	var app *cli.App = &cli.App{
		Name:  "seestone",
		Usage: "The seeing stone",
		Action: func(c *cli.Context) error {
			fmt.Println("no subcommand called")
			return nil
		},
		Commands: []*cli.Command{
			{
				Name: "books",
				Action: func(c *cli.Context) error {
					listBooks()
					return nil
				},
			},
			{
				Name: "movies",
				Action: func(c *cli.Context) error {
					listMovies()
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
