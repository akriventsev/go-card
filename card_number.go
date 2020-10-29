package card

import (
	"fmt"
	"strconv"
)

type CardNumber struct {
	number string
}

func (c *CardNumber) String() string {
	return c.number
}

func NewCardNumber(number string) (*CardNumber, error) {
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
	return &CardNumber{number: number}, nil
}
