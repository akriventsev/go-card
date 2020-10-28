package card

import (
	"reflect"
	"testing"
)

var testCard = &Card{number: []int{4, 7, 1, 6, 3, 3, 9, 2, 3, 9, 4, 6, 6, 8, 9, 8}, cvv: "334", month: 12, year: 2023, cardHolder: "Ivanov Ivan"}
var expiredCard = &Card{number: []int{4, 7, 1, 6, 3, 3, 9, 2, 3, 9, 4, 6, 6, 8, 9, 8}, cvv: "334", month: 12, year: 2019, cardHolder: "Ivanov Ivan"}

func TestNewCard(t *testing.T) {
	type args struct {
		number     string
		cvv        string
		month      int
		year       int
		cardHolder string
	}
	tests := []struct {
		name    string
		args    args
		want    *Card
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "VISA valid test",
			args:    args{number: "4716339239466898", cvv: "334", year: 2023, month: 12, cardHolder: "Ivanov Ivan"},
			want:    testCard,
			wantErr: false,
		},
		{name: "VISA invalid number charactertest",
			args:    args{number: "471633923z466898", cvv: "334", year: 2023, month: 12},
			want:    nil,
			wantErr: true,
		},
		{name: "VISA invalid number test",
			args:    args{number: "4716339239466897", cvv: "334", year: 2023, month: 12},
			want:    nil,
			wantErr: true,
		},
		{name: "VISA invalid year test",
			args:    args{number: "4716339239466898", cvv: "334", year: -2023, month: 12},
			want:    nil,
			wantErr: true,
		},
		{name: "VISA invalid month test sub zero",
			args:    args{number: "4716339239466898", cvv: "334", year: 2023, month: -12},
			want:    nil,
			wantErr: true,
		},
		{name: "VISA invalid month test over 12",
			args:    args{number: "4716339239466898", cvv: "334", year: 2023, month: 13},
			want:    nil,
			wantErr: true,
		},
		{name: "VISA invalid number length over 19",
			args:    args{number: "471633923946689811111", cvv: "334", year: 2023, month: 12},
			want:    nil,
			wantErr: true,
		},
		{name: "VISA invalid number length lower than 14",
			args:    args{number: "4716339239466", cvv: "334", year: 2023, month: 12},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewCard(tt.args.number, tt.args.cvv, tt.args.month, tt.args.year, tt.args.cardHolder)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewCard() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCard() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCard_Number(t *testing.T) {
	tests := []struct {
		name string
		c    *Card
		want string
	}{
		// TODO: Add test cases.
		{name: "Test number getter",
			c:    testCard,
			want: "4716339239466898",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Number(); got != tt.want {
				t.Errorf("Card.Number() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCard_Cvv(t *testing.T) {
	tests := []struct {
		name string
		c    *Card
		want string
	}{
		{
			name: "Test cvv getter",
			c:    testCard,
			want: "334",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Cvv(); got != tt.want {
				t.Errorf("Card.Cvv() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCard_Month(t *testing.T) {
	tests := []struct {
		name string
		c    *Card
		want int
	}{
		{
			name: "Test cvv getter",
			c:    testCard,
			want: 12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Month(); got != tt.want {
				t.Errorf("Card.Month() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCard_Year(t *testing.T) {
	tests := []struct {
		name string
		c    *Card
		want int
	}{
		{
			name: "Test cvv getter",
			c:    testCard,
			want: 2023,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Year(); got != tt.want {
				t.Errorf("Card.Year() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCard_CardHolder(t *testing.T) {
	tests := []struct {
		name string
		c    *Card
		want string
	}{
		{
			name: "Test cvv getter",
			c:    testCard,
			want: "Ivanov Ivan",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.CardHolder(); got != tt.want {
				t.Errorf("Card.CardHolder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCard_MaskedNumber(t *testing.T) {
	type args struct {
		headlen int
		taillen int
	}
	tests := []struct {
		name string
		c    *Card
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "Test cvv getter",
			c:    testCard,
			args: args{headlen: 6, taillen: 4},
			want: "471633******6898",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.MaskedNumber(tt.args.headlen, tt.args.taillen); got != tt.want {
				t.Errorf("Card.MaskedNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCard_Expired(t *testing.T) {
	tests := []struct {
		name string
		c    *Card
		want bool
	}{
		{
			name: "Test not expired card",
			c:    testCard,
			want: false,
		},
		{
			name: "Test expired card",
			c:    expiredCard,
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Expired(); got != tt.want {
				t.Errorf("Card.Expired() = %v, want %v", got, tt.want)
			}
		})
	}
}
