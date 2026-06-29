package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func ParseScore(raw string) (int, error) {
	text := strings.TrimSpace(raw)
	if text == "" {
		return 0, errors.New("score is empty")
	}

	score, err := strconv.Atoi(text)
	if err != nil {
		return 0, fmt.Errorf("invalid score %q", raw)
	}

	if err := ValidateScore(score); err != nil {
		return 0, err
	}
	return score, nil
}

func ValidateScore(score int) error {
	if score < 0 || score > 100 {
		return fmt.Errorf("score %d out of range 0-100", score)
	}
	return nil
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

func MaxAndCount(nums []int) (int, int) {
	if len(nums) == 0 {
		return 0, 0
	}

	max := nums[0]
	count := 0
	for _, num := range nums {
		if num > max {
			max = num
			count = 1
			continue
		}
		if num == max {
			count++
		}
	}
	return max, count
}
