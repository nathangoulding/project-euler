package problems

import (
	"fmt"
	"errors"
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

func RunProblem(name string) {
	if f, ok := problems[name]; ok {
		f.Run()
		return
	}
	panic(fmt.Sprintf("Invalid problem %s specific", name))
}

func TestProblem(name string) {
	if f, ok := problems[name]; ok {
		f.Test()
		return
	}
	panic(fmt.Sprintf("Invalid problem %s specific", name))
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
