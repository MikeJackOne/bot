package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	Messenger "github.com/loser02/bot/messenger"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
)

var db *sql.DB

func ReadConfig() Messenger.Bot {
	var conf Messenger.Bot
	configFile, err := os.Open("config.json")
	if err != nil {
		log.Fatal("opening config file", err.Error())
	}

	jsonParser := json.NewDecoder(configFile)
	if err = jsonParser.Decode(&conf); err != nil {
		log.Fatal("parsing config file", err.Error())
	}
	return conf
}

func initSqlite() {
	dbTmp, err := sql.Open("sqlite3", "order.db")
	if err != nil {
		log.Fatal("failed to connect sqlite! ", err)
	}
	db = dbTmp
}

func main() {

	messenger := ReadConfig()
	initSqlite()
	f, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()
	log.SetOutput(os.Stdout)

	messenger.AppendJob(HelloBack)
	messenger.AppendJob(GenerateOrder)
	messenger.AppendJob(GetOrder)

	messenger.Start(f)
}

func HelloBack(m Messenger.Message) bool {
	log.Println(m.Text)
	if m.Text == "thank you" {
		Messenger.SendSimpleMessage(m.SenderId, "You are welcome!")
	}
	return true
}

func GenerateOrder(m Messenger.Message) bool {
	if m.Text == "generate order" {
		result := generateRandomOrder(m.SenderId)
		if result {
			Messenger.SendSimpleMessage(m.SenderId, "You have successfully generate an order!")
		} else {
			Messenger.SendSimpleMessage(m.SenderId, "order generate failed!")
		}
	}
	return true
}

func GetOrder(m Messenger.Message) bool {
	if m.Text == "get last order" {
		user_order, err := getLastOrder(m.SenderId)
		if err != nil {
			Messenger.SendSimpleMessage(m.SenderId, "no order")
		} else {
			order_message := "your last order information is following:\n"
			order_message += "order name:" + user_order.Name + "\n"
			order_message += "deliver address:" + user_order.Address + "\n"
			order_message += fmt.Sprintf("amount:%.2f\n", user_order.Amount)
			Messenger.SendSimpleMessage(m.SenderId, order_message)
		}
	}
	return true
}
