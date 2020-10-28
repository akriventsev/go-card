
# go-card 

Basic credit card structure with validation using the [Luhn algorithm](http://en.wikipedia.org/wiki/Luhn_algorithm)


## Installation

```bash
go get github.com/anaximen/go-card
```

## Quick Start

```go
// Initialize a new card:
card,err := card.NewCard("4716339239466898", "334", 2023, 12, "Ivanov Ivan")

if err!= nil {
    fmt.Print(err)
}
...

expired := card.Expired()

maskedNumber := card.MaskedNumber()