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
				function: false,
			},
			want: "1",
		},
		{
			l: LambdishType{
				t:    Bool,
				list: 1,
				function: false,
			},
			want: "[3]",
		},
		{
			l: LambdishType{
				t:    Char,
				list: 2,
				function: false,
			},
			want: "[[2]]",
		},
		{
			l: LambdishType{
				t:    Bool,
				list: 5,
				function: false,
			},
			want: "[[[[[3]]]]]",
		},
		{
			l: LambdishType{
				t:    Num,
				list: 0,
				params: []*LambdishType{
					{t: Num, list: 0, function: false},
				},
				function: true,
			},
			want: "(1=>1)",
		},
		{
			l: LambdishType{
				t:    Bool,
				list: 0,
				params: []*LambdishType{
					{t: Num, list: 1, function: false},
					{t: Num, list: 1, function: false},
				},
				function: true,
			},
			want: "([1][1]=>3)",
		},
		{
			l: LambdishType{
				t:    Bool,
				list: 0,
				params: []*LambdishType{
					{
						t: Num, 
						list: 0, 
						function: true, 
						params: []*LambdishType{
							{t: Num, list: 0, function: false},
							{t: Num, list: 0, function: false},
						},
					},
					{t: Num, list: 1, function: false},
				},
				function: true,
			},
			want: "((11=>1)[1]=>3)",
		},
	}

	for _, test := range tests {
		got := test.l.String()

		if got != test.want {
			t.Errorf("%v.String() = %v, want: %v", test.l, got, test.want)
		}
	}
}
