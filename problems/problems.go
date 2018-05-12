package problems

import (
	"fmt"
	"errors"
	"math"
)

var (
	problems = make(map[string]Problem)
)

type Problem interface {
	Run() int
	Test() int
}

func GetProblem(name string) (Problem, error) {
	if problems[name] != nil {
		return problems[name], nil
	}
	return nil, errors.New("problem does not exist")
}

func RegisterProblem(name string, problem Problem) {
	if problems[name] != nil {
		panic(fmt.Sprintf("Problem %s already registered", name))
	}
	problems[name] = problem
}

func IntInSlice(a int, list []int) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func PrimeFactorsOf(value int) []int {
	if value < 2 {
		panic("not a valid number")
	}

	if IsPrime(value) {
		return []int{value}
	}
	sqrt := int(math.Sqrt(float64(value)))
	var primes []int

L1:
	for {
L2:
		for i := 2; i <= sqrt; i++ {
			if value % i == 0 {
				primes = append(primes, i)
				value = value / i
				if IsPrime(value) {
					primes = append(primes, value)
					break L1
				}
				break L2
			}
		}
	}
	return primes
}

func IsPrime(value int) bool {
	if value == 2 {
		return true
	}
	for i := 2; i < value/2; i++ {
		if value % i == 0 {
			return false
		}
	}
	return true
}