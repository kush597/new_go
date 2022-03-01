package main

import (
	"fmt"
	"log"
	"os"

	"github.com/dastergon/oomutil"
)

func main() {
	pid := os.Getpid()
	ps, err := oomutil.NewOOMProcess(int32(pid))
	if err != nil {
		log.Fatal(err)
	}

	oomScore, err := ps.OOMScore()
	if err != nil {
		log.Fatal(err)
	}

	oomScoreAdj, err := ps.OOMScoreAdj()
	if err != nil {
		log.Fatal(err)
	}

	memoryOvercommit, err := ps.MemoryOvercommit()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Memory Overcommit: %d\nPID: %d => OOM Score: %d, OOM Score Adj: %d\n", memoryOvercommit, pid, oomScore, oomScoreAdj)
}
