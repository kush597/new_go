package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"sync"
)

const (
	leakyServer = "http://localhost:8001"
)

// get the count of the number of go routines in the server.
func getRoutineCount() (int, error) {
	body, err := getReq("/_count")

	if err != nil {
		return -1, err
	}
	count, err := strconv.Atoi(string(body))
	if err != nil {
		return -1, err
	}
	return count, nil
}

// Send get request and return the repsonse body.
func getReq(endPoint string) ([]byte, error) {
	response, err := http.Get(leakyServer + endPoint)
	if err != nil {
		return []byte{}, err
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return []byte{}, err
	}
	return body, nil
}


func main() {
	// get the number of go routines in the leaky server.
	count, err := getRoutineCount()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("\n %d Go routines before the load test in the system.", count)

	var wg sync.WaitGroup
	// send 50 concurrent request to the leaky endpoint.
	for i := 0; i < 50; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			_, err = getReq("/sum")
			if err != nil {
				log.Fatal(err)
			}

		}()
	}
	wg.Wait()
	// get the cout of number of goroutines in the system after the load test.
	count, err = getRoutineCount()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("\n %d Go routines after the load test in the system.", count)
}