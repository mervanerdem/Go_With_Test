package main

import "testing"

func TestAdder(t *testing.T) {

	assertCorrectMessage := func(t testing.TB, got, want float64) {
		t.Helper()
		if got != want {
			t.Errorf("got %f want %f", got, want)
		}
	}

	t.Run("Adding test", func(t *testing.T) {
		sum := Add(2, 2)
		expected := 4.0

		assertCorrectMessage(t, sum, expected)
	})
	t.Run("Sub Test", func(t *testing.T) {
		sub := Sub(5, 4)
		expected := 1.0

		assertCorrectMessage(t, sub, expected)
	})
	t.Run("Div Test", func(t *testing.T) {
		div, _ := Div(20, 4)
		expected := 5.0

		assertCorrectMessage(t, div, expected)
	})
	t.Run("Zero can not be divided Test", func(t *testing.T) {
		_, err := Div(20, 0)

		if err != nil {
			t.Errorf("0 can not be divided")
		}
	})

	t.Run("Multiplication Test", func(t *testing.T) {
		div := Mult(3, 4)
		expected := 12.0

		assertCorrectMessage(t, div, expected)
	})

}
