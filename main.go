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

func main() {
	var app *cli.App = &cli.App{
		Name:  "books",
		Usage: "list books",
		Action: func(c *cli.Context) error {
			fmt.Println("main")
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

func respToString(resp http.Response) string {
	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	return string(bytes)
}
