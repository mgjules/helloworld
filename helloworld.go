// Prints Hello World to Github in various forms.
package main

import (
	"fmt"

	"github.com/julesmike/stringreverse"
)

func main() {
	str := "Hello World to Github"
	fmt.Println(str)
	fmt.Println(stringreverse.Reverse(str))
}
