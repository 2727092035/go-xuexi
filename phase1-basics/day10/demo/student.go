package main

type Student struct {
	Name   string
	Scores []int
}

func (s Student) Average() float64 {
	if len(s.Scores) == 0 {
		return 0
	}

	total := 0
	for _, score := range s.Scores {
		total += score
	}
	return float64(total) / float64(len(s.Scores))
}
