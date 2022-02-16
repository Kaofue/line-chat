package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/line/line-bot-sdk-go/linebot"
)

var (
	client *linebot.Client
	err    error
)

func main() {
	fmt.Println("run!!")
	// 建立客戶端
	client, err = linebot.New("949137fe35fc862b5ef1dfe5b0fb2de9", "7yuKcXUPEWS94OSGeORHagrXAWVtQmwjMpMANJOlOqTD5PEaaq00YuNiATkJXVvU3pgcrYOcqGIkF3oBpcH9dP7gz29fG5SS+ZD4ITwBQAZw+10WG4VDp3qM8ygtJhoLNRTbR9U87Q/nklUnvYfV1wdB04t89/1O/w1cDnyilFU=")

	if err != nil {
		log.Println(err.Error())
	}

	http.HandleFunc("/callback", callbackHandler)

	log.Fatal(http.ListenAndServe(":84", nil))
}

func callbackHandler(w http.ResponseWriter, r *http.Request) {
	// 接收請求
	events, err := client.ParseRequest(r)

	if err != nil {
		if err == linebot.ErrInvalidSignature {
			w.WriteHeader(400)
		} else {
			w.WriteHeader(500)
		}

		return
	}

	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				// 回覆訊息
				if _, err = client.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.Text)).Do(); err != nil {
					log.Println(err.Error())
				}
			}
		}
	}
}