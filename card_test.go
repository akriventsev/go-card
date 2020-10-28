package gocard

import (
	"reflect"
	"testing"
)

func TestNewCard(t *testing.T) {
	type args struct {
		number string
		cvv    string
		month  int
		year   int
	}
	tests := []struct {
		name    string
		args    args
		want    *Card
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "VISA valid test",
			args:    args{number: "4716339239466898", cvv: "334", year: 2023, month: 12},
			want:    &Card{number: []int{4, 7, 1, 6, 3, 3, 9, 2, 3, 9, 4, 6, 6, 8, 9, 8}, cvv: "334", month: 12, year: 2023},
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewCard(tt.args.number, tt.args.cvv, tt.args.month, tt.args.year)
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
