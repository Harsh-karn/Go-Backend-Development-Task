package service

import (
	"testing"
	"time"
)

func TestCalculateAge(t *testing.T) {
	tests := []struct {
		name        string
		dobStr      string
		nowStr      string
		expectedAge int
	}{
		{
			name:        "Birthday has passed this year",
			dobStr:      "1990-05-10",
			nowStr:      "2026-08-15",
			expectedAge: 36,
		},
		{
			name:        "Birthday has not passed this year",
			dobStr:      "1990-12-20",
			nowStr:      "2026-08-15",
			expectedAge: 35,
		},
		{
			name:        "Today is the birthday",
			dobStr:      "1990-08-15",
			nowStr:      "2026-08-15",
			expectedAge: 36,
		},
		{
			name:        "Born this year",
			dobStr:      "2026-01-01",
			nowStr:      "2026-08-15",
			expectedAge: 0,
		},
		{
			name:        "Born tomorrow (future date)",
			dobStr:      "2026-08-16",
			nowStr:      "2026-08-15",
			expectedAge: -1, // Our logic returns -1 for future dates in the same year, which is mathematically consistent though practically invalid (we should validate no future dates on input).
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dob, _ := time.Parse("2006-01-02", tt.dobStr)
			now, _ := time.Parse("2006-01-02", tt.nowStr)

			age := CalculateAge(dob, now)
			if age != tt.expectedAge {
				t.Errorf("expected %d, got %d", tt.expectedAge, age)
			}
		})
	}
}
