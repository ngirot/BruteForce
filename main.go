package main

import (
	"fmt"
)

func main() {
	fmt.Printf("Start brute-forcing...\n")
	var result = TestAllStrings(testCurrent, displayCurrent)

	if result != "" {
		fmt.Printf("Found : %s\n", result)
	} else {
		fmt.Printf("Not found")
	}
}

var i = 0
func displayCurrent(data string)  {
	i++
	//if(i%1000 == 0) {
		fmt.Printf("Done: %d [%s]\n", i, data)
	//}
}

func testCurrent(data string) bool {
	return Hash(data) == "2e7d2c03a9507ae265ecf5b5356885a53393a2029d241394997265a1a25aefc6"
}