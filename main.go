package main

import (
	"fmt"
	"task-golang/utils"
)

func main() {
	fmt.Println(utils.MagicSum(16))
	fmt.Println(utils.MagicPow(3))
	fmt.Println(utils.MagicOdd(4))
	fmt.Println(utils.MagicGrade(4))
	fmt.Println(utils.MagicName(2))
	fmt.Println(utils.MagicTria(3))

	n := 16
	utils.MagicChange(&n)
	fmt.Println("Changed value:", n)
}
