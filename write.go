package bakar

import (
	"math"
	"os"
)

/*****************Read********************/
func Read(filename string) ([]byte, error) {

	sampleFile, err := os.ReadFile(filename)
	Check(err)

	return sampleFile, err

}

/*****************Check********************/
func Check(err error) {
	if err != nil {
		panic(err)
	}
}

/*****************Write********************/
func Write(filename string, content []string, sep string) error {
	var tabByte []byte
	for index, v := range content {
		if index != len(content)-1 {
			tabByte = append(tabByte, []byte(v+sep)...)
		} else {
			tabByte = append(tabByte, []byte(v)...)
		}
	}
	return os.WriteFile(filename, tabByte, 0777)
}

/*****************Average********************/
func Average(tabFloat []float64) float64 {
	sum := 0.0
	for _, v := range tabFloat {
		sum += v
	}
	return sum / float64(len(tabFloat))
}

/*****************Median********************/
func Median(tabFloat []float64) int {
	sortFloat64(tabFloat)
	length := len(tabFloat)
	middle := len(tabFloat) / 2
	median := 0.0
	if length%2 == 0 {
		middleLeft := tabFloat[middle-1]
		middleRight := tabFloat[middle]
		median = (middleLeft + middleRight) / 2
	} else if length%2 != 0 {
		median = tabFloat[middle]
	}
	return int(math.Round(median))
}

/*****************Variance********************/
func Variance(tabFloat []float64, mean float64) int {
	sum := 0.0
	for _, v := range tabFloat {
		sum += math.Pow(v-mean, 2)
	}
	return int(math.Round(sum / float64(len(tabFloat))))
}

/*****************Standard Deviation********************/
func StandardDeviation(tabFloat []float64, mean float64) int {
	return int(math.Round(math.Sqrt(float64(Variance(tabFloat, mean)))))
}

/*****************SortFloat64********************/
func sortFloat64(numbers []float64) {
	for i := 0; i < len(numbers); i++ {
		for j := i + 1; j < len(numbers); j++ {
			if numbers[j] < numbers[i] {
				temp := numbers[i]
				numbers[i] = numbers[j]
				numbers[j] = temp
			}
		}
	}
}
