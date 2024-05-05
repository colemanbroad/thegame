package main

import (
	"fmt"
	"log"
	"net/http"
)

type GlobalState struct {
	Count int
}

var global GlobalState
var user GlobalState
var playermap map[string][]string

func getHandler(w http.ResponseWriter, r *http.Request) {
	// component := page(global.Count, user.Count-global.Count)
	component := renderplayertable(playermap)

	component.Render(r.Context(), w)
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	// Update state.
	r.ParseForm()

	// Check to see if the global button was pressed.

	for k, v := range r.Form {
		fmt.Println(k)
		fmt.Println(v)
	}
	if r.Form.Has("global") {
		global.Count++
	}
	if r.Form.Has("user") {
		user.Count++
	}
	//TODO: Update session.

	// Display the form.
	getHandler(w, r)
}

func startServer() {

	// Handle POST and GET requests.
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			postHandler(w, r)
			return
		}
		getHandler(w, r)
	})

	// Start the server.
	fmt.Println("listening on http://localhost:8000")
	if err := http.ListenAndServe("localhost:8000", nil); err != nil {
		log.Printf("error listening: %v", err)
	}
}
