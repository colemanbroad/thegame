package main

import (
	// "encoding"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
	// "strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	// "gonum.org/v1/plot/text"
	// "strconv"
)

type Tournament struct {
	id string
	// competitors[2]
	// the matchID is a short string that uniquely identifies each game
	// somehow we need to encode the structure of the Tournament.
	matchID2idx map[string]uint32
}

// type Match struct {
// 	home       string
// 	away       string
// 	home_score int8
// 	away_score int8
// 	pens       bool
// 	home_pens  int8
// 	away_pens  int8
// 	extra_time uint16
// 	group      string
// 	round      string
// }

// type Room struct {
// }

// add team to db
// add player to db
// mapping player >< teams
// Game(id, home, away, homescore, awayscore)

// User api

// RoomHost api
func addPlayer(player string)       {}
func addTeam(team string, cost int) {}
func createRoom(room string)        {}

func newMatch(home, away, home_score, away_score, date, pens string) Match {

	c1 := Competitor{ID: "", Name: home, Score: home_score, ExtraScore: pens}
	c2 := Competitor{ID: "", Name: away, Score: away_score, ExtraScore: pens}
	return Match{ID: "", Date: date, Competitors: []Competitor{c1, c2}}

}

// type Team string
// type Player string

func loadGames(csvname string) []Match {
	records := readCsvFile(csvname)
	games := make([]Match, 0, 64)

	titler := cases.Title(language.AmericanEnglish)

	for id, row := range records {
		if row[0] == "" {
			continue
		}
		homescore := row[1][0]
		awayscore := row[1][2]
		date := row[4]
		pens := row[3]
		home := titler.String(row[0])
		away := titler.String(row[2])
		g := newMatch(home, away, string(homescore), string(awayscore), date, pens)
		g.ID = fmt.Sprint(id)
		games = append(games, g)
	}

	return games
}

func load_dan_table(csv_teams_prices_owners, csv_match_results string) Data {

	records := readCsvFile(csv_teams_prices_owners)
	countries := records[0][1:]
	prices := records[1][1:]

	for i, p := range prices {
		c := countries[i]
		// countries[i] = strings.Title(c)
		countries[i] = cases.Title(language.AmericanEnglish).String(c)
		fmt.Println(c, p)
	}

	fmt.Println("\n\nNow we're trying to print out the names of the team's players.")
	fmt.Println()
	fmt.Println()

	teammap := make(map[string][]string)
	playermap := make(map[string][]string)

	// can each team (one per column). User names start at row 15 and stop with a blank cell.
	for c := 1; c < len(records[0]); c++ {
		land := countries[c-1]
		for r := 15; r < len(records); r++ {
			x := records[r][c]
			if x == "" {
				break
			}
			teammap[land] = append(teammap[land], x)
		}
	}

	// Invert the team->players map to get player->teams
	for team, players := range teammap {
		fmt.Println(team, players)
		for _, p := range players {
			playermap[p] = append(playermap[p], team)
		}
	}

	fmt.Println("\n\nThe following players own the following teams.")
	fmt.Println()
	fmt.Println()

	for player, teams := range playermap {
		fmt.Println(player, teams)
	}

	room := Room{ID: "0", Name: "Dan's Game WWC 2023"}
	team2ID := make(map[string]string)
	teams := make([]Team, 0, len(teammap))
	i := 0
	for team, _ := range teammap {
		id := fmt.Sprint(i)
		teams = append(teams, Team{ID: id, Name: team, Tournament: room.Name})
		team2ID[team] = id
		i++
	}

	fmt.Println(team2ID)

	users := make([]User, 0, 50)
	i = 0
	for player, _ := range playermap {
		users = append(users, User{ID: fmt.Sprint(i), Name: player})
		i++
	}

	// games := loadGames(csv_match_results)
	// for _, g := range games {
	// 	fmt.Println(g)
	// }
	matches := loadGames(csv_match_results)

	// add IDs to matches and competitors
	for i := range matches {

		match := &matches[i]
		match.ID = fmt.Sprint(i + 99)
		var t string

		t = match.Competitors[0].Name
		fmt.Println("T = ", t, team2ID[t])
		match.Competitors[0].ID = team2ID[t]

		t = match.Competitors[1].Name
		fmt.Println("T = ", t, team2ID[t])
		match.Competitors[1].ID = team2ID[t]

	}

	user_scores := make([]UserScore, 1, 1400)

	// for i, user := range users {
	// 	user_scores[0]
	// }

	// type Data struct {
	// 	Room       Room        `json:"room"`
	// 	Teams      []Team      `json:"teams"`
	// 	Users      []User      `json:"users"`
	// 	Matches    []Match     `json:"matches"`
	// 	UserScores []UserScore `json:"user_scores"`
	// }
	// Teams := Team{}
	// Room := Room{ID: "0", Name: "Dan's Game WWC 2023"}

	data := Data{Room: room, Teams: teams, Users: users, Matches: matches, UserScores: user_scores}

	return data

	// return Data{Room, Teams, Users, Matches, UserScores}
}

func fill_in_user_scores(data Data) {
}

// We need to convert from our internal representation using maps to the REST API json representation
// This will not be trivial.
// Foundational tables are Tournament, Team, Match, Room, User,
// All need unique id.
// User(id string, admin bool, name string, email string)
// Tournament(id string, )
// Then we need a relation between Users >< Teams (who owns which teams?)
// And a relation between Match < Team (potentially built into the match table?)
// And a relation between Tournament < Match
// And a relation between Room - Tournament
// Room < User

func main() {

	data := load_dan_table("../data/wc2023.csv", "../data/wc2023_games.csv")

	prettyJSON, err := json.MarshalIndent(data, "", "    ") // Using four spaces for indentation
	if err != nil {
		log.Fatalf("Error generating pretty JSON: %v", err)
	}
	fmt.Printf("%s\n", prettyJSON)

}

func readCsvFile(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	return records
}
