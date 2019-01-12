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
}
