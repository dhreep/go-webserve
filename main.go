package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

var counter int
var mutex = &sync.Mutex{}
var api_key string

func increment(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	counter++
	fmt.Fprint(w, strconv.Itoa(counter))
	mutex.Unlock()
}

func handle_query(w http.ResponseWriter, r *http.Request) {
	//TODO: write the code to pull data from the query
	query := r.URL.Query().Get("query")
	if len(query) == 0 {
		fmt.Println("The request lacks a query")
		return
	}
	result := handle_response(query, api_key)
	//fmt.Fprint(w,result)
	fmt.Fprint(w, "Api Call Works "+result)
}

//	func homefunc(w http.ResponseWriter, r *http.Request) {
//		http.ServeFile(w,r,r.URL.Path[1:])
//	}
func main() {
	fmt.Print("Please enter the API_Key for Gemini model or 0 to continue without: ")
	fmt.Scan(&api_key)
	if len(api_key) == 2 {
		fmt.Println(handle_response("Give a simple poem about fishes", api_key))
	}
	//TODO: Add the nececessary modals in HTML file to get the data and make req
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/increment", increment)
	http.HandleFunc("/query", handle_query) //url splitting
	http.HandleFunc("/decrement", func(w http.ResponseWriter, r *http.Request) {

		mutex.Lock()
		counter--
		fmt.Fprint(w, strconv.Itoa(counter))
		mutex.Unlock()
	})
	port := ":8081"
	fmt.Printf("Now serving port: %s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
