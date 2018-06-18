package problems

import (
	"sort"
	"math"
	"strconv"
	"math/big"
)

func BigFactorial(num int) *big.Int {
	factorial := big.NewInt(1)
	for i := num; i > 1; i-- {
		factorial.Mul(factorial, big.NewInt(int64(i)))
	}
	return factorial
}


func Factorial(num int) uint64 {
	factorial := uint64(1)
	for i := num; i > 1; i-- {
		factorial = uint64(i) * factorial
	}
	return factorial
}

func SumBigIntDigits(num *big.Int) int {
	var digits []int
	for _, value := range num.String() {
		i, _ := strconv.Atoi(string(value))
		digits = append(digits, i)
	}
	return Summation(digits)
}

func PascalsTriangle(rows int) [][]int {
	grid := make([][]int, rows)
	for i := range grid {
		grid[i] = make([]int, i + 1)
	}
	for i := 0; i < rows; i++ {
		for j := range grid[i] {
			if j == 0 || j == i {
				grid[i][j] = 1
			} else {
				grid[i][j] = grid[i - 1][j] + grid[i - 1][j - 1]
			}
		}
	}
	return grid
}

func IsDeficientNumber(num int) bool {
	return Summation(DivisorsOf(num)) < num
}

func IsPerfectNumber(num int) bool {
	return Summation(DivisorsOf(num)) == num
}

func IsAbundantNumber(num int) bool {
	return Summation(DivisorsOf(num)) > num
}

func BinomialCoefficientAsGrid(n int, k int) [][]int {
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

func PadRight(str, pad string, length int) string {
	for {
		str += pad
		if len(str) > length {
			return str[0:length]
		}
	}
}

func PadLeft(str, pad string, length int) string {
	for {
		str = pad + str
		if len(str) > length {
			return str[0:length]
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
	divisors := []int{1}
	for j := 2; j <= int(math.Sqrt(float64(value))); j++ {
		if value % j == 0 {
			if value / j == j {
				divisors = append(divisors, j)
			} else {
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

func AlphabeticValues(word string) []int {
	var values []int
	for _, letter := range word {
		values = append(values, AlphabeticValue(string(letter)))
	}
	return values
}

func AlphabeticValue(letter string) int {
	if letter == "a" || letter == "A" {
		return 1
	}
	if letter == "b" || letter == "B" {
		return 2
	}
	if letter == "c" || letter == "C" {
		return 3
	}
	if letter == "d" || letter == "D" {
		return 4
	}
	if letter == "e" || letter == "E" {
		return 5
	}
	if letter == "f" || letter == "F" {
		return 6
	}
	if letter == "g" || letter == "G" {
		return 7
	}
	if letter == "h" || letter == "H" {
		return 8
	}
	if letter == "i" || letter == "I" {
		return 9
	}
	if letter == "j" || letter == "J" {
		return 10
	}
	if letter == "k" || letter == "K" {
		return 11
	}
	if letter == "l" || letter == "L" {
		return 12
	}
	if letter == "m" || letter == "M" {
		return 13
	}
	if letter == "n" || letter == "N" {
		return 14
	}
	if letter == "o" || letter == "O" {
		return 15
	}
	if letter == "p" || letter == "P" {
		return 16
	}
	if letter == "q" || letter == "Q" {
		return 17
	}
	if letter == "r" || letter == "R" {
		return 18
	}
	if letter == "s" || letter == "S" {
		return 19
	}
	if letter == "t" || letter == "T" {
		return 20
	}
	if letter == "u" || letter == "U" {
		return 21
	}
	if letter == "v" || letter == "V" {
		return 22
	}
	if letter == "w" || letter == "W" {
		return 23
	}
	if letter == "x" || letter == "X" {
		return 24
	}
	if letter == "y" || letter == "Y" {
		return 25
	}
	if letter == "z" || letter == "Z" {
		return 26
	}
	return -65535
}