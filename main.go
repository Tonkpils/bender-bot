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

	bot.Hear(`kill all (humans|robots)`, func(res *grot.Response) {
		if res.Matches[1] == "robots" {
			res.Send("This is the worst kind of discrimination. The kind against me!")
		} else {
			res.Send("Now we're talking!")
		}
	})

	if err := bot.Run(); err != nil {
		log.Fatal(err)
	}
}
