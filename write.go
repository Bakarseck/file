package file

import (
	"os"
	"strings"
)



func Read(filename string, sep string) ([]string, error) {

	sampleFile, err := os.ReadFile(filename)
	Check(err)

	return strings.Split(string(sampleFile), sep), err

}

func Check(err error) {
	if err != nil {
		panic(err)
	}
}

func Write(filename string, content []string, sep string) error {
	var tabByte []byte
	for index, v := range content {
		if index != len(content) -1 {
			tabByte = append(tabByte, []byte(v + sep)...)
		} else {
			tabByte = append(tabByte, []byte(v)...)
		}
	}
	return os.WriteFile(filename, tabByte, 0777)
}