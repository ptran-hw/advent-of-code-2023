package main

import (
	"log"
	"os"
	"time"
)

// use pointer to Solver for mutable instance
var solvers = map[string]Solver{}

type Solver interface {
	Solve()
}

func main() {
	arguments := os.Args[1:]

	if len(arguments) != 1 {
		log.Panic("incorrect number of arguments used")
	}

	problemNumber := arguments[0]
	solver := solvers[problemNumber]
	if solver == nil {
		log.Panicf("unable to find day %s solver", problemNumber)
	}

	start := time.Now()
	defer func() { log.Printf("time elapsed: %v\n", time.Now().Sub(start)) }()

	solver.Solve()
}
