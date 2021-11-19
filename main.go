package main

import (
	"fmt"
	"log"
	"os"

	api "github.com/resonantjacket5/seestone/pkg"

	"github.com/urfave/cli/v2"
)

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
					var books []api.Book = api.ListBooks()
					fmt.Println(books)
					return nil
				},
			},
			{
				Flags: *defaultFlags,
				Name:  "movies",
				Action: func(c *cli.Context) error {
					api.ListMovies()
					return nil
				},
			},
			{
				Flags: *defaultFlags,
				Name:  "characters",
				Action: func(c *cli.Context) error {
					api.ListCharacters(c)
					return nil
				},
				Subcommands: []*cli.Command{
					{
						Flags: *defaultFlags,
						Name:  "describe",
						Action: func(c *cli.Context) error {
							fmt.Printf("Hello %q", c.Args().Get(0))
							return nil
						},
					},
				},
			},
			{
				Flags: *defaultFlags,
				Name:  "chapters",
				Action: func(c *cli.Context) error {
					api.ListChapters(c)
					return nil
				},
			},
			{
				Flags: *defaultFlags,
				Name:  "quotes",
				Action: func(c *cli.Context) error {
					var quotes []api.Quote = api.ListQuotes(c)
					for _, v := range quotes {
						fmt.Println(v.Character)
					}
					// fmt.Println(quotes)
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
