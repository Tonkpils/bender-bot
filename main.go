package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/Tonkpils/grot"
	_ "github.com/Tonkpils/grot/adapters/slack"
)

func init() {
	log.SetOutput(os.Stdout)
	log.SetFlags(0)
}

type EchoSkillResponse struct {
	Version           string                 `json:"version"`
	SessionAttributes map[string]interface{} `json:"sessionAttributes,omitempty"`
	Response          EchoResponse           `json:"response"`
}

type EchoResponse struct {
	OutputSpeech OutputSpeech `json:"outputSpeech,omitempty"`
}

type OutputSpeech struct {
	Type string `json:"type"`
	Text string `json:"text,omitempty"`
	SSML string `json:"ssml,omitempty"`
}

func NewEchoResponse() *EchoSkillResponse {
	return &EchoSkillResponse{
		Version: "1.0",
		Response: EchoResponse{
			OutputSpeech: OutputSpeech{
				Type: "PlainText",
				Text: "Hello World!",
			},
		},
	}
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

	bot.Router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		bdy, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.Write([]byte("PONG"))
		}

		fmt.Println(string(bdy))

		resp := NewEchoResponse()
		jsonResp, err := json.Marshal(resp)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
		w.Write(jsonResp)
	})

	if err := bot.Run(); err != nil {
		log.Fatal(err)
	}
}
