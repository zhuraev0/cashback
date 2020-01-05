package main

import "testing"

func Test_cashback(t *testing.T) {

	tests := []struct {
		name string
		amount int
		want int
	}{
		// TODO: Add test cases.
		{name:"Have cashback", amount:5000, want:250},
		{name:"No cashback", amount:1000, want:0},
		{name:"Cashback on bound", amount:3000, want:150},
	}
	for _, test := range tests {
		got := cashback(test.amount)
		if got!= test.want{
			t.Error("For cashback test:", test.name,"got:", got, "want:", test.want)
		}
	}
}