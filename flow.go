package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func j(x int) {
	for i := 0; i < 10; i++ {
		x = +i
	}
	fmt.Println(x)
}

func timeHandler(format string) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		tm := time.Now().Format(format)
		w.Write([]byte("The time is: " + tm))
	}
	return http.HandlerFunc(fn)
}

func main() {
	go j(1)
	go j(101)

	fmt.Println("This program for learning flow")

	mux := http.NewServeMux()

	th := timeHandler(time.RFC1123)
	mux.Handle("/time", th)

	log.Println("Listening...")
	http.ListenAndServe(":3000", mux)
}
