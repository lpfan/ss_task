package main

import (
	"crawler"
	"flag"
)

func main() {
	var actionFlag = flag.String("action", "scrape", "Provide action name to execute")
	var pageNumFlag = flag.Int("pages", 5, "Set number of pages to scrape")
	flag.Parse()

	switch *actionFlag {
	case "scrape":
		crawler.StartScraping(*pageNumFlag)
	}
}
