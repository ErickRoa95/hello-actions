package calculator

import "testing"

type testCase struct{
	name string
	a,b int
	want float32
}

func TestCalculator(t *testing.T) {
	testCases := []testCase{
		{
			name: "ADD",
			a: 10,
			b: 5,
			want: 15,
		},
		{
			name: "SUBTRACT",
			a: 10,
			b: 5,
			want: 5,
		},
		{
			name: "MULTIPLY",
			a: 10,
			b: 5,
			want: 50.0,
		},
		{
			name: "DIVIDE",
			a: 10,
			b: 5,
			want: 2.0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
				calc := NewCalculator()
				switch tc.name {
				case "ADD":
					if got := calc.Add(tc.a, tc.b); got != int(tc.want) {
						t.Errorf("Add(%d, %d) = %d; want %d", tc.a, tc.b, got, int(tc.want))
					}
					case "SUBTRACT":
					if got := calc.Subtract(tc.a, tc.b); got != int(tc.want) {
						t.Errorf("Subtract(%d, %d) = %d; want %d", tc.a, tc.b, got, int(tc.want))
					}
				case "MULTIPLY":
					if got := calc.Multiply(tc.a, tc.b); got != tc.want {
						t.Errorf("Multiply(%d, %d) = %.2f; want %.2f", tc.a, tc.b, got, tc.want)
					}
				case "DIVIDE":
					if got := calc.Divide(tc.a, tc.b); got != tc.want {
						t.Errorf("Divide(%d, %d) = %.2f; want %.2f", tc.a, tc.b, got, tc.want)
				}
			}
		})
	}
}