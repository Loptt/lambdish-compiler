package types

import (
	"testing"
)

func TestConvert(t *testing.T) {
	tests := []struct {
		b    BasicType
		want rune
	}{
		{
			b:    Num,
			want: '1',
		},
		{
			b:    Char,
			want: '2',
		},
		{
			b:    Bool,
			want: '3',
		},
	}

	for _, test := range tests {
		got := test.b.convert()

		if got != test.want {
			t.Errorf("convert() = %v, want: %v", got, test.want)
		}
	}
}

func TestString(t *testing.T) {
	tests := []struct {
		l    LambdishType
		want string
	}{
		{
			l: LambdishType{
				t:    Num,
				list: 0,
			},
			want: "1",
		},
		{
			l: LambdishType{
				t:    Bool,
				list: 1,
			},
			want: "[3]",
		},
		{
			l: LambdishType{
				t:    Char,
				list: 2,
			},
			want: "[[2]]",
		},
		{
			l: LambdishType{
				t:    Bool,
				list: 5,
			},
			want: "[[[[[3]]]]]",
		},
	}

	for _, test := range tests {
		got := test.l.String()

		if got != test.want {
			t.Errorf("%v.String() = %v, want: %v", test.l, got, test.want)
		}
	}
}
