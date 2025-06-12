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

func increment(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	counter++
	fmt.Fprint(w, strconv.Itoa(counter))
	mutex.Unlock()
}

//	func homefunc(w http.ResponseWriter, r *http.Request) {
//		http.ServeFile(w,r,r.URL.Path[1:])
//	}
func main() {
	fmt.Print("Please enter the API_Key for Gemini model or 0 to continue without: ")
	var api_key string
	fmt.Scan(&api_key)
	if len(api_key) == 2 {
		fmt.Println(handle_response("Give a simple poem about fishes", api_key))
	}
	//TODO: Add a function to handle get requests for the query
	//TODO: Add the nececessary modals in HTML file to get the data and make req
	//TODO: Add function call for the get request
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/increment", increment)
	http.HandleFunc("/decrement", func(w http.ResponseWriter, r *http.Request) {

		mutex.Lock()
		counter--
		fmt.Fprint(w, strconv.Itoa(counter))
		mutex.Unlock()
	})
	port := ":8081"
	fmt.Printf("Now serving port: %s", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
