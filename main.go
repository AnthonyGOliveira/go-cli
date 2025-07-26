package main

import (
	"log"
	"os"

	"github.com/AnthonyGOliveira/go-cli/app"
)

func main() {
	application := app.GenerateApp()
	if err := application.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
