package main

import (
	"fmt"
	"os"
	"strconv"
	"task-golang/utils" // Import the utils package
)

func main() {
	// Read n from command line arguments
	n := 1 // default value
	if len(os.Args) > 1 {
		var err error
		n, err = strconv.Atoi(os.Args[1]) // convert string to int
		if err != nil {
			fmt.Println("Invalid number. Using default n=1")
			n = 1
		}
	}

	fmt.Println(utils.MagicSum(16))
	fmt.Println(utils.MagicPow(3))
	fmt.Println(utils.MagicOdd(4))
	fmt.Println(utils.MagicGrade(4))
	fmt.Println(utils.MagicName(2))

	// MagicChange
	fmt.Println("Before MagicChange:", n)
	utils.MagicChange(&n)
	fmt.Println("After MagicChange:", n)

	// MagicTria
	fmt.Println("MagicTria:", utils.MagicTria(n))

	// Initialize MagicNumber struct and use Multiply method
	magicNumber := utils.MagicNumber{Number: n}
	fmt.Println("Before Multiply:", magicNumber.Number)
	magicNumber.Multiply(n)
	fmt.Println("After Multiply:", magicNumber.Number)
}
