package money

import "testing"

func TestRoundUsesTwoDecimalCommercialRounding(t *testing.T) {
	tests := []struct {
		name  string
		value float64
		want  float64
	}{
		{name: "rounds down", value: 10.124, want: 10.12},
		{name: "rounds up", value: 10.125, want: 10.13},
		{name: "rounds negative half away from zero", value: -10.125, want: -10.13},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := Round(test.value); got != test.want {
				t.Fatalf("Round(%v) = %v, want %v", test.value, got, test.want)
			}
		})
	}
}
