package square

import (
	"math"
	"testing"
)

const errorMsg = "Got %f, expected %f"

func TestCalcSquare(t *testing.T) {
	sideLen := 10.0
	triangleExpected := sideLen * sideLen * math.Sqrt(3) / 4
	triangleResult := CalcSquare(sideLen, SidesTriangle)
	if triangleResult != triangleExpected {
		t.Errorf(errorMsg, triangleResult, triangleExpected)
	}

	squareExpected := sideLen * sideLen
	squareResult := CalcSquare(sideLen, SidesSquare)
	if squareResult != squareExpected {
		t.Errorf(errorMsg, squareResult, squareExpected)
	}

	circleExpected := sideLen * sideLen * math.Pi
	circleResult := CalcSquare(sideLen, SidesCircle)
	if circleResult != circleExpected {
		t.Errorf(errorMsg, circleResult, circleExpected)
	}
}
