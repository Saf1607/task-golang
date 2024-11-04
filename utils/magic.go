package utils

func MagicSum(n int) int {
	return n + n
}

// MagicPow takes an integer and returns n raised to the power of n
func MagicPow(n int) int {
	result := 1
	for i := 0; i < n; i++ {
		result *= n
	}
	return result
}

// MagicOdd takes an integer and returns true if it is an odd number
func MagicOdd(n int) bool {
	return n%2 != 0
}

// MagicGrade returns a string representing the grade based on the input number
func MagicGrade(n int) string {
	switch n {
	case 0:
		return "Zero"
	case 1:
		return "Bad"
	case 2:
		return "Ok"
	case 3:
		return "Nice"
	case 4:
		return "Awesome"
	case 5:
		return "Exceptional"
	default:
		return "Unknown"
	}
}

// MagicName returns a slice of strings containing the name repeated n times
func MagicName(n int) []string {
	names := []string{}
	for i := 0; i < n; i++ {
		names = append(names, "safira") // Replace "radit" with your name
	}
	return names
}

// MagicTria returns the sum of numbers from 1 to n
func MagicTria(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum
}

// MagicChange modifies the input integer by doubling its value
func MagicChange(n *int) {
	*n = *n * 2
}
