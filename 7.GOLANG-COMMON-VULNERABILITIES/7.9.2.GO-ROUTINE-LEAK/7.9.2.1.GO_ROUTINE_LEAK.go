package main

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
	"strconv"
)

// get the count of number of go routines in the system.
func countGoRoutines() int {
	return runtime.NumGoroutine()
}

func getGoroutinesCountHandler(w http.ResponseWriter, r *http.Request) {
	// Get the count of number of go routines running.
	count := countGoRoutines()
	w.Write([]byte(strconv.Itoa(count)))
}

// function to add an array of numbers.
func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	// writes the sum to the go routines.
	c <- sum // send sum to c
}

// HTTP handler for /sum
func sumConcurrent(w http.ResponseWriter, r *http.Request) {
	s := []int{7, 2, 8, -9, 4, 0}

	c1 := make(chan int)
	c2 := make(chan int)
	// spin up a goroutine.
	go sum(s[:len(s)/2], c1)
	// spin up a goroutine.
	go sum(s[len(s)/2:], c2)
	// not reading from c2.
	// go routine writing to c2 will be blocked.
  // Since we are not reading from c2, 
  // the goroutine attempting to write to c2 
  // will be blocked forever resulting in leak.
	x := <-c1
	// write the response.
	fmt.Fprintf(w, strconv.Itoa(x))
}

func main() {
	// get the sum of numbers.
	http.HandleFunc("/sum", sumConcurrent)
	// get the count of number of go routines in the system.
	http.HandleFunc("/_count", getGoroutinesCountHandler)
	err := http.ListenAndServe(":8001", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}