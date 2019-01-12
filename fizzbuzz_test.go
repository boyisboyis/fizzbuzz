package fizzbuzz_test

import (
	"testing"

	"github.com/boyisboyis/fizzbuzz"
)

func TestFizzBuzz(t *testing.T) {
	t.Run("1 should say 1", func(t *testing.T) {
		get := fizzbuzz.Say(1)

		want := "1"

		if get != want {
			t.Errorf("it should say %s but get %s", get, want)
		}
	})

	t.Run("2 should say 2", func(t *testing.T) {
		get := fizzbuzz.Say(2)

		want := "2"

		if get != want {
			t.Errorf("it should say %s but get %s", get, want)
		}
	})

	t.Run("3 should say Fizz", func(t *testing.T) {
		get := fizzbuzz.Say(3)

		want := "Fizz"

		if get != want {
			t.Errorf("it should say %s but get %s", get, want)
		}
	})

	t.Run("4 should say 4", func(t *testing.T) {
		get := fizzbuzz.Say(4)

		want := "4"

		if get != want {
			t.Errorf("it should say %s but get %s", get, want)
		}
	})

	t.Run("5 should say Buzz", func(t *testing.T) {
		get := fizzbuzz.Say(5)

		want := "Buzz"

		if get != want {
			t.Errorf("it should say %s but get %s", get, want)
		}
	})
}
