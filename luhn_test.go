package card

import "testing"

func TestValid(t *testing.T) {
	type args struct {
		digits []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{name: "VISA valid test", args: args{digits: []int{4, 7, 1, 6, 3, 3, 9, 2, 3, 9, 4, 6, 6, 8, 9, 8}}, want: true},
		{name: "VISA invalid test", args: args{digits: []int{4, 7, 1, 6, 3, 3, 9, 2, 3, 9, 4, 6, 6, 8, 9, 9}}, want: false},
		{name: "Master card valid test", args: args{digits: []int{2, 2, 2, 1, 1, 5, 5, 7, 2, 1, 2, 2, 3, 1, 8, 9}}, want: true},
		{name: "Master card invalid test", args: args{digits: []int{2, 2, 2, 1, 1, 5, 5, 7, 2, 1, 2, 2, 3, 1, 8, 7}}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := valid(tt.args.digits); got != tt.want {
				t.Errorf("Valid() = %v, want %v", got, tt.want)
			}
		})
	}
}
