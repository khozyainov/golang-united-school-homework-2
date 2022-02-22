package square

import (
	"math"
	"testing"
)

func TestCalcSquare(t *testing.T) {
	side_len := 10.0
	triangle_expected := side_len * side_len / 2
	triangle_result := CalcSquare(side_len, SidesTriangle)
	if triangle_result != triangle_expected {
		t.Errorf("Got %f, expected %f", triangle_result, triangle_expected)
	}

	square_expected := side_len * side_len
	square_result := CalcSquare(side_len, SidesSquare)
	if square_result != square_expected {
		t.Errorf("Got %f, expected %f", square_result, square_expected)
	}

	circle_expected := side_len * side_len * math.Pi
	circle_result := CalcSquare(side_len, SidesCircle)
	if circle_result != circle_expected {
		t.Errorf("Got %f, expected %f", circle_result, circle_expected)
	}
}
