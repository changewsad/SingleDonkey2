package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"math/rand"
	
	"github.com/line/line-bot-sdk-go/linebot"
)

var bot *linebot.Client

func main() {
	var err error
	bot, err = linebot.New(os.Getenv("ChannelSecret"), os.Getenv("ChannelAccessToken"))
	log.Println("Bot:", bot, " err:", err)
	
	http.HandleFunc("/callback", callbackHandler)
	port := os.Getenv("PORT")
	addr := fmt.Sprintf(":%s", port)
	http.ListenAndServe(addr, nil)
}

func callbackHandler(w http.ResponseWriter, r *http.Request) {
	events, err := bot.ParseRequest(r)

	if err != nil {
		if err == linebot.ErrInvalidSignature {
			w.WriteHeader(400)
		} else {
			w.WriteHeader(500)
		}
		return
	}
	//if _, err := bot.PushMessage("U2c68fd429a99dceccc8956571baa7d00", linebot.NewTextMessage("hello")).Do(); err != nil {
	//	txt= txt + err.Error()
	//}
	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				//var txt = Send(message.Text);
				//rand.Seed(99)
				answers :=[]string{"咿喔","咿~喔~","咿喔咿喔喔","咿喔~咿喔","咿咿喔喔~","咿~喔咿~喔","咿喔喔~~","喔~咿喔~~","咿喔咿喔咿喔"}
				
				var txt = message.Text+","+answers[rand.Intn(len(answers))]
				if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(txt)).Do(); err != nil {
					log.Print(err)
				}
				
			}
		}
		
	}
}
