package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	t.Run("sum of 2 and 2 is 4", func(t *testing.T) {
		got := Add(2, 2)
		want := 4

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("sum of 3 and 3 is 6", func(t *testing.T) {
		got := Add(3, 3)
		want := 6

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
