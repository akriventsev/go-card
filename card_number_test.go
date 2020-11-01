package card

import (
	"reflect"
	"testing"
)

func TestNewCardNumber(t *testing.T) {
	type args struct {
		number string
	}
	tests := []struct {
		name    string
		args    args
		want    *Number
		wantErr bool
	}{
		{name: "VISA valid test",
			args:    args{number: "4716339239466898"},
			want:    &Number{number: "4716339239466898"},
			wantErr: false,
		},
		{name: "VISA invalid number charactertest",
			args:    args{number: "471633923z466898"},
			want:    nil,
			wantErr: true,
		},
		{name: "VISA invalid number checksum",
			args:    args{number: "4716339239466899"},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewCardNumber(tt.args.number)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewCardNumber() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCardNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
