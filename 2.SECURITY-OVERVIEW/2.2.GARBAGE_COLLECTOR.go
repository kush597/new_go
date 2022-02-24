package main


import (
	"fmt"
	"runtime"
)

func printStats(mem runtime.MemStats) {
	runtime.ReadMemStats(&mem)
	fmt.Println(" Mem Alloc : ", mem.Alloc)
	fmt.Println(" Mem TotalAlloc : ", mem.TotalAlloc)
	fmt.Println(" Mem HeapAlloc : ", mem.HeapAlloc)
	fmt.Println(" Mem NumGC : ", mem.NumGC,"\n")
}

func main() {
	var mem runtime.MemStats
	printStats(mem)
	for i := 0; i<10; i++ {
		s := make([]byte,5000000)
		if s == nil {
			fmt.Println("operation failed")
		}
	}
	printStats(mem)
}