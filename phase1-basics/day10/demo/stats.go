package main

import "errors"

type ClassStats struct {
	Average float64
	Highest int
	Lowest  int
}

func CalculateStats(students []Student) (ClassStats, error) {
	if len(students) == 0 {
		return ClassStats{}, errors.New("students is empty")
	}

	total := 0
	count := 0
	highest := 0
	lowest := 0
	hasScore := false

	for _, student := range students {
		for _, score := range student.Scores {
			if !hasScore {
				highest = score
				lowest = score
				hasScore = true
			}
			total += score
			count++
			if score > highest {
				highest = score
			}
			if score < lowest {
				lowest = score
			}
		}
	}

	if count == 0 {
		return ClassStats{}, errors.New("scores is empty")
	}

	return ClassStats{
		Average: float64(total) / float64(count),
		Highest: highest,
		Lowest:  lowest,
	}, nil
}

func Grade(score float64) string {
	switch {
	case score >= 90:
		return "A"
	case score >= 80:
		return "B"
	case score >= 70:
		return "C"
	case score >= 60:
		return "D"
	default:
		return "F"
	}
}

func GradeCounts(students []Student) map[string]int {
	counts := map[string]int{"A": 0, "B": 0, "C": 0, "D": 0, "F": 0}
	for _, student := range students {
		counts[Grade(student.Average())]++
	}
	return counts
}
