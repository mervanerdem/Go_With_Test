package main

import "testing"

func TestAdder(t *testing.T) {

	t.Run("Adding test", func(t *testing.T) {
		sum := Add(2, 2)
		expected := 4

		if sum != expected {
			t.Errorf("expected: %d but got: %d", expected, sum)
		}
	})
	t.Run("Sub Test", func(t *testing.T) {
		sub := Sub(5, 4)
		expected := 1

		if sub != expected {
			t.Errorf("expected: %d but got: %d", expected, sub)
		}
	})

}
