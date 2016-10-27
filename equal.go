package equal

// Interface is a type, typically a struct, that satisfies equal.Interface
// and can be tested for equality with another type that satisfies
// equal.Interface by routines in this package.
type Interface interface {
	// Equals reports whether the element is equal to data.
	Equals(data Interface) bool
}

// Contains reports whether the provided slice contains the
// value passed.
//
// Passing a nil slice or a nil value always returns false.
func Contains(slice []interface{}, value Interface) bool {
	if slice == nil || value == nil {
		return false
	}
	for i := 0; i < len(slice); i++ {
		eq := slice[i].(Interface)
		if eq.Equals(value) {
			return true
		}
	}
	return false
}

// Convenience types for common cases

// String attaches the methods of Interface to primitive string.
type String string

func (a String) Equals(b Interface) bool { return a == b }

// Int attaches the methods of Interface to primitive int.
type Int int

func (a Int) Equals(b Interface) bool { return a == b }

// Float64 attaches the methods of Interface to primitive float64.
type Float64 float64

func (a Float64) Equals(b Interface) bool { return a == b }

// Convenience wrappers for common cases

// StringContains reports whether a slice of strings contains the string b.
func StringContains(a []string, b string) bool {
	slice := make([]interface{}, len(a))
	for i, s := range a {
		slice[i] = String(s)
	}
	return Contains(slice, String(b))
}

// IntContains reports whether a slice of ints contains the int b.
func IntContains(a []int, b int) bool {
	slice := make([]interface{}, len(a))
	for i, s := range a {
		slice[i] = Int(s)
	}
	return Contains(slice, Int(b))
}

// Float64Contains reports whether a slice of float64s contains the float64 b.
func Float64Contains(a []float64, b float64) bool {
	slice := make([]interface{}, len(a))
	for i, s := range a {
		slice[i] = Float64(s)
	}
	return Contains(slice, Float64(b))
}
