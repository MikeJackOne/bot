package Messenger

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"os"
)

func convertGenericMessage(message FbMessage) (Message, error) {
	var tmp Message
	if len(message.Entry) > 0 && len(message.Entry[0].Messaging) > 0 {
		entry := message.Entry[0]
		tmessage := message.Entry[0].Messaging[0]
		tmp.ID = entry.ID
		tmp.Time = entry.Time
		tmp.SenderId = tmessage.Sender.ID
		tmp.RecipientId = tmessage.Recipient.ID
		tmp.Text = tmessage.Message.Text
		return tmp, nil
	}
	return tmp, nil
}

func (bot *Bot) processMessage(c *gin.Context) {

	//jsonData, err := io.ReadAll(c.Request.Body)

	var message FbMessage

	err := c.ShouldBindJSON(&message)

	if err != nil {
		log.Println("error whiling getting json body")
		c.String(http.StatusBadRequest, "wrong message")
	}

	quickMessage, err := convertGenericMessage(message)

	if err != nil {
		log.Println("error while receving message:", message)
		log.Println("returning")
		c.String(http.StatusBadRequest, "wrong message")
	}

	log.Println(quickMessage)

	for _, f := range bot.core {
		f(quickMessage)
	}

	c.String(http.StatusOK, "ok")
}

func newMessage(receipt string, m string) Reply {
	return Reply{
		Message:   SmallMessage{Text: m},
		Recipient: Recipient{ID: receipt},
	}
}

func SendSimpleMessage(receipt string, m string) {

	var replyMessage Reply

	replyMessage = newMessage(receipt, m)

	message, err := json.Marshal(replyMessage)

	if err != nil {
		log.Println("error while converting message into JSON ", message)
		return
	}

	resp, err := http.Post(apiUrl, "application/json", bytes.NewBuffer(message))

	if err != nil {
		//Try again
		resp, err = http.Post(apiUrl, "application/json", bytes.NewBuffer(message))

		if err != nil {
			log.Println("error sending1 the message:", err)
			return
		}
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		data, _ := io.ReadAll(resp.Body)
		log.Println("error sending the message:", string(data), " ", string(message), " ", apiUrl)
	}
}

func (bot *Bot) simpleCheck(c *gin.Context) {
	var challenge string
	challenge = c.Query("hub.challenge")

	token := c.Query("hub.verify_token")

	if token != bot.Token {
		c.String(http.StatusBadRequest, "authentication failed")
		return
	}

	//jsonData, err := io.ReadAll(c.Request.Body)

	c.String(http.StatusOK, challenge)
}

func (bot *Bot) AppendJob(function func(Message) bool) {
	bot.core = append(bot.core, function)
}

func (bot *Bot) Start(f *os.File) {
	//gin.DisableConsoleColor()

	//gin.DefaultWriter = io.MultiWriter(f)

	apiUrl = bot.ApiUrl + bot.AccessToken

	r := gin.Default()

	r.POST("/webhook", bot.processMessage)
	r.GET("/webhook", bot.simpleCheck)
	r.Run()
}
