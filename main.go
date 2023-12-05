package main

import (
	"github.com/ptran-hw/advent-of-code-2023/day1"
	"github.com/ptran-hw/advent-of-code-2023/day2"
	"github.com/ptran-hw/advent-of-code-2023/day3"
	"github.com/ptran-hw/advent-of-code-2023/day4"
	"github.com/ptran-hw/advent-of-code-2023/day5"
	"log"
	"os"
	"time"
)

// use pointer to Solver for mutable instance
var solvers = map[string]Solver{
	"1": day1.Solver{},
	"2": day2.Solver{},
	"3": day3.Solver{},
	"4": day4.Solver{},
	"5": day5.Solver{},
}

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
