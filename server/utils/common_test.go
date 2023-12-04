package utils

import (
	"testing"
	"time"
)

func TestIsWithinBusinessHours(t *testing.T) {
	tests := []struct {
		name      string
		inputTime string
		start     string
		end       string
		expected  bool
	}{
		{"WithinBusinessHours", "2023-12-01T10:30:45Z", "08:00:00", "17:00:00", true},
		{"OutsideBusinessHours", "2023-12-01T22:52:45Z", "08:00:00", "22:52:00", false},
		{"OnWeekend", "2023-12-02T10:30:45Z", "08:00:00", "17:00:00", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inputTime, err := time.Parse(time.RFC3339, tt.inputTime)
			if err != nil {
				t.Fatalf("failed to parse input time: %v", err)
			}

			result := IsWithinBusinessHours(inputTime, tt.start, tt.end)

			if result != tt.expected {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}
