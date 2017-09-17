package main

import (
	"fmt"
	"log"
	"os"
)

const (
	frontendDirName = "frontend"
	serverPort      = 8080
)

type configuration struct {
	ServerPort   int
	FrontendPath string
}

// create configuration object
func GetConfiguration() *configuration {
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatal("Can't find configuration directory for clien application")
	}

	config := configuration{}
	config.ServerPort = serverPort
	config.FrontendPath = fmt.Sprintf("%s/%s", currentDir, frontendDirName)
	return &config
}
