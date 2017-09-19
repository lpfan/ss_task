package main

import (
	"os"

	"client"
	"crawler"

	"github.com/urfave/cli"
)

func main() {
	var pageNum int

	app := cli.NewApp()
	app.Flags = []cli.Flag{
		cli.IntFlag{
			Name:        "pages",
			Value:       5,
			Usage:       "Set number of pages to scrape",
			Destination: &pageNum,
		},
	}
	app.Commands = []cli.Command{
		{
			Name:    "scrape",
			Aliases: []string{"s", "scr"},
			Usage:   "Start scraping process",
			Action: func(c *cli.Context) error {
				crawler.StartScraping(pageNum)
				return nil
			},
		},
		{
			Name:    "run_client",
			Aliases: []string{"rcl", "run_cl"},
			Usage:   "Start client",
			Action: func(c *cli.Context) error {
				client.RunClient()
				return nil
			},
		},
	}

	app.Run(os.Args)
}
