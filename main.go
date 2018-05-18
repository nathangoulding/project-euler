package main

import (
	"flag"
	"os"
	"fmt"
	"github.com/nathangoulding/project-euler/problems"
	"time"
)

var (
	all     bool
	run     bool
	test    bool
	problem int
)

func init() {

	flag.BoolVar(&run, "r", false, "Run a problem")
	flag.BoolVar(&test, "t", false, "Test a problem")
	flag.IntVar(&problem, "p", 0, "The problem to run/test")

	flag.BoolVar(&all, "a", false, "Test all problems")
	flag.Parse()
}

func main() {
	if !all && (problem == 0 || (!run && !test)) {
		flag.Usage()
		if problem == 0 {
			fmt.Print("Must specify a problem to run/test, or test all!\n")
		} else {
			fmt.Print("Must either run/test a problem!\n")
		}
		os.Exit(1)
	}
	if problem != 0 {
		if p, err := problems.GetProblem(problem); err == nil {
			result := timedRun(p, true)
			if test {
				if result != p.Test() {
					fmt.Print("Test failed!\n")
					os.Exit(1)
				}
				fmt.Print("Test passed!\n")
			}
		} else {
			panic(fmt.Sprintf("Invalid problem %s\n", problem))
		}
	} else {
		numFailed := 0
		for _, p := range problems.GetProblems() {
			fmt.Printf("Testing Problem #%d: %s...", p.Number(), p.Description())
			result := timedRun(p, false)
			if result != p.Test() {
				numFailed += 1
				fmt.Print("Failed!\n")
			} else {
				fmt.Print("Passed!\n")
			}
		}
		if numFailed > 0 {
			fmt.Printf("Warning! %d tests failed!\n", numFailed)
			os.Exit(1)
		} else {
			fmt.Print("Congrats! All tests passed!\n")
		}
	}
}

func timedRun(p problems.Problem, printResults bool) int {
	start := time.Now()
	result := p.Run()
	if printResults {
		fmt.Printf("Found: %d\n", result)
		fmt.Printf("Run time: %s\n", time.Since(start))
	}
	return result
}
