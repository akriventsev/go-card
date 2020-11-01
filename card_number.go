package card

import (
	"fmt"
	"strconv"
)

//Number - holds a credit card number
type Number struct {
	number string
}

func (c *Number) String() string {
	return c.number
}

//NewCardNumber - return a new Number object
func NewCardNumber(number string) (*Number, error) {
	digits := []int{}
	for i := 0; i < len(number); i++ {
		d, err := strconv.Atoi(string(number[i]))
		if err != nil {
			return nil, fmt.Errorf("Invalid characters in number")
		}
		digits = append(digits, d)
	}

	if !valid(digits) {
		return nil, fmt.Errorf("Card number is not valid")
	}
	return &Number{number: number}, nil
}
