package main

import (
	"io/ioutil"
	"log"
	"os"
)

func main() {

	bytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	floor := 0
	enteredBasement := false
	for idx, b := range bytes {
		if b == '(' {
			floor += 1
		} else if b == ')' {
			floor -= 1
		}
		if !enteredBasement && floor < 0 {
			enteredBasement = true
			log.Printf("position where santa enters basement: %d", idx)
		}
	}
	log.Printf("floor = %d", floor)

}
