package main

import (
	"fmt"
	"log"
)

func main() {
	var number int

	if _, err := fmt.Scanln(&number); err != nil {
		log.Fatal(err)
	}

	b := false

	for {
		if number%2 == 0 {
			if number/2 == 1 {
				fmt.Print("iya")
				b = true
				break
			}
		}

		number = number / 2
	}

	if !b {
		fmt.Print("bukan")
	}

}
