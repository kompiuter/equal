package main

import (
	"fmt"

	"github.com/kompiuter/equal"
)

type Invoice struct {
	amount     float64
	customerID string
}

func (i Invoice) Equals(b equal.Interface) bool {
	return i.customerID == b.(Invoice).customerID
}

func main() {

	inv1 := Invoice{55.5, "George"}
	inv2 := Invoice{28, "John"}
	inv3 := Invoice{100, "George"}
	invoices := []Invoice{inv1, inv2, inv3}
	fmt.Println(equal.Contains(invoices, inv3))

	slice := []string{"hello", "world", "gopher"}
	str := "gopher"
	fmt.Println(equal.StringContains(slice, str))
}
