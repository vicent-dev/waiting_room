package main

import (
	"log"
	"waiting_room/app"
)

const ()

func main() {
	if err := app.Run(); err != nil {
		log.Fatalf("Output error: %s \n", err.Error())
	}
}
