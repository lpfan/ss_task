package client

import (
	"fmt"
	"log"
	"os"
)

const (
	frontendDirName        = "frontend"
	serverPort             = "8080"
	frontendApplicationDir = "dist"
)

type configuration struct {
	ServerPort                    string
	FrontendPath                  string
	FrontendApplicationStaticPath string
}

// create configuration object
func GetConfiguration() *configuration {
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatal("Can't find configuration directory for clien application")
	}

	config := configuration{}
	config.ServerPort = serverPort
	config.FrontendPath = fmt.Sprintf("%s/%s/%s", currentDir, "src/client", frontendDirName)
	config.FrontendApplicationStaticPath = fmt.Sprintf("%s/%s", config.FrontendPath, frontendApplicationDir)
	return &config
}
