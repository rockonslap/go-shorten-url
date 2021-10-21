package main

import (
	"fmt"
	"log"
	"os"
	shortenurl "shorten-url/app/shorten_url"
	database "shorten-url/config"

	"github.com/urfave/cli/v2"
)

func main() {
	database.CreateDatabaseConnection()

	app := &cli.App{
		Name: "shortenurl",

		Commands: newCommands(),
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func newCommands() []*cli.Command {
	return []*cli.Command{
		{
			Name:  "encode",
			Usage: "encode url to shorten url",
			Action: func(c *cli.Context) error {
				url := c.Args().Slice()[0]
				shortenUrl := shortenurl.EncodeShortenUrl(url)

				fmt.Println("short url: " + shortenUrl.ShortUrl)

				return nil
			},
		},
		{
			Name:  "decode",
			Usage: "decode shorten url to original url",
			Action: func(c *cli.Context) error {
				url := c.Args().Slice()[0]
				shortenUrl := shortenurl.DecodeShortenUrl(url)

				fmt.Println("original url: " + shortenUrl.Url)

				return nil
			},
		},
	}
}
