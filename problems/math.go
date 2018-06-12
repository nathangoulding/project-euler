package problems

import (
	"sort"
	"math"
)

func BinomialCoefficient(n int, k int) [][]int {
	grid := make([][]int, n)
	for i := range grid {
		grid[i] = make([]int, k)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < k; j++ {
			if i == 0 || j == 0 {
				grid[i][j] = 1
			} else {
				grid[i][j] = grid[i - 1][j] + grid[i][j - 1]
			}
		}
	}
	return grid
}

func PadRight(str, pad string, lenght int) string {
	for {
		str += pad
		if len(str) > lenght {
			return str[0:lenght]
		}
	}
}

func PadLeft(str, pad string, lenght int) string {
	for {
		str = pad + str
		if len(str) > lenght {
			return str[0:lenght]
		}
	}
}

func Summation(values []int) int {
	sum := 0
	for _, value := range values {
		sum += value
	}
	return sum
}

func MakeRange(min int, max int) []int {
	var set []int
	for i := min; i <= max; i++ {
		set = append(set, i)
	}
	return set
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
		return divisors[i] < divisors[j]
	})
	return divisors
}

func GreatestPrimeFactorOf(value int) int {
	factors := PrimeFactorsOf(value)
	return factors[len(factors) - 1]
}

func PrimeFactorsOf(value int) []int {
	if value < 2 {
		panic("not valid input")
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
	for i := 2; i <= value/2; i++ {
		if value % i == 0 {
			return false
		}
	}
	return true
}

// sieve of eratosthenes
func PrimesBelow(value int64) []int {

	multipliers := make([]int, value - 1)
	multipliers[0] = 1
	multipliers[1] = 1

	for i := int64(2); i <= value / 2; i++ {
		for j := int64(2); j < value / j; j++ {
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
