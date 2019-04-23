package main
// Test case program of the temperature.go
import(
	"testing"
	"./temperature"
)

type temperatureTest struct{
	i float64
	expected float64
}

var CtoFTests = [] temperatureTest {
	{4.1, 39.28},
	{10, 50},
	{-10, 14},
}

var FtoCTest = [] temperatureTest{
	{32, 0},
	{50, 10},
	{5, 15},
}

func TestCtoF(t *testing.T){
	for _, tt := range CtoFTests{
		actual := temperature.CtoF(tt.i)
		if actual != tt.expected{
			t.Errorf("expectd %v, actural %v,", tt.expected, actual)
		}
	}
}

func testFtoC(t *testing.T){
	for _, ff := range FtoCTest{
		actual := temperature.FtoC(ff.i)
		if actual != ff.expected{
			t.Errorf("expect %v, actural %v", ff.expected, actual)
		}
	}
}
