package card

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Card - credit card structure
type Card struct {
	number        []int
	cvv           string
	month         int
	year          int
	paymentSystem PaymentSystem
	cardHolder    string
}

//Number - return a number of credit card
func (c *Card) Number() string {
	return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(c.number)), ""), "[]")
}

//MaskedNumber - return a masked credit card number like 427615xxxxxx1234
//@headlen - length of first series
//@taillen - length of last series
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

//Cvv - return a cvv of credit card
func (c *Card) Cvv() string {
	return c.cvv
}

//Month - return expiration month of credit card
func (c *Card) Month() int {
	return c.month
}

//Year - return expiration year of credit card
func (c *Card) Year() int {
	return c.year
}

//CardHolder - return cardholder of credit card
func (c *Card) CardHolder() string {
	return c.cardHolder
}

//Expired - return true if card is expired and false if not
func (c *Card) Expired() bool {

	now := time.Now()
	expirationDate := time.Date(c.year, time.Month(c.month), 1, 0, 0, 0, 0, now.Location())
	expirationDate.AddDate(0, 1, 0)
	return expirationDate.Before(now)

}

//PaymentSystem holds a short and long names of who has issued the credit card
type PaymentSystem struct {
	Short, Full string
}

//NewCard - create a new credit card object
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
