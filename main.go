package main

import (
	// "log"
	// "github.com/gorilla/mux"
	// "net/http"
	"fmt"
	"math/rand"
	"time"
)

func printSomething(id int) {
	waitTime := rand.Intn(1000)
	time.Sleep(time.Duration(waitTime) * time.Millisecond)
	fmt.Println("Something here", id, "-", waitTime)
}

func start() {

	for i := 1; i <= 1000; i++ {
		go printSomething(i)
	}

	time.Sleep(time.Duration(10000) * time.Millisecond)

	// router := mux.NewRouter()
	// log.Fatal(http.ListenAndServe(":8000", router))

}
