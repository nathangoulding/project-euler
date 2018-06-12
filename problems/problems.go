package problems

import (
	"fmt"
	"errors"
	"sort"
)

var (
	problems = Problems{}
)

type Problem interface {
	Number() int
	Description() string
	Run() int
	Test() int
}

type Problems []Problem

// Ensure it satisfies sort.Interface
func (p Problems) Len() int           { return len(p) }
func (p Problems) Less(i, j int) bool { return p[i].Number() < p[j].Number() }
func (p Problems) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func GetProblem(problem int) (Problem, error) {
	for _, p := range problems {
		if p.Number() == problem {
			return p, nil
		}
	}
	return nil, errors.New("problem does not exist")
}

func GetProblems() Problems {
	return problems
}

func RegisterProblem(problem Problem) {
	if _, err := GetProblem(problem.Number()); err == nil {
		panic(fmt.Sprintf("Problem %d already registered", problem.Number()))
	}
	problems = append(problems, problem)
	sort.Sort(problems)
}