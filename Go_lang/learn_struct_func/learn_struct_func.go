package main

import (
	"fmt"
	"strconv"
)

type Student struct {
	name   string
	grades []int
	age    int
}

func (s *Student) getAverageGrade() float32 {
	sum := 0
	for _, v := range s.grades {
		sum += v
	}
	return float32(sum) / float32(len(s.grades))
}

func (s *Student) getMaxGrade() int {
	best := 0
	for _, v := range s.grades {
		if best < v {
			best = v
		}
	}
	return best
}

func (s *Student) getInformation() string {
	return "Name: " + s.name + ", Age: " + strconv.FormatInt(int64(s.age), 10)
}

func main() {
	s1 := Student{"Joe", []int{80, 85, 40, 24, 32}, 20}
	s2 := Student{"Mike", []int{20, 99, 100, 30, 80, 81, 76, 65}, 22}
	fmt.Printf("%v,  Average grade: %v, Max Grade: %v \n", s1.getInformation(), s1.getAverageGrade(), s1.getMaxGrade())
	fmt.Printf("%v,  Average grade: %v, Max Grade: %v \n", s2.getInformation(), s2.getAverageGrade(), s2.getMaxGrade())
}
