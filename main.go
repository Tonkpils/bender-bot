package main

import (
	"log"
	"os"

	"github.com/Tonkpils/grot"
	_ "github.com/Tonkpils/grot/adapters/slack"
)

func init() {
	log.SetOutput(os.Stdout)
	log.SetFlags(0)
}

func main() {
	log.Println("system going online...")
	bot := grot.NewBot()

	if err := bot.Run(); err != nil {
		log.Fatal(err)
	}
}
