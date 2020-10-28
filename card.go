package gocard

import (
	"fmt"
	"strconv"
)

type Card struct {
	number        []int
	cvv           string
	month         int
	year          int
	paymentSystem PaymentSystem
	cardHolder    string
}

// Company holds a short and long names of who has issued the credit card
type PaymentSystem struct {
	Short, Full string
}

func NewCard(number string, cvv string, month int, year int) (*Card, error) {
	if (len(number) < 13) || (len(number) > 19) {
		return nil, fmt.Errorf("Invalid number length")
	}
	if (month > 12) || (month < 0) || (year < 0) {
		return nil, fmt.Errorf("Invalid expiration date")
	}
	digits := []int{}
	for i := 0; i < len(number); i++ {
		d, err := strconv.Atoi(string(number[i]))
		if err != nil {
			return nil, fmt.Errorf("Invalid characters in number")
		}
		digits = append(digits, d)
	}

	if !Valid(digits) {
		return nil, fmt.Errorf("Card number is not valid")
	}

	c := &Card{number: digits, cvv: cvv, month: month, year: year}
	return c, nil

}
