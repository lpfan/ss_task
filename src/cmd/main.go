package main

import (
	"crawler"
	"flag"
)

func main() {
	var actionFlag = flag.String("action", "scrape", "Provide action name to execute")
	flag.Parse()

	switch *actionFlag {
	case "scrape":
		crawler.ScrapeNewTopics()
	}
}
