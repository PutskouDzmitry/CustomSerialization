package main

import (
	"encoding/json"
	"fmt"
)

type Message struct {
	From string     `json:"from"`
	To   string     `json:"to"`
	Data string `json:"data"`
}

func main() {
	msg := Message{
		From: "XiaoMing",
		To:   "LiGang",
		Data: `{"title":"test","body":"something"}`,
	}
	jsonData, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(jsonData))
}


