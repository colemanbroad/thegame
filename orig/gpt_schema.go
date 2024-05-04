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
	ID   string `json:"id"`
	Name string `json:"name"`
}

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Competitor struct {
	ID         string `json:"id"`
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

// func main() {
// 	jsonStr := `{
//         "room" : {
//           "id" : "1",
//           "name" : "Vlad's Pool Party"
//         },
//         "teams": [
//             {
//                 "id": "1",
//                 "name": "Arsenal"
//             }
//         ],
//         "users": [
//             {
//                 "id": "1",
//                 "name": "Vlad"
//             }
//         ],
//         "matches": [
//             {
//                 "id": "1",
//                 "date":"2024-12-15T000",
//                 "competitors": [
//                     {
//                         "id": "1",
//                         "score": "3",
//                         "extra_score": "5"
//                     },
//                     {
//                         "id": "1",
//                         "score": "4",
//                         "extra_score": "0"
//                     }
//                 ]
//             }
//         ],
//         "user_scores" : [
//             {
//                 "date":"2024-12-15T000",
//                 "match":"1",
//                 "users":[
//                     {
//                         "id":"1",
//                         "score":23424
//                     }
//                 ]
//             }
//         ]
//     }`

// 	var data Data
// 	err := json.Unmarshal([]byte(jsonStr), &data)
// 	if err != nil {
// 		fmt.Println("Error parsing JSON: ", err)
// 	}

// 	prettyJSON, err := json.MarshalIndent(data, "", "    ") // Using four spaces for indentation
// 	if err != nil {
// 		log.Fatalf("Error generating pretty JSON: %v", err)
// 	}

// 	fmt.Printf("%s\n", prettyJSON)

// 	// fmt.Printf("%+v\n", data)
// }
