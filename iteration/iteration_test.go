package iteration

import (
	"testing"
)

// TestRepeat tests Repeat by checking that it returns a string repeated 5 times.
// func TestRepeat(t *testing.T) {
// 	repeated := Repeat("a")
// 	expected := "aaaaa"

// 	if repeated != expected {
// 		t.Errorf("expected %q but got %q", expected, repeated)
// 	}
// }

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}

// func BenchmarkRandInt(b *testing.B) {
// 	max := big.NewInt(15000) // maximum value
// 	for range b.N {
// 		rand.Int(rand.Reader, max)
// 	}
// }
