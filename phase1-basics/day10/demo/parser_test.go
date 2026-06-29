package main

import (
	"context"
	"testing"
)

func TestParseStudentsConcurrent(t *testing.T) {
	input := `
Tom,90|80
Jerry,78|82
`

	students, err := ParseStudentsConcurrent(context.Background(), input)
	if err != nil {
		t.Fatalf("ParseStudentsConcurrent unexpected error: %v", err)
	}
	if len(students) != 2 {
		t.Fatalf("len(students) = %d, want 2", len(students))
	}
	if students[0].Name != "Tom" || students[1].Name != "Jerry" {
		t.Fatalf("students order changed: %+v", students)
	}
}
