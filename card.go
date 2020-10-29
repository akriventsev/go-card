package card

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Card struct {
	number        []int
	cvv           string
	month         int
	year          int
	paymentSystem PaymentSystem
	cardHolder    string
}

func (c *Card) Number() string {
	return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(c.number)), ""), "[]")
}

func (c *Card) MaskedNumber(headlen int, taillen int) string {

	hl := headlen
	tl := taillen
	if (hl < 0) || (hl > len(c.number)) {
		hl = 6
	}
	if (tl < 0) || (tl > len(c.number)) {
		tl = 4
	}
	head := c.number[:hl]
	tail := c.number[len(c.number)-tl:]

	maskednumber := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(head)), ""), "[]")
	masklen := len(c.number) - len(head) - len(tail)
	maskednumber = maskednumber + strings.Repeat("*", masklen)
	maskednumber = maskednumber + strings.Trim(strings.Join(strings.Fields(fmt.Sprint(tail)), ""), "[]")
	return maskednumber
}

func (c *Card) Cvv() string {
	return c.cvv
}

func (c *Card) Month() int {
	return c.month
}

func (c *Card) Year() int {
	return c.year
}

func (c *Card) CardHolder() string {
	return c.cardHolder
}

func (c *Card) Expired() bool {

	now := time.Now()
	expirationDate := time.Date(c.year, time.Month(c.month), 1, 0, 0, 0, 0, now.Location())
	expirationDate.AddDate(0, 1, 0)
	return expirationDate.Before(now)

}

// Company holds a short and long names of who has issued the credit card
type PaymentSystem struct {
	Short, Full string
}

func NewCard(number string, cvv string, month int, year int, cardHolder string) (*Card, error) {
	length := len(number)

	if (length < 14) || (length > 19) {
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

	if !valid(digits) {
		return nil, fmt.Errorf("Card number is not valid")
	}

	c := &Card{number: digits, cvv: cvv, month: month, year: year, cardHolder: cardHolder}
	return c, nil

}
