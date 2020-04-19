package dir

import (
	"github.com/Loptt/lambdish-compiler/types"
	"reflect"
	"testing"
)

func TestStringDir(t *testing.T) {
	tests := []struct {
		e    funcentry
		want string
	}{
		{
			e: funcentry{
				"getMax",
				types.NewLambdishType(types.Bool, 0),
				2,
				[]types.LambdishType{
					types.NewLambdishType(types.Bool, 0),
					types.NewLambdishType(types.Bool, 0),
				},
			},
			want: "getMax@3@33",
		},
		{
			e: funcentry{
				"exists",
				types.NewLambdishType(types.Bool, 0),
				1,
				[]types.LambdishType{
					types.NewLambdishType(types.Bool, 0),
				},
			},
			want: "exists@3@3",
		},
		{
			e: funcentry{
				"goGet",
				types.NewLambdishType(types.Num, 0),
				1,
				[]types.LambdishType{
					types.NewLambdishType(types.Bool, 0),
				},
			},
			want: "goGet@1@3",
		},
		{
			e: funcentry{
				"getEven",
				types.NewLambdishType(types.Num, 1),
				2,
				[]types.LambdishType{
					types.NewLambdishType(types.Num, 1),
					types.NewLambdishType(types.Num, 0),
				},
			},
			want: "getEven@[1]@[1]1",
		},
	}

	for _, test := range tests {
		got := test.e.String()

		if got != test.want {
			t.Errorf("String() = %v, want: %v", got, test.want)
		}
	}
}

func TestAddDir(t *testing.T) {
	tests := []struct {
		e    funcentry
		fd   FuncDirectory
		want FuncDirectory
	}{
		{
			fd: FuncDirectory{
				make(map[string]funcentry),
			},
			e: funcentry{
				"getMax",
				types.NewLambdishType(types.Bool, 0),
				2,
				[]types.LambdishType{
					types.NewLambdishType(types.Bool, 0),
					types.NewLambdishType(types.Bool, 0),
				},
			},
			want: FuncDirectory{
				map[string]funcentry{
					"getMax@3@33": funcentry{
						"getMax",
						types.NewLambdishType(types.Bool, 0),
						2,
						[]types.LambdishType{
							types.NewLambdishType(types.Bool, 0),
							types.NewLambdishType(types.Bool, 0),
						},
					},
				},
			},
		},
		{
			fd: FuncDirectory{
				map[string]funcentry{
					"getMax@3@33": funcentry{
						"getMax",
						types.NewLambdishType(types.Bool, 0),
						2,
						[]types.LambdishType{
							types.NewLambdishType(types.Bool, 0),
							types.NewLambdishType(types.Bool, 0),
						},
					},
				},
			},
			e: funcentry{
				"exists",
				types.NewLambdishType(types.Bool, 0),
				1,
				[]types.LambdishType{
					types.NewLambdishType(types.Bool, 0),
				},
			},
			want: FuncDirectory{
				map[string]funcentry{
					"getMax@3@33": funcentry{
						"getMax",
						types.NewLambdishType(types.Bool, 0),
						2,
						[]types.LambdishType{
							types.NewLambdishType(types.Bool, 0),
							types.NewLambdishType(types.Bool, 0),
						},
					},
					"exists@3@3": funcentry{
						"exists",
						types.NewLambdishType(types.Bool, 0),
						1,
						[]types.LambdishType{
							types.NewLambdishType(types.Bool, 0),
						},
					},
				},
			},
		},
		{
			fd: FuncDirectory{
				map[string]funcentry{
					"getMax@3@33": funcentry{
						"getMax",
						types.NewLambdishType(types.Bool, 0),
						2,
						[]types.LambdishType{
							types.NewLambdishType(types.Bool, 0),
							types.NewLambdishType(types.Bool, 0),
						},
					},
					"exists@3@3": funcentry{
						"exists",
						types.NewLambdishType(types.Bool, 0),
						1,
						[]types.LambdishType{
							types.NewLambdishType(types.Bool, 0),
						},
					},
				},
			},
			e: funcentry{
				"getEven",
				types.NewLambdishType(types.Num, 1),
				2,
				[]types.LambdishType{
					types.NewLambdishType(types.Num, 1),
					types.NewLambdishType(types.Num, 0),
				},
			},
			want: FuncDirectory{
				map[string]funcentry{
					"getMax@3@33": funcentry{
						"getMax",
						types.NewLambdishType(types.Bool, 0),
						2,
						[]types.LambdishType{
							types.NewLambdishType(types.Bool, 0),
							types.NewLambdishType(types.Bool, 0),
						},
					},
					"exists@3@3": funcentry{
						"exists",
						types.NewLambdishType(types.Bool, 0),
						1,
						[]types.LambdishType{
							types.NewLambdishType(types.Bool, 0),
						},
					},
					"getEven@[1]@[1]1": funcentry{
						"getEven",
						types.NewLambdishType(types.Num, 1),
						2,
						[]types.LambdishType{
							types.NewLambdishType(types.Num, 1),
							types.NewLambdishType(types.Num, 0),
						},
					},
				},
			},
		},
	}
	for _, test := range tests {
		test.fd.Add(test.e)
		got := test.fd
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("fd.Add(%v) = %v, want: %v", test.e, got, test.want)
		}
	}
}

func TestGetDir(t *testing.T) {
	tests := []struct {
		key  string
		fd   FuncDirectory
		want *funcentry
	}{
		{
			key: "getEven@[1]@[1]1",
			fd: FuncDirectory{
				map[string]funcentry{
					"getMax@3@33": funcentry{
						"getMax",
						types.NewLambdishType(types.Bool, 0),
						2,
						[]types.LambdishType{
							types.NewLambdishType(types.Bool, 0),
							types.NewLambdishType(types.Bool, 0),
						},
					},
					"exists@3@3": funcentry{
						"exists",
						types.NewLambdishType(types.Bool, 0),
						1,
						[]types.LambdishType{
							types.NewLambdishType(types.Bool, 0),
						},
					},
					"getEven@[1]@[1]1": funcentry{
						"getEven",
						types.NewLambdishType(types.Num, 1),
						2,
						[]types.LambdishType{
							types.NewLambdishType(types.Num, 1),
							types.NewLambdishType(types.Num, 0),
						},
					},
				},
			},
			want: &funcentry{
				"getEven",
				types.NewLambdishType(types.Num, 1),
				2,
				[]types.LambdishType{
					types.NewLambdishType(types.Num, 1),
					types.NewLambdishType(types.Num, 0),
				},
			},
		},
		{
			key: "exists@3@3",
			fd: FuncDirectory{
				map[string]funcentry{
					"getMax@3@33": funcentry{
						"getMax",
						types.NewLambdishType(types.Bool, 0),
						2,
						[]types.LambdishType{
							types.NewLambdishType(types.Bool, 0),
							types.NewLambdishType(types.Bool, 0),
						},
					},
					"exists@3@3": funcentry{
						"exists",
						types.NewLambdishType(types.Bool, 0),
						1,
						[]types.LambdishType{
							types.NewLambdishType(types.Bool, 0),
						},
					},
					"getEven@[1]@[1]1": funcentry{
						"getEven",
						types.NewLambdishType(types.Num, 1),
						2,
						[]types.LambdishType{
							types.NewLambdishType(types.Num, 1),
							types.NewLambdishType(types.Num, 0),
						},
					},
				},
			},
			want: &funcentry{
						"exists",
						types.NewLambdishType(types.Bool, 0),
						1,
						[]types.LambdishType{
							types.NewLambdishType(types.Bool, 0),
						},
			},
		},
		{
			key: "exists@3@4",
			fd: FuncDirectory{
				map[string]funcentry{
					"getMax@3@33": funcentry{
						"getMax",
						types.NewLambdishType(types.Bool, 0),
						2,
						[]types.LambdishType{
							types.NewLambdishType(types.Bool, 0),
							types.NewLambdishType(types.Bool, 0),
						},
					},
					"exists@3@3": funcentry{
						"exists",
						types.NewLambdishType(types.Bool, 0),
						1,
						[]types.LambdishType{
							types.NewLambdishType(types.Bool, 0),
						},
					},
					"getEven@[1]@[1]1": funcentry{
						"getEven",
						types.NewLambdishType(types.Num, 1),
						2,
						[]types.LambdishType{
							types.NewLambdishType(types.Num, 1),
							types.NewLambdishType(types.Num, 0),
						},
					},
				},
			},
			want: nil,
		},
	}
	for _, test := range tests {
		got := test.fd.Get(test.key)
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("fd.Get(%+v) = %+v, want: %+v", test.key, got, test.want)
		}
	}
}

func TestExistsDir(t *testing.T) {
	tests := []struct {
		key  string
		fd   FuncDirectory
		want bool
	}{
		{
			key: "getEven@[1]@[1]1",
			fd: FuncDirectory{
				map[string]funcentry{
					"getMax@3@33": funcentry{
						"getMax",
						types.NewLambdishType(types.Bool, 0),
						2,
						[]types.LambdishType{
							types.NewLambdishType(types.Bool, 0),
							types.NewLambdishType(types.Bool, 0),
						},
					},
					"exists@3@3": funcentry{
						"exists",
						types.NewLambdishType(types.Bool, 0),
						1,
						[]types.LambdishType{
							types.NewLambdishType(types.Bool, 0),
						},
					},
					"getEven@[1]@[1]1": funcentry{
						"getEven",
						types.NewLambdishType(types.Num, 1),
						2,
						[]types.LambdishType{
							types.NewLambdishType(types.Num, 1),
							types.NewLambdishType(types.Num, 0),
						},
					},
				},
			},
			want: true,
		},
		{
			key: "exists@3@3",
			fd: FuncDirectory{
				map[string]funcentry{
					"getMax@3@33": funcentry{
						"getMax",
						types.NewLambdishType(types.Bool, 0),
						2,
						[]types.LambdishType{
							types.NewLambdishType(types.Bool, 0),
							types.NewLambdishType(types.Bool, 0),
						},
					},
					"exists@3@3": funcentry{
						"exists",
						types.NewLambdishType(types.Bool, 0),
						1,
						[]types.LambdishType{
							types.NewLambdishType(types.Bool, 0),
						},
					},
					"getEven@[1]@[1]1": funcentry{
						"getEven",
						types.NewLambdishType(types.Num, 1),
						2,
						[]types.LambdishType{
							types.NewLambdishType(types.Num, 1),
							types.NewLambdishType(types.Num, 0),
						},
					},
				},
			},
			want: true,
		},
		{
			key: "exists@3@4",
			fd: FuncDirectory{
				map[string]funcentry{
					"getMax@3@33": funcentry{
						"getMax",
						types.NewLambdishType(types.Bool, 0),
						2,
						[]types.LambdishType{
							types.NewLambdishType(types.Bool, 0),
							types.NewLambdishType(types.Bool, 0),
						},
					},
					"exists@3@3": funcentry{
						"exists",
						types.NewLambdishType(types.Bool, 0),
						1,
						[]types.LambdishType{
							types.NewLambdishType(types.Bool, 0),
						},
					},
					"getEven@[1]@[1]1": funcentry{
						"getEven",
						types.NewLambdishType(types.Num, 1),
						2,
						[]types.LambdishType{
							types.NewLambdishType(types.Num, 1),
							types.NewLambdishType(types.Num, 0),
						},
					},
				},
			},
			want: false,
		},
	}
	for _, test := range tests {
		got := test.fd.Exists(test.key)
		if got != test.want {
			t.Errorf("fd.Exists(%+v) = %+v, want: %+v", test.key, got, test.want)
		}
	}
}

