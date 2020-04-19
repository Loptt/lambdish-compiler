package dir

import (
	"github.com/Loptt/lambdish-compiler/types"
	"testing"
	"reflect"
)

func TestStringVar(t *testing.T) {
	tests := []struct {
		e    varentry
		want string
	}{
		{
			e: varentry{
				"minVal",
				types.NewLambdishType(types.Num, 0),
				"1",
			},
			want: "minVal",
		},
		{
			e: varentry{
				"isEnabled",
				types.NewLambdishType(types.Bool, 0),
				"true",
			},
			want: "isEnabled",
		},
		{
			e: varentry{
				"maxValue",
				types.NewLambdishType(types.Num, 0),
				"55675437564564758547896",
			},
			want: "maxValue",
		},
		{
			e: varentry{
				"arrBool",
				types.NewLambdishType(types.Bool, 10),
				"[[[[[[[[[[true,true]]]]]]]]]]",
			},
			want: "arrBool",
		},
	}

	for _, test := range tests {
		got := test.e.String()

		if got != test.want {
			t.Errorf("String() = %v, want: %v", got, test.want)
		}
	}
}

func TestAddVar(t *testing.T) {
	tests := []struct {
		e    varentry
		vd   VarDirectory
		want VarDirectory
	}{
		{
			vd: VarDirectory{
				make(map[string]varentry),
			},
			e: varentry{
				"minVal",
				types.NewLambdishType(types.Num, 0),
				"1",
			},
			want: VarDirectory{
				map[string]varentry{
					"minVal": varentry{
					"minVal",
					types.NewLambdishType(types.Num, 0),
					"1",
					},
				},
			},
		},
		{
			vd: VarDirectory{
				map[string]varentry{
					"isEnabled": varentry{
						"isEnabled",
						types.NewLambdishType(types.Bool, 0),
						"true",
					},
				},
			},
			e: varentry{
				"minVal",
				types.NewLambdishType(types.Num, 0),
				"1",
			},
			want: VarDirectory{
				map[string]varentry{
					"isEnabled": varentry{
						"isEnabled",
						types.NewLambdishType(types.Bool, 0),
						"true",
					},
					"minVal": varentry{
						"minVal",
						types.NewLambdishType(types.Num, 0),
						"1",
					},
				},
			},
		},
		{
			vd: VarDirectory{
				map[string]varentry{
					"isEnabled": varentry{
						"isEnabled",
						types.NewLambdishType(types.Bool, 0),
						"true",
					},
					"minVal": varentry{
						"minVal",
						types.NewLambdishType(types.Num, 0),
						"1",
					},
				},
			},
			e: varentry{
				"arrBool",
				types.NewLambdishType(types.Bool, 10),
				"[[[[[[[[[[true,true]]]]]]]]]]",
			},
			want: VarDirectory{
				map[string]varentry{
					"isEnabled": varentry{
						"isEnabled",
						types.NewLambdishType(types.Bool, 0),
						"true",
					},
					"minVal": varentry{
						"minVal",
						types.NewLambdishType(types.Num, 0),
						"1",
					},
					"arrBool": varentry{
						"arrBool",
						types.NewLambdishType(types.Bool, 10),
						"[[[[[[[[[[true,true]]]]]]]]]]",
					},
				},
			},
		},
	}

	for _, test := range tests {
		test.vd.Add(test.e)
		got := test.vd
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("vd.Add(%v) = %v, want: %v", test.e, got, test.want)
		}
	}
}
func TestGetVar(t *testing.T) {
	tests := []struct {
		key  string
		vd   VarDirectory
		want *varentry
	}{
		{
			key: "minVal",
			vd: VarDirectory{
				map[string]varentry{
					"isEnabled": varentry{
						"isEnabled",
						types.NewLambdishType(types.Bool, 0),
						"true",
					},
					"minVal": varentry{
						"minVal",
						types.NewLambdishType(types.Num, 0),
						"1",
					},
					"arrBool": varentry{
						"arrBool",
						types.NewLambdishType(types.Bool, 10),
						"[[[[[[[[[[true,true]]]]]]]]]]",
					},
				},
			},
			want: &varentry{
						"minVal",
						types.NewLambdishType(types.Num, 0),
						"1",
			},
		},
		{
			key: "x",
			vd: VarDirectory{
				map[string]varentry{
					"isEnabled": varentry{
						"isEnabled",
						types.NewLambdishType(types.Bool, 0),
						"true",
					},
					"minVal": varentry{
						"minVal",
						types.NewLambdishType(types.Num, 0),
						"1",
					},
					"arrBool": varentry{
						"arrBool",
						types.NewLambdishType(types.Bool, 10),
						"[[[[[[[[[[true,true]]]]]]]]]]",
					},
				},
			},
			want: nil,
		},
		{
			key: "ARRBOOL",
			vd: VarDirectory{
				map[string]varentry{
					"isEnabled": varentry{
						"isEnabled",
						types.NewLambdishType(types.Bool, 0),
						"true",
					},
					"minVal": varentry{
						"minVal",
						types.NewLambdishType(types.Num, 0),
						"1",
					},
					"arrBool": varentry{
						"arrBool",
						types.NewLambdishType(types.Bool, 10),
						"[[[[[[[[[[true,true]]]]]]]]]]",
					},
				},
			},
			want: nil,
		},
	}

	for _, test := range tests {
		got := test.vd.Get(test.key)
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("vd.Get(%v) = %v, want: %v", test.key, got, test.want)
		}
	}
}
func TestExistsVar(t *testing.T) {
	tests := []struct {
		key  string
		vd   VarDirectory
		want bool
	}{
		{
			key: "minVal",
			vd: VarDirectory{
				map[string]varentry{
					"isEnabled": varentry{
						"isEnabled",
						types.NewLambdishType(types.Bool, 0),
						"true",
					},
					"minVal": varentry{
						"minVal",
						types.NewLambdishType(types.Num, 0),
						"1",
					},
					"arrBool": varentry{
						"arrBool",
						types.NewLambdishType(types.Bool, 10),
						"[[[[[[[[[[true,true]]]]]]]]]]",
					},
				},
			},
			want: true,
		},
		{
			key: "x",
			vd: VarDirectory{
				map[string]varentry{
					"isEnabled": varentry{
						"isEnabled",
						types.NewLambdishType(types.Bool, 0),
						"true",
					},
					"minVal": varentry{
						"minVal",
						types.NewLambdishType(types.Num, 0),
						"1",
					},
					"arrBool": varentry{
						"arrBool",
						types.NewLambdishType(types.Bool, 10),
						"[[[[[[[[[[true,true]]]]]]]]]]",
					},
				},
			},
			want: false,
		},
		{
			key: "ARRBOOL",
			vd: VarDirectory{
				map[string]varentry{
					"isEnabled": varentry{
						"isEnabled",
						types.NewLambdishType(types.Bool, 0),
						"true",
					},
					"minVal": varentry{
						"minVal",
						types.NewLambdishType(types.Num, 0),
						"1",
					},
					"arrBool": varentry{
						"arrBool",
						types.NewLambdishType(types.Bool, 10),
						"[[[[[[[[[[true,true]]]]]]]]]]",
					},
				},
			},
			want: false,
		},
	}

	for _, test := range tests {
		got := test.vd.Exists(test.key)
		if got != test.want {
			t.Errorf("vd.Exists(%v) = %v, want: %v", test.key, got, test.want)
		}
	}
}
