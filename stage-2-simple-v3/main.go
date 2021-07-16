package main

import (
	"encoding/json"
	"os"
	"time"
)

func (u *MyUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		LastSeen int64 `json:"lastSeen"`
		*MyUser
	}{
		LastSeen: u.LastSeen.Unix(),
		MyUser:   u,
	})
}

type MyUser struct {
	ID       int64     `json:"id"`
	Name     string    `json:"name"`
	LastSeen time.Time `json:"lastSeen"`
}

func main() {
	_ = json.NewEncoder(os.Stdout).Encode(
		&MyUser{1, "Ken", time.Now()},
	)
}
