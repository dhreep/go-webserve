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
// func homefunc(w http.ResponseWriter, r *http.Request) {
// 	http.ServeFile(w,r,r.URL.Path[1:])
// }
func main() {
	fmt.Print("Hello World\n")
	http.Handle("/",http.FileServer(http.Dir("./static")))
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
