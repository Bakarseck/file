package bakar

import (
	"strings"
	"testing"
)

func TestWrite(t *testing.T) {

	s := "hello how are you ?"

	tabString := strings.Split(s, " ")

	err := Write("a.txt", tabString, " ")

	if err != nil {
		t.Errorf("Error !!!")
	}
}

func TestRead(t *testing.T)  {
	_, err := Read("a.txt")

	if err != nil {
		t.Errorf("Error")
	}

}

func TestAverage(t *testing.T) {

	result := Average([]float64{1.0, 2.0, 3.0})

	if result != 2 {
		t.Errorf("Error !!!")
	}
}

func TestVariance(t *testing.T) {

	result := Variance([]float64{1.0, 2.0, 3.0}, Average([]float64{1.0, 2.0, 3.0}))

	if result != 1 {
		t.Errorf("Error !!!")
	}
}

func TestStandardDeviation(t *testing.T) {
	result := StandardDeviation([]float64{1.0, 2.0, 3.0}, Average([]float64{1.0, 2.0, 3.0}))

	if result != 1 {
		t.Errorf("Error !!!")
	}
}

func TestMedian(t *testing.T) {
	result := Median([]float64{1.0, 2.0, 3.0})

	if result != 2 {
		t.Errorf("Error !!!")
	}
}