package gocard

// CalculateLuhn return the check number
func CalculateLuhn(digits []int) int {
	checkNumber := checksum(digits)

	if checkNumber == 0 {
		return 0
	}
	return 10 - checkNumber
}

// Valid check number is valid or not based on Luhn algorithm
func Valid(digits []int) bool {
	return (digits[len(digits)-1]+checksum(digits))%10 == 0
}

func checksum(digits []int) int {
	var luhn int

	for i := len(digits) - 2; i > -1; i-- {
		cur := digits[i]
		if i%2 == 0 { // even
			cur = cur * 2
			if cur > 9 {
				cur = cur%10 + 1
			}
		}
		luhn += cur
	}

	return luhn % 10
}
