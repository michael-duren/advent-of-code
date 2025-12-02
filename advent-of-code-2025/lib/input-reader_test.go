package lib

import (
	"reflect"
	"testing"
)

func Test_ReadInput(t *testing.T) {
	tests := []struct {
		name string
		// function input
		day     string
		want    []string
		wantErr bool
	}{
		{
			name: "correctly reads day one",
			day:  "day1-s",
			want: []string{
				"L68",
				"L30",
				"R48",
				"L5",
				"R60",
				"L55",
				"L1",
				"L99",
				"R14",
				"L82",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := readInput(tt.day)

			// Check result
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadInput() = %v, want %v", got, tt.want)
			}
		})

	}
}
