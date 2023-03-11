package file

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
	_, err := Read("a.txt", " ")

	if err != nil {
		t.Errorf("Error")
	}

}