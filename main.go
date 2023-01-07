package main

import (
	"log"

	"github.com/kmrtftech/tg-framework/internal/app"
)

func main() {
	if err := app.Main(); err != nil {
		log.Fatal(err.Error())
	}
}
