package main

import (
	"context"
	"fmt"
)

type parseJob struct {
	Index int
	Line  string
}

type parseResult struct {
	Index   int
	Student Student
	Err     error
}

func ParseStudentsConcurrent(ctx context.Context, input string) ([]Student, error) {
	lines, err := splitStudentLines(input)
	if err != nil {
		return nil, err
	}

	jobs := make(chan parseJob)
	results := make(chan parseResult)

	workerCount := 2
	for i := 0; i < workerCount; i++ {
		go parseWorker(ctx, jobs, results)
	}

	go func() {
		defer close(jobs)
		for index, line := range lines {
			select {
			case <-ctx.Done():
				return
			case jobs <- parseJob{Index: index, Line: line}:
			}
		}
	}()

	students := make([]Student, len(lines))
	for i := 0; i < len(lines); i++ {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case result := <-results:
			if result.Err != nil {
				return nil, fmt.Errorf("line %d invalid: %w", result.Index+1, result.Err)
			}
			students[result.Index] = result.Student
		}
	}
	return students, nil
}

func parseWorker(ctx context.Context, jobs <-chan parseJob, results chan<- parseResult) {
	for {
		select {
		case <-ctx.Done():
			return
		case job, ok := <-jobs:
			if !ok {
				return
			}

			student, err := ParseStudent(job.Line)
			select {
			case <-ctx.Done():
				return
			case results <- parseResult{Index: job.Index, Student: student, Err: err}:
			}
		}
	}
}
