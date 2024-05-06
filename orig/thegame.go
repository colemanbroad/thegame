package main

import (
	// "encoding"
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

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

	var p0, p1 string
	if pens == "" {
		p0 = ""
		p1 = ""
	} else {
		p := strings.Split(pens, ":")
		if len(p) != 2 {
			log.Fatalf("Format Error: The format for penalties must be"+`\d+:\d+`+" but we got %v .\n", pens)
		}
		p0 = p[0]
		p1 = p[1]
	}

	c1 := Competitor{ID: "", Name: home, Score: home_score, ExtraScore: p0}
	c2 := Competitor{ID: "", Name: away, Score: away_score, ExtraScore: p1}
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

type AllData struct {
	Data     Data
	team_map map[string][]string
	user_map map[string][]string
}

func load_dan_table(csv_teams_prices_owners, csv_match_results string) AllData {

	csv_countries_prices_owners := readCsvFile(csv_teams_prices_owners)
	countries := csv_countries_prices_owners[0][1:]
	prices := csv_countries_prices_owners[1][1:]

	for i, p := range prices {
		c := countries[i]
		// countries[i] = strings.Title(c)
		countries[i] = cases.Title(language.AmericanEnglish).String(c)
		fmt.Println(c, p)
	}

	fmt.Println("\n\nNow we're trying to print out the names of the team's players.")
	fmt.Println()
	fmt.Println()

	team_map := make(map[string][]string)
	user_map := make(map[string][]string)

	// can each team (one per column). User names start at row 15 and stop with a blank cell.
	for c := 1; c < len(csv_countries_prices_owners[0]); c++ {
		land := countries[c-1]
		for r := 15; r < len(csv_countries_prices_owners); r++ {
			x := csv_countries_prices_owners[r][c]
			if x == "" {
				break
			}
			team_map[land] = append(team_map[land], x)
		}
	}

	// Invert the team->players map to get player->teams
	for team, players := range team_map {
		fmt.Println(team, players)
		for _, p := range players {
			user_map[p] = append(user_map[p], team)
		}
	}

	fmt.Println("\n\nThe following players own the following teams.")
	fmt.Println()
	fmt.Println()

	for player, teams := range user_map {
		fmt.Println(player, teams)
	}

	// now let's build the basic tables

	// teams are all the countries listed at the top of Dan's spreadsheet
	// TODO: confirm that they match all the teams from the games spreadsheet
	room := Room{ID: "0", Name: "Dan's Game WWC 2023"}
	team2ID := make(map[string]string)
	teams := make([]Team, len(countries))
	for i, team := range countries {
		id := fmt.Sprint(i)
		teams[i] = Team{ID: id, Name: team, Tournament: room.Name}
		team2ID[team] = id
	}

	fmt.Println(team2ID)

	// Users are all the people who bought teams
	// We can get all the users by looking at the keys of the user_map
	users := make([]User, 0, len(user_map))
	i := 0
	for u := range user_map {
		id := fmt.Sprint(i)
		users = append(users, User{ID: id, Name: u})
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
		match.ID = fmt.Sprint(i)
		var t string

		t = match.Competitors[0].Name
		fmt.Println("T = ", t, team2ID[t])
		match.Competitors[0].ID = team2ID[t]

		t = match.Competitors[1].Name
		fmt.Println("T = ", t, team2ID[t])
		match.Competitors[1].ID = team2ID[t]

	}

	// build indexes

	data := Data{Room: room, Teams: teams, Users: users, Matches: matches, Results: nil}

	alldata := AllData{
		Data:     data,
		team_map: team_map,
		user_map: user_map,
	}

	return alldata
}

func fill_in_user_scores(alldata *AllData) {
	// fill in user scores
	// this means we generate scores for users (after each round).
	// we can build a map from users to scores given a list of Matches
	// and we can filter matches by date / round.
	// as long as we don't need to do any prediction then we don't really need to know the tournament structure!
	// just the dates will be fine...
	// this means we generate scores for users given

	data := &alldata.Data

	results := make([]MatchResults, 0, len(data.Matches))

	teams_points := make(map[string]int, len(data.Matches))
	for _, team := range data.Teams {
		teams_points[team.Name] = 0
	}

	point_totals := make(map[string]int)

	// first calculate team points
	for idx, match := range data.Matches {
		if idx == 62 {
			// TODO: FIXME: We don't count the 3rd place game!!!
			continue
		}
		s0 := match.Competitors[0].Score
		s1 := match.Competitors[1].Score
		pen0 := match.Competitors[0].ExtraScore
		pen1 := match.Competitors[1].ExtraScore
		n0 := match.Competitors[0].Name
		n1 := match.Competitors[1].Name

		type PtsShared struct{ p0, p1 int }
		pts := PtsShared{p0: 0, p1: 0}

		// update teams_points the total number
		if pen0 == "" && pen1 == "" {
			if s0 > s1 {
				pts.p0 = 3
			} else if s0 < s1 {
				pts.p1 = 3
			} else {
				pts.p0 = 1
				pts.p1 = 1
			}
		} else if pen0 > pen1 {
			pts.p0 = 3
		} else if pen0 < pen1 {
			pts.p1 = 3
		} else {
			log.Fatalf("Invalid format: Cannot have equal, nonempty penalty scores. %v vs %v score: %v == %v .\n", n0, n1, pen0, pen1)
		}

		teams_points[n0] += pts.p0
		teams_points[n1] += pts.p1

		// loop over users that own home team and

		f_user := func(user_str string) (*User, error) {
			for _, u := range data.Users {
				if u.Name == user_str {
					return &u, nil
				}
			}
			return nil, errors.New("User not found")
		}

		pointslist := make([]UserPoints, 0, len(data.Users))
		for _, user_str := range alldata.team_map[n0] {
			user, err := f_user(user_str)
			if err != nil {
				log.Panic(err)
			}
			pointslist = append(pointslist, UserPoints{UserID: user.ID, Points: pts.p0})
			point_totals[user.Name] += pts.p0
			if pts.p0 > 0 && user.Name == "SHAWN" {
				fmt.Printf("Shawn got %v points for match %v vs %v ", pts.p0, match.Competitors[0].Name, match.Competitors[1].Name)
			}
		}
		for _, user_str := range alldata.team_map[n1] {
			user, err := f_user(user_str)
			if err != nil {
				log.Panic(err)
			}
			pointslist = append(pointslist, UserPoints{UserID: user.ID, Points: pts.p1})
			point_totals[user.Name] += pts.p1
			if pts.p1 > 0 && user.Name == "SHAWN" {
				fmt.Printf("Shawn got %v points for match %v vs %v ", pts.p1, match.Competitors[0].Name, match.Competitors[1].Name)
			}
		}

		results = append(results, MatchResults{Match: match, PointsList: pointslist})

	}

	for k, v := range teams_points {
		fmt.Println(k, v)
	}

	fmt.Println("The point_totals in the end are .......")
	for k, v := range point_totals {
		fmt.Println(k, v)
	}

	data.Results = results
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

	alldata := load_dan_table("../data/wc2023.csv", "../data/wc2023_games.csv")

	fill_in_user_scores(&alldata)

	prettyJSON, err := json.MarshalIndent(alldata.Data, "", "    ") // Using four spaces for indentation
	if err != nil {
		log.Fatalf("Error generating pretty JSON: %v", err)
	}
	// _ = prettyJSON
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
