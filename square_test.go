package square

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCalcSqaureCustom(t *testing.T) {
	cases := []struct {
		name        string
		sidesNum    sidesAmount
		sideLen     float64
		expectedRes float64
	}{
		{
			name:        "triangle",
			sidesNum:    SidesTriangle,
			sideLen:     1.0,
			expectedRes: 0.4330127018922193,
		},
		{
			name:        "triangle 2",
			sidesNum:    SidesTriangle,
			sideLen:     2.0,
			expectedRes: 1.7320508075688772,
		},
		{
			name:        "triangle 3",
			sidesNum:    SidesTriangle,
			sideLen:     3.0,
			expectedRes: 3.8971143170299736,
		},
		{
			name:        "circle 1",
			sidesNum:    SidesCircle,
			sideLen:     1.0,
			expectedRes: 3.141592653589793,
		},
		{
			name:        "circle 4",
			sidesNum:    SidesCircle,
			sideLen:     4.0,
			expectedRes: 50.26548245743669,
		},
		{
			name:        "circle 2.3",
			sidesNum:    SidesCircle,
			sideLen:     2.3,
			expectedRes: 16.619025137490002,
		},
		{
			name:        "square 2.3",
			sidesNum:    SidesSquare,
			sideLen:     6.785,
			expectedRes: 6.785 * 6.785,
		},
	}

	for i, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			require.Equal(
				t,
				c.expectedRes,
				CalcSquare(c.sideLen, c.sidesNum),
				i+1,
			)
		})
	}

}
