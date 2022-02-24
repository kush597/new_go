							UNIT 7 :



Data races and race condition in golang:

A data race occurs when: two or more threads in a single process access the same memory location concurrently, and. at least one of the accesses is for writing, and. the threads are not using any exclusive locks to control their accesses to that memory.


A race condition occurs when the timing or order of events affects the correctness of a piece of code whereas data race occurs when one thread accesses a mutable object while another thread is writing to it.

use the following command to detect data race : $ go run -race filename.go (filename is the name of the file in which go program is stored)



Dead lock in golang:
A deadlock happens when a group of goroutines are waiting for each other and none of them is able to proceed.

Example for deadlock detector: TOY DEADLOCK DETECTOR
PACKAGES TO BE INSTALLED:

Use -d depending on system . 


go get -d "github.com/y-taka-23/ddsv-go/deadlock"
go get -d "github.com/y-taka-23/ddsv-go/deadlock/rule"
go get -d "github.com/y-taka-23/ddsv-go/deadlock/rule/do"
go get -d "github.com/y-taka-23/ddsv-go/deadlock/rule/vars"
go get -d "github.com/y-taka-23/ddsv-go/deadlock/rule/when"


Description:

The dining philosophers problem is one of the best-known examples of concurrent programming. In this model, some philosophers are sitting on a round table and forks are served between each philosopher. A pasta bawl is also served at the centre of the table, but philosophers have to hold both of left/right forks to help themselves. Here the philosophers are analogues of processes/threads, and the forks are that of shared resources.

philosophers and forks around a table

In a naive implementation of this setting, for example, all philosophers act as following:

    Pick up the fork on the left side
    Pick up the fork on the right side
    Eat the pasta
    Put down the fork on the right side
    Put down the fork on the left side

When multiple philosophers act like this concurrently, as you noticed, it results in a deadlock. Let's model the situation and detect the deadlocked state by this package.

As the simplest case, assume that only two philosophers sitting on the table. We define two processes P1, P2 to represent the philosophers, and two shared variables f1, f2 for forks. The fork f1 is on P1's left side, and the f2 is on his right side.

The red color show you an error trace from the initial state (blue) to a deadlock (red.) In the error firstly P1 gets f1 (P1.up_l) then P2 gets f2 (P2.up_l.) At the deadlock, P1 waits f2 and P2 waits f1 respectively forever.



GO-ROUTINE LEAK:

Goroutines are light weight threads managed by Go runtime.

DESCRIPTION: Go-routine leak is acheived by following instructions below.

Leak Detection.

The formula for detecting leaks in webservers is to add instrumentation endpoints and use them alongside load tests.

[// get the count of number of go routines in the system.
func countGoRoutines() int {
        return runtime.NumGoroutine()
}      

func getGoroutinesCountHandler(w http.ResponseWriter, r *http.Request) {
        // Get the count of number of go routines running.
        count := countGoRoutines()
        w.Write([]byte(strconv.Itoa(count)))
}
func main()
   http.HandleFunc("/_count", getGoroutinesCountHandler)
}]

Use instrumentation endpoint which responds with number of goroutines alive in the system before and after your load test.

Here is the flow of your load test program:

Step 1: Call the instrumentation endpoint and get the count of number of goroutines alive in your webserver.
Step 2: Perform load test.Lets the load be concurrent. 
     for i := 0; i < 100 ; i++ {
          go callEndpointUnderInvestigation()
     }
Step 3: Call the instrumentation endpoint and get the count of number of goroutines alive in your webserver.

There is an evidence of existence of leak if there is an unusual increase in number of goroutines alive in the system after the load test.

Here is a small example with webserver having a leaky endpoint. With a simple test we figure identify the existence of leak the server.

[package main

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
}]

[package main

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
]
// First run the leaky server 
$ go run leaky-server.go
// Run the load test now.
$ go run load.go
 3 Go routines before the load test in the system.
 54 Go routines after the load test in the system.

You can clearly see that with 50 concurrent request to the leaky endpoint there’s a increase of 50 go routines in the system.

Lets run the load test again.

$ go run load.go
 53 Go routines before the load test in the system.
 104 Go routines after the load test in the system.

Its clear that with every run of the load test the number of go routines in the server is increasing and its not dipping down. That’s a clear evidence of a leak.
Identifying the origin of leaks.
Using stack trace instrumentation.

Once you’ve identified that the leaks exist in your web server, now you need to identify the origin of the leak.

Adding endpoint which would return the stack trace of your webserver can help you identify the origin of the leak.
[
import (
  "runtime/debug"
  "runtime/pprof"
)
func getStackTraceHandler(w http.ResponseWriter, r *http.Request) {
       stack := debug.Stack()
       w.Write(stack)
       pprof.Lookup("goroutine").WriteTo(w, 2)
}
func main() {
http.HandleFunc("/_stack", getStackTraceHandler)
}
]
After you identify the existence of leaks, use the endpoint to obtain the stack trace before and after your load to identify the origin of the leak.

Adding the stack trace instrumentation to the leaky-server and performing the load test again. Here is the code:

[package main

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
	"runtime/debug"
	"runtime/pprof"
	"strconv"
)

// get the count of number of go routines in the system.
func countGoRoutines() int {
	return runtime.NumGoroutine()
}

// respond with number of go routines in the system.
func getGoroutinesCountHandler(w http.ResponseWriter, r *http.Request) {
	// Get the count of number of go routines running.
	count := countGoRoutines()
	w.Write([]byte(strconv.Itoa(count)))
}

// respond with the stack trace of the system.
func getStackTraceHandler(w http.ResponseWriter, r *http.Request) {
	stack := debug.Stack()
	w.Write(stack)
	pprof.Lookup("goroutine").WriteTo(w, 2)
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
	x := <-c1
	// write the response.
	fmt.Fprintf(w, strconv.Itoa(x))
}

func main() {
	// get the sum of numbers.
	http.HandleFunc("/sum", sumConcurrent)
	// get the count of number of go routines in the system.
	http.HandleFunc("/_count", getGoroutinesCountHandler)
	// respond with the stack trace of the system.
	http.HandleFunc("/_stack", getStackTraceHandler)
	err := http.ListenAndServe(":8001", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
]
[package main

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

// obtain stack trace of the server.
func getStackTrace() (string, error) {
	body, err := getReq("/_stack")

	if err != nil {
		return "", err
	}
	return string(body), nil
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
	// obtain the stack trace of the system.
	trace, err := getStackTrace()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("\nStack trace after the load test : \n %s",trace)
}
]

//output
// First run the leaky server
$ go run leaky-server.go
// Run the load test now.
$ go run load.go
 3 Go routines before the load test in the system.
 54 Go routines after the load test in the system.
 goroutine 149 [chan send]:
main.sum(0xc420122e58, 0x3, 0x3, 0xc420112240)
        /home/karthic/gophercon/count-instrument.go:39 +0x6c
created by main.sumConcurrent
        /home/karthic/gophercon/count-instrument.go:51 +0x12b

goroutine 243 [chan send]:
main.sum(0xc42021a0d8, 0x3, 0x3, 0xc4202760c0)
        /home/karthic/gophercon/count-instrument.go:39 +0x6c
created by main.sumConcurrent
        /home/karthic/gophercon/count-instrument.go:51 +0x12b

goroutine 259 [chan send]:
main.sum(0xc4202700d8, 0x3, 0x3, 0xc42029c0c0)
        /home/karthic/gophercon/count-instrument.go:39 +0x6c
created by main.sumConcurrent
        /home/karthic/gophercon/count-instrument.go:51 +0x12b

goroutine 135 [chan send]:
main.sum(0xc420226348, 0x3, 0x3, 0xc4202363c0)
        /home/karthic/gophercon/count-instrument.go:39 +0x6c
created by main.sumConcurrent
        /home/karthic/gophercon/count-instrument.go:51 +0x12b

goroutine 166 [chan send]:
main.sum(0xc4202482b8, 0x3, 0x3, 0xc42006b8c0)
        /home/karthic/gophercon/count-instrument.go:39 +0x6c
created by main.sumConcurrent
        /home/karthic/gophercon/count-instrument.go:51 +0x12b

goroutine 199 [chan send]:
main.sum(0xc420260378, 0x3, 0x3, 0xc420256480)
        /home/karthic/gophercon/count-instrument.go:39 +0x6c
created by main.sumConcurrent
        /home/karthic/gophercon/count-instrument.go:51 +0x12b
........

The stack trace clearly points to the epi-center of the leak.