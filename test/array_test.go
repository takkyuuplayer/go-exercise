package test

import "testing"

func TestArrayInitializationWithIndex(t *testing.T) {
	type Currency int

	const (
		USD Currency = iota
		EUR
		JPY
	)

	symbol := [...]string{
		USD: "$",
		EUR: "EURO",
		JPY: "円",
	}

	if JPY != 2 {
		t.Errorf(`JPY = %#v, want %#v`, JPY, 2)
	}

	if symbol[JPY] != "円" {
		t.Errorf(`symbol[JPY] = %#v, want %#v`, symbol[JPY], "円")
	}
}

func TestArrayInitializationWithSkippingIndex(t *testing.T) {
	r := [...]int{99: -1}

	if len(r) != 100 {
		t.Errorf(`len(r) = %#v, want %#v`, len(r), 100)
	}

	if r[0] != 0 {
		t.Errorf(`r[0] = %#v, want %#v`, r[0], 0)
	}

	if r[99] != -1 {
		t.Errorf(`r[99] = %#v, want %#v`, r[99], -1)
	}
}
