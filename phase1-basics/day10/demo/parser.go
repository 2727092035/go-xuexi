package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func ParseStudents(input string) ([]Student, error) {
	lines, err := splitStudentLines(input)
	if err != nil {
		return nil, err
	}

	students := make([]Student, 0, len(lines))
	for index, line := range lines {
		student, err := ParseStudent(line)
		if err != nil {
			return nil, fmt.Errorf("line %d invalid: %w", index+1, err)
		}
		students = append(students, student)
	}
	return students, nil
}

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
	if len(lines) == 0 {
		return nil, errors.New("students input is empty")
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
	if strings.TrimSpace(input) == "" {
		return nil, errors.New("scores is empty")
	}

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
	if err := validateScore(score); err != nil {
		return 0, err
	}
	return score, nil
}

func validateScore(score int) error {
	if score < 0 || score > 100 {
		return fmt.Errorf("score %d out of range 0-100", score)
	}
	return nil
}
