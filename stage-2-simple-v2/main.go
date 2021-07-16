package main

import (
	"encoding/json"
	"os"
	"time"
)

func (u *MyUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		ID       int64  `json:"id"`
		Name     string `json:"name"`
		LastSeen int64  `json:"lastSeen"`
	}{
		ID:       u.ID,
		Name:     u.Name,
		LastSeen: u.LastSeen.Unix(),
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
