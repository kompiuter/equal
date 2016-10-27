package equal

import "testing"

func TestStringContains(t *testing.T) {
	var tests = []struct {
		slice []string
		value string
		want  bool
	}{
		{[]string{"hello", "gopher", "golang", "world"}, "golang", true},
		{[]string{"hello", "go", "hello"}, "hello", true},
		{[]string{"hello", ""}, "", true},
		{[]string{"hello", "gopher", "golang", "world"}, "GOLANG", false},
		{[]string{"hello", "gopher", "golang", "world"}, "", false},
		{[]string{}, "", false},
		{nil, "go", false},
	}
	for _, test := range tests {
		if got := StringContains(test.slice, test.value); got != test.want {
			t.Errorf("StringContains(%#v, %#v) = %v", test.slice, test.value, got)
		}
	}
}

func TestIntContains(t *testing.T) {
	var tests = []struct {
		slice []int
		value int
		want  bool
	}{
		{[]int{5, 10, 2, 60, 3, 5}, 5, true},
		{[]int{0, 0, 0, 0}, 0, true},
		{[]int{0, 0, 0, 0}, 5, false},
		{[]int{5, 3, 2, 60, 3, 5}, 60, true},
		{[]int{-1, -2, -3, -4}, -5, false},
		{[]int{-1, -2, -3, -4}, 4, false},
		{[]int{1, 2}, 0, false},
		{[]int{1e6, 1e5}, 1e5, true},
		{[]int{1e6, 1e5}, 1e0, false},
		{[]int{}, 0, false},
		{nil, 0, false},
	}
	for _, test := range tests {
		if got := IntContains(test.slice, test.value); got != test.want {
			t.Errorf("IntContains(%#v, %#v) = %v", test.slice, test.value, got)
		}
	}
}

func TestFloat64Contains(t *testing.T) {
	var tests = []struct {
		slice []float64
		value float64
		want  bool
	}{
		{[]float64{0.001, 5, 0.202, 3.3}, 0.202, true},
		{[]float64{1.53e-3, 2.33e2, 3.1414, 99}, 2.33e2, true},
		{[]float64{0, 0, 0, 0}, 0, true},
		{[]float64{0.001, 0.00001, 1e-5, 0, 10}, 0.0001, false},
		{[]float64{-5.5, -10.3}, 5.5, false},
		{[]float64{0}, 1e-30, false},
		{[]float64{}, 0, false},
		{nil, 1e-9, false},
	}
	for _, test := range tests {
		if got := Float64Contains(test.slice, test.value); got != test.want {
			t.Errorf("Float64Contains(%#v, %#v) = %v", test.slice, test.value, got)
		}
	}
}
