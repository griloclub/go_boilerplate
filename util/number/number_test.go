package number

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFloatingPoint(t *testing.T) {

	y := 0.1
	y1 := 0.2
	totaly := y + y1

	require.Less(t, 0.3, totaly)                 //different value because of floating point
	require.Equal(t, 0.3, RoundFloat(totaly, 2)) //fixed floating point value

	x := 0.07
	x1 := 0.07
	x2 := 0.07

	totalx := x + x1 + x2
	require.Less(t, 0.21, totalx)                 //different value because of floating point
	require.Equal(t, 0.21000000000000002, totalx) //floating point problem
	require.Equal(t, 0.21, RoundFloat(totalx, 2)) //fixed floating point value

	tax := totalx * 0.05
	require.Less(t, 0.0105, tax)                         //different value because of floating point
	require.Equal(t, 0.010500000000000002, tax)          //floating point problem
	require.Equal(t, 0.0105, RoundFloat(totalx, 4)*0.05) //fixed floating point value

	final := totalx + tax
	require.Less(t, 0.2205, final)                 //different value because of floating point
	require.Equal(t, 0.22050000000000003, final)   //floating point problem
	require.Equal(t, 0.2205, RoundFloat(final, 4)) //fixed floating point value

}
