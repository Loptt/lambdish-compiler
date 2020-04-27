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
		{
			b: Null,
			want: 'n',
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
				basic:    Num,
				list: 0,
				function: false,
			},
			want: "1",
		},
		{
			l: LambdishType{
				basic:    Bool,
				list: 1,
				function: false,
			},
			want: "[3]",
		},
		{
			l: LambdishType{
				basic:    Char,
				list: 2,
				function: false,
			},
			want: "[[2]]",
		},
		{
			l: LambdishType{
				basic:    Bool,
				list: 5,
				function: false,
			},
			want: "[[[[[3]]]]]",
		},
		{
			l: LambdishType{
				retval: &LambdishType{
					basic: Num,
					list: 0,
					function: false,
				},
				list: 0,
				params: []*LambdishType{
					{basic: Num, list: 0, function: false},
				},
				function: true,
			},
			want: "(1=>1)",
		},
		{
			l: LambdishType{
				retval: &LambdishType{
					basic: Bool,
					list: 0,
					function: false,
				},
				list: 0,
				params: []*LambdishType{
					{basic: Num, list: 1, function: false},
					{basic: Num, list: 1, function: false},
				},
				function: true,
			},
			want: "([1][1]=>3)",
		},
		{
			l: LambdishType{
				retval: &LambdishType{
					basic: Bool,
					list: 0,
					function: false,
				},
				list: 0,
				params: []*LambdishType{
					{
						retval: &LambdishType{
							basic: Num,
							list: 0,
							function: false,
						},
						list: 0, 
						function: true, 
						params: []*LambdishType{
							{basic: Num, list: 0, function: false},
							{basic: Num, list: 0, function: false},
						},
					},
					{basic: Num, list: 1, function: false},
				},
				function: true,
			},
			want: "((11=>1)[1]=>3)",
		},
		{
			l: LambdishType{
				retval: &LambdishType{
					retval: &LambdishType{
						basic: Char, 
						list: 0, 
						function: false,
					},
					list: 0,
					function: true,
					params: []*LambdishType{
						{basic: Num, list: 0, function: false},
					},
				},
				list: 0,
				params: []*LambdishType{
					{
						retval: &LambdishType{
							basic: Num, 
							list: 0, 
							function: false,
						}, 
						list: 0, 
						function: true, 
						params: []*LambdishType{
							{basic: Num, list: 0, function: false},
							{basic: Num, list: 0, function: false},
						},
					},
					{basic: Num, list: 1, function: false},
				},
				function: true,
			},
			want: "((11=>1)[1]=>(1=>2))",
		},
	}

	for i, test := range tests {
		got := test.l.String()

		if got != test.want {
			t.Errorf("Number %d = %v, want: %v", i, got, test.want)
		}
	}
}
