package main

import (
	"log"
	"math/rand"
	"time"
)

type order struct {
	Id      int
	Name    string
	Address string
	Amount  float32
	UserId  string
	Time    int64
}

func generateRandString(n int) string {
	var chars = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0987654321")
	str := make([]rune, n)
	for i := range str {
		str[i] = chars[rand.Intn(len(chars))]
	}
	return string(str)
}

func generateRandomOrder(userId string) bool {
	_, err := db.Exec("insert into order_info(name,deliver_address,amount,user_id,time) values(?,?,?,?,?)", generateRandString(20), generateRandString(20), rand.Intn(999), userId, time.Now())
	if err != nil {
		log.Println("failed to generate order:", err)
		return false
	}
	return true
}

func getLastOrder(userId string) (order, error) {
	var o order
	err := db.QueryRow("select name,deliver_address,amount from order_info where user_id = ? order by time desc limit 1", userId).Scan(&o.Name, &o.Address, &o.Amount)
	if err != nil {
		log.Println("get order error:", err)
		return order{}, err
	} else {
		return o, nil
	}
}
