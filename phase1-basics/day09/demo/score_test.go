package main

import "testing"

func TestGrade(t *testing.T) {
	tests := []struct {
		name  string
		score float64
		want  string
	}{
		{name: "A boundary", score: 90, want: "A"},
		{name: "B boundary", score: 80, want: "B"},
		{name: "C boundary", score: 70, want: "C"},
		{name: "D boundary", score: 60, want: "D"},
		{name: "F below 60", score: 59.9, want: "F"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Grade(tt.score)
			if got != tt.want {
				t.Fatalf("Grade(%v) = %q, want %q", tt.score, got, tt.want)
			}
		})
	}
}

func TestParseScore(t *testing.T) {
	tests := []struct {
		name    string
		raw     string
		want    int
		wantErr bool
	}{
		{name: "valid", raw: "90", want: 90},
		{name: "trim spaces", raw: " 80 ", want: 80},
		{name: "invalid number", raw: "abc", wantErr: true},
		{name: "negative", raw: "-1", wantErr: true},
		{name: "too large", raw: "101", wantErr: true},
		{name: "empty", raw: "", wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseScore(tt.raw)
			if tt.wantErr {
				if err == nil {
					t.Fatalf("ParseScore(%q) expected error", tt.raw)
				}
				return
			}
			if err != nil {
				t.Fatalf("ParseScore(%q) unexpected error: %v", tt.raw, err)
			}
			if got != tt.want {
				t.Fatalf("ParseScore(%q) = %d, want %d", tt.raw, got, tt.want)
			}
		})
	}
}

func TestMaxAndCount(t *testing.T) {
	tests := []struct {
		name      string
		nums      []int
		wantMax   int
		wantCount int
	}{
		{name: "empty", nums: nil, wantMax: 0, wantCount: 0},
		{name: "single", nums: []int{7}, wantMax: 7, wantCount: 1},
		{name: "two max values", nums: []int{90, 100, 80, 100}, wantMax: 100, wantCount: 2},
		{name: "all same", nums: []int{5, 5, 5}, wantMax: 5, wantCount: 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotMax, gotCount := MaxAndCount(tt.nums)
			if gotMax != tt.wantMax || gotCount != tt.wantCount {
				t.Fatalf("MaxAndCount(%v) = (%d, %d), want (%d, %d)",
					tt.nums, gotMax, gotCount, tt.wantMax, tt.wantCount)
			}
		})
	}
}

func BenchmarkGrade(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Grade(88.5)
	}
}
