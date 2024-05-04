package main

import (
	"encoding/json"
	"fmt"
	// "io/ioutil"
	// "log"
	// "os"
	"testing"
	// "go-hep.org/x/hep/xrootd/xrdproto/open"
)

func TestJson(t *testing.T) {
	var err error

	// f, err := os.Open("../data/rest_schema.json")
	// filebytes, err := os.ReadFile("../data/rest_example.json")
	// if err != nil {
	// 	t.Errorf("Couldn't read file.")
	// }
	// jsonStr := string(filebytes)
	// jsonStr := f.Read()
	// jsonStr := readCsvFile()

	jsonStr := `{
	       "room" : {
	         "id" : "1",
	         "name" : "Vlad's Pool Party"
	       },
	       "teams": [
	           {
	               "id": "1",
	               "name": "Arsenal"
	           }
	       ],
	       "users": [
	           {
	               "id": "1",
	               "name": "Vlad"
	           }
	       ],
	       "matches": [
	           {
	               "id": "1",
	               "date":"2024-12-15T000",
	               "competitors": [
	                   {
	                       "id": "1",
	                       "score": "3",
	                       "extra_score": "5"
	                   },
	                   {
	                       "id": "1",
	                       "score": "4",
	                       "extra_score": "0"
	                   }
	               ]
	           }
	       ],
	       "user_scores" : [
	           {
	               "date":"2024-12-15T000",
	               "match":"1",
	               "users":[
	                   {
	                       "id":"1",
	                       "score":23424
	                   }
	               ]
	           }
	       ]
	   }`

	var data Data
	err = json.Unmarshal([]byte(jsonStr), &data)
	if err != nil {
		t.Errorf("Error parsing JSON: %v", err)
	}

	prettyJSON, err := json.MarshalIndent(data, "", "    ") // Using four spaces for indentation
	if err != nil {
		t.Errorf("Error generating pretty JSON: %v", err)
	}

	fmt.Printf("%s\n", prettyJSON)

	// fmt.Printf("%+v\n", data)
}
