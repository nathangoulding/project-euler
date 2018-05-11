package main

import (
	"flag"
	"os"
	"fmt"
	"github.com/nathangoulding/project-euler/problems"
	"time"
)

var (
	run   bool
	test  bool
	problem string
)

func init() {

	// serve flags
	flag.BoolVar(&run, "r", false, "Run a problem")

	// authenticate flags
	flag.BoolVar(&test, "t", false, "Test a problem")
	flag.StringVar(&problem, "p", "", "The problem to run/test")
	flag.Parse()
}

func main() {
	if problem == "" || (!run && !test) {
		flag.Usage()
		if problem == "" {
			fmt.Print("Must specify a problem to run/test!\n")
		} else {
			fmt.Print("Must either run/test a problem!\n")
		}
		os.Exit(1)
	}
	if p, err := problems.GetProblem(problem); err == nil {
		if run {
			start := time.Now()
			p.Run()
			fmt.Printf("Run time: %s\n", time.Since(start))
		} else {
			p.Test()
		}
	} else {
		panic(fmt.Sprintf("Invalid problem %s\n", problem))
	}
}
