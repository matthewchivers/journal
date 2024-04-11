package caltools

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestWeekOfMonth(t *testing.T) {
	tests := []struct {
		name string
		date time.Time
		want int
	}{
		{
			name: "Week 1 - 1st March 2024 (Friday)",
			date: time.Date(2024, time.March, 1, 0, 0, 0, 0, time.UTC),
			want: 1,
		},
		{
			name: "Week 1 - 3rd March 2024 (Sunday)",
			date: time.Date(2024, time.March, 3, 0, 0, 0, 0, time.UTC),
			want: 1,
		},
		{
			name: "Week 2 - 4th March 2024 (Monday)",
			date: time.Date(2024, time.March, 4, 0, 0, 0, 0, time.UTC),
			want: 2,
		},
		{
			name: "Week 5 - 25th March 2024 (Monday)",
			date: time.Date(2024, time.March, 25, 0, 0, 0, 0, time.UTC),
			want: 5,
		},
		{
			name: "Week 5 - 31st March 2024 (Sunday)",
			date: time.Date(2024, time.March, 31, 0, 0, 0, 0, time.UTC),
			want: 5,
		},
		{
			name: "Week 1 - 1st April 2024 (Monday)",
			date: time.Date(2024, time.April, 1, 0, 0, 0, 0, time.UTC),
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := WeekOfMonth(tt.date)
			assert.Equal(t, tt.want, got, "WeekOfMonth() = %v, want %v", got, tt.want)
		})
	}
}

func TestWeekCommencing(t *testing.T) {
	tests := []struct {
		name string
		date time.Time
		want time.Time
	}{
		{
			name: "Week commencing for 26th February 2024 (Monday)",
			date: time.Date(2024, time.February, 26, 0, 0, 0, 0, time.UTC),
			want: time.Date(2024, time.February, 26, 0, 0, 0, 0, time.UTC),
		},
		{
			name: "Week commencing for 1st March 2024 (Friday)",
			date: time.Date(2024, time.March, 1, 0, 0, 0, 0, time.UTC),
			want: time.Date(2024, time.February, 26, 0, 0, 0, 0, time.UTC),
		},
		{
			name: "Week commencing for 3rd March 2024 (Sunday)",
			date: time.Date(2024, time.March, 3, 0, 0, 0, 0, time.UTC),
			want: time.Date(2024, time.February, 26, 0, 0, 0, 0, time.UTC),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := WeekCommencing(tt.date)
			assert.Equal(t, tt.want, got, "WeekCommencing() = %v, want %v", got, tt.want)
		})
	}
}
