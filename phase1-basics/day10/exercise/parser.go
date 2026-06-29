package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func splitStudentLines(input string) ([]string, error) {
	text := strings.TrimSpace(input)
	if text == "" {
		return nil, errors.New("students input is empty")
	}

	rawLines := strings.Split(text, "\n")
	lines := make([]string, 0, len(rawLines))
	for _, line := range rawLines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		lines = append(lines, line)
	}
	return lines, nil
}

func ParseStudent(line string) (Student, error) {
	parts := strings.Split(line, ",")
	if len(parts) != 2 {
		return Student{}, fmt.Errorf("invalid student line %q, want name,scores", line)
	}

	name := strings.TrimSpace(parts[0])
	if name == "" {
		return Student{}, errors.New("student name is empty")
	}

	scores, err := parseScores(parts[1])
	if err != nil {
		return Student{}, err
	}
	return Student{Name: name, Scores: scores}, nil
}

func parseScores(input string) ([]int, error) {
	parts := strings.Split(input, "|")
	scores := make([]int, 0, len(parts))
	for _, part := range parts {
		score, err := parseScore(part)
		if err != nil {
			return nil, err
		}
		scores = append(scores, score)
	}
	return scores, nil
}

func parseScore(raw string) (int, error) {
	text := strings.TrimSpace(raw)
	if text == "" {
		return 0, errors.New("score is empty")
	}

	score, err := strconv.Atoi(text)
	if err != nil {
		return 0, fmt.Errorf("invalid score %q", raw)
	}
	if score < 0 || score > 100 {
		return 0, fmt.Errorf("score %d out of range 0-100", score)
	}
	return score, nil
}
