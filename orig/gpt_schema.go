package main

import (
// "encoding/json"
// "fmt"
// "log"
)

// ----- Back -> Front schema
// These types have already been organized by the backend with valid IDs
//

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

type MatchResults struct {
	// Date   string            `json:"date"`
	// MatchID    string       `json:"match_id"`
	Match      Match        `json:"match"`
	PointsList []UserPoints `json:"points_list"`
}

type UserPoints struct {
	UserID string `json:"user_id"`
	Points int    `json:"points"`
}

type Data struct {
	Room    Room           `json:"room"`
	Teams   []Team         `json:"teams"`
	Users   []User         `json:"users"`
	Matches []Match        `json:"matches"`
	Results []MatchResults `json:"match_results"`
}

// --------

// ---- Front -> Back Schema

type InitialPrices struct {
	Teams  []Team
	Prices []float64
}
