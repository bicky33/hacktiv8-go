package helper

import (
	"fmt"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	test := HelloWorld("asep")

	if test != "Helloasep" {
		t.Error(test)
	}
	fmt.Println("heheh")
}
