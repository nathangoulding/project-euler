package problems

import (
	"fmt"
	"errors"
	"math"
	sort "sort"
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

func GreatestCommonDivisorOf(value int) int {
	factors := DivisorsOf(value)
	return factors[len(factors) - 1]
}

func DivisorsOf(value int) []int {
	var divisors []int
	for j := 2; j <= int(math.Sqrt(float64(value))); j++ {
		if value % j == 0 {
			if value / j == j {
				divisors = append(divisors, j)
			} else {
				divisors = append(divisors, j)
				divisors = append(divisors, j, value / j)
			}
		}
	}
	sort.Slice(divisors, func(i, j int) bool {
		return divisors[i]< divisors[j]
	})
	return divisors
}

func GreatestPrimeFactorOf(value int) int {
	factors := PrimeFactorsOf(value)
	return factors[len(factors) - 1]
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
	factors := PrimeFactorsOf(value)
	return len(factors) == 0
	//for i := 2; i <= value/2; i++ {
	//	if value % i == 0 {
	//		return false
	//	}
	//}
	//return true
}

// sieve of eratosthenes
func PrimesBelow(value int) []int {

	multipliers := make([]int, value - 1)
	multipliers[0] = 1
	multipliers[1] = 1

	for i := 2; i <= value / 2; i++ {
		for j := 2; j < value / j; j++ {
			if i * j < value - 1 {
				multipliers[i * j] = 1
			}
		}
	}

	var primes []int
	for i, value := range multipliers {
		if value != 1 {
			primes = append(primes, i)
		}
	}

	return primes
}