package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	// "strconv"
)

type Game struct {
	home       string
	away       string
	home_score int8
	away_score int8
	pens       bool
	home_pens  int8
	away_pens  int8
	extra_time uint16
	group      string
	round      string
}

func newGame(home, away string, home_score, away_score int8) Game {
	return Game{
		home:       home,
		away:       away,
		home_score: home_score,
		away_score: away_score,
	}
}

// type Team string
// type Player string

func loadGames() []Game {
	records := readCsvFile("wc2023_games.csv")
	games := make([]Game, 0, 64)

	for _, row := range records {
		if row[0] == "" {
			continue
		}
		homescore := uint8(row[1][0])
		awayscore := uint8(row[1][2])
		// homescore, err := strconv.Atoi(row[1][0])
		// if err != nil {
		// 	panic(err)
		// }
		// awayscore, err := strconv.Atoi(row[1][2])
		// if err != nil {
		// 	panic(err)
		// }
		g := newGame(row[0], row[2], int8(homescore), int8(awayscore))
		// fmt.Println(g)
		games = append(games, g)
	}

	return games

}

func runTheGame() map[string][]string {

	games := loadGames()
	for _, g := range games {
		fmt.Println(g)
	}
	// fmt.Println(games)

	panic("i've done a panic!")

	records := readCsvFile("wc2023.csv")
	countries := records[0][1:]

	prices := records[1][1:]

	for i, p := range prices {
		c := countries[i]
		fmt.Println(c, p)
	}

	fmt.Println("\n\nNow we're trying to print out the names of the team's players.")
	fmt.Println()
	fmt.Println()

	teammap := make(map[string][]string)
	playermap := make(map[string][]string)

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

	return playermap
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