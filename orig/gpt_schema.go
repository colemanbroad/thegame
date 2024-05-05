package main

import (
// "encoding/json"
// "fmt"
// "log"
)

type Room struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Team struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Tournament string `json:"tournament"`
}

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Competitor struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Score      string `json:"score"`
	ExtraScore string `json:"extra_score"`
}

type Match struct {
	ID          string       `json:"id"`
	Date        string       `json:"date"`
	Competitors []Competitor `json:"competitors"`
}

type UserScore struct {
	Date  string            `json:"date"`
	Match string            `json:"match"`
	Users []UserScoreDetail `json:"users"`
}

type UserScoreDetail struct {
	ID    string `json:"id"`
	Score int    `json:"score"`
}

type Data struct {
	Room       Room        `json:"room"`
	Teams      []Team      `json:"teams"`
	Users      []User      `json:"users"`
	Matches    []Match     `json:"matches"`
	UserScores []UserScore `json:"user_scores"`
}
