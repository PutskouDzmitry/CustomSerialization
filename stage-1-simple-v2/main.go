package main

import (
	"encoding/json"
	"fmt"
)

type JsonString string

func (j JsonString) MarshalJSON() ([]byte, error) {
	return []byte(j), nil
}

type Message struct {
	From string     `json:"from"`
	To   string     `json:"to"`
	Data JsonString `json:"data"`
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

