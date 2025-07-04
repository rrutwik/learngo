package structs

import (
	"math"
	"testing"
)

func TestPerimeter(t *testing.T) {

	t.Run("rectangle", func(t *testing.T) {
		got := Rectangle{10.0, 10.0}.Perimeter()
		want := 40.0

		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})

	t.Run("circle", func(t *testing.T) {
		got := Circle{10.0}.Perimeter()
		want := 2 * math.Pi * 10
		// compare float got and want upto 2 digits
		if math.Abs(got-want) > 0.01 {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})
}

func TestArea(t *testing.T) {
	t.Run("rectangle", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		got := rectangle.Area()
		want := 100.0

		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})

	t.Run("circle", func(t *testing.T) {
		circle := Circle{10.0}
		got := circle.Area()
		want := math.Pi * 10 * 10
		// compare float got and want upto 2 digits
		if math.Abs(got-want) > 0.01 {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})
}

func TestArea2(t *testing.T) {

	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if math.Abs(got-want) > 0.01 {
			t.Errorf("got %g want %g", got, want)
		}
	}

	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{12, 6}
		checkArea(t, rectangle, 72.0)
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{10}
		checkArea(t, circle, math.Pi*10*10)
	})

}
