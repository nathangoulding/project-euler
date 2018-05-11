package problems

import (
	"fmt"
)

func init() {
	RegisterProblem("12", Problem12{})
}

type Problem12 struct {
}

func (p Problem12) Run() {
	desiredDivisors := 350
	fmt.Printf("Looking for %d divisors\n", desiredDivisors)

	triangle := 0
	for i := 1; i > 0; i++ {
		triangle += i
		divisors := 0
		for j := 1; j <= triangle; j++ {
			if triangle % j == 0 {
				divisors += 1
			}
		}
		if divisors > desiredDivisors {
			break
		}
	}
	fmt.Printf("Found triangle number: %d\n", triangle)
}

func (p Problem12) Test() {
	fmt.Print("there")
}
