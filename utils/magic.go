package utils

// MagicSum returns the sum of n + n
func MagicSum(n int) int {
	return n + n
}

// MagicPow returns n raised to the power of n
func MagicPow(n int) int {
	result := 1
	for i := 0; i < n; i++ {
		result *= n
	}
	return result
}

// MagicOdd returns true if n is odd, otherwise false
func MagicOdd(n int) bool {
	return n%2 != 0
}

// MagicGrade returns a string based on the grade
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

// MagicName returns a slice filled with the name repeated n times
func MagicName(n int) []string {
	names := []string{}
	for i := 0; i < n; i++ {
		names = append(names, "safira")
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

// MagicChange changes the value of n to n * 2
func MagicChange(n *int) {
	*n *= 2
}

// MagicNumber struct with a method to multiply its Number
type MagicNumber struct {
	Number int
}

// Multiply multiplies Number with n
func (m *MagicNumber) Multiply(n int) {
	m.Number *= n
}
