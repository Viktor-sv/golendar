package db

import "testing"

func Test_addTwoNum(t *testing.T) {
	type args struct {
		x int
		y int
	}

	tests := []struct {
		name string
		args args
		want int
	}{

		{name: "mytest", args: struct {
			x int
			y int
		}{x: 2, y: 5},
			want: 7},

		{name: "mytest2", args: struct {
			x int
			y int
		}{x: 23, y: 9}, want: 32},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addTwoNum(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("addTwoNum() = %v, want %v", got, tt.want)
			}
		})
	}
}
