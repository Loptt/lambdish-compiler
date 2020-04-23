package dir

import (
	"github.com/Loptt/lambdish-compiler/types"
	"reflect"
	"testing"
)

func TestStringDir(t *testing.T) {
	tests := []struct {
		e    *FuncEntry
		want string
	}{
		{
			e: &FuncEntry{
				"getMax",
				types.NewLambdishType(types.Bool, 0),
				2,
				[]types.LambdishType{
					types.NewLambdishType(types.Bool, 0),
					types.NewLambdishType(types.Bool, 0),
				},
				&VarDirectory{
					map[string]*VarEntry{
						"isEnabled": &VarEntry{
							"isEnabled",
							types.NewLambdishType(types.Bool, 0),
						},
						"minVal": &VarEntry{
							"minVal",
							types.NewLambdishType(types.Num, 0),
						},
						"arrBool": &VarEntry{
							"arrBool",
							types.NewLambdishType(types.Bool, 10),
						},
					},
				},
			},
			want: "getMax@3@33",
		},
		{
			e: &FuncEntry{
				"exists",
				types.NewLambdishType(types.Bool, 0),
				1,
				[]types.LambdishType{
					types.NewLambdishType(types.Bool, 0),
				},
				&VarDirectory{
					map[string]*VarEntry{
						"isEnabled": &VarEntry{
							"isEnabled",
							types.NewLambdishType(types.Bool, 0),
						},
						"minVal": &VarEntry{
							"minVal",
							types.NewLambdishType(types.Num, 0),
						},
						"arrBool": &VarEntry{
							"arrBool",
							types.NewLambdishType(types.Bool, 10),
						},
					},
				},
			},
			want: "exists@3@3",
		},
		{
			e: &FuncEntry{
				"goGet",
				types.NewLambdishType(types.Num, 0),
				1,
				[]types.LambdishType{
					types.NewLambdishType(types.Bool, 0),
				},
				&VarDirectory{
					map[string]*VarEntry{
						"isEnabled": &VarEntry{
							"isEnabled",
							types.NewLambdishType(types.Bool, 0),
						},
						"minVal": &VarEntry{
							"minVal",
							types.NewLambdishType(types.Num, 0),
						},
						"arrBool": &VarEntry{
							"arrBool",
							types.NewLambdishType(types.Bool, 10),
						},
					},
				},
			},
			want: "goGet@1@3",
		},
		{
			e: &FuncEntry{
				"getEven",
				types.NewLambdishType(types.Num, 1),
				2,
				[]types.LambdishType{
					types.NewLambdishType(types.Num, 1),
					types.NewLambdishType(types.Num, 0),
				},
				&VarDirectory{
					map[string]*VarEntry{
						"isEnabled": &VarEntry{
							"isEnabled",
							types.NewLambdishType(types.Bool, 0),
						},
						"minVal": &VarEntry{
							"minVal",
							types.NewLambdishType(types.Num, 0),
						},
						"arrBool": &VarEntry{
							"arrBool",
							types.NewLambdishType(types.Bool, 10),
						},
					},
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
		e    *FuncEntry
		fd   FuncDirectory
		want FuncDirectory
	}{
		{
			fd: FuncDirectory{
				make(map[string]*FuncEntry),
			},
			e: &FuncEntry{
				"getMax",
				types.NewLambdishType(types.Bool, 0),
				2,
				[]types.LambdishType{
					types.NewLambdishType(types.Bool, 0),
					types.NewLambdishType(types.Bool, 0),
				},
				&VarDirectory{
					map[string]*VarEntry{
						"minVal": &VarEntry{
							"minVal",
							types.NewLambdishType(types.Num, 0),
						},
					},
				},
			},
			want: FuncDirectory{
				map[string]*FuncEntry{
					"getMax@3@33": &FuncEntry{
						"getMax",
						types.NewLambdishType(types.Bool, 0),
						2,
						[]types.LambdishType{
							types.NewLambdishType(types.Bool, 0),
							types.NewLambdishType(types.Bool, 0),
						},
						&VarDirectory{
							map[string]*VarEntry{
								"minVal": &VarEntry{
									"minVal",
									types.NewLambdishType(types.Num, 0),
								},
							},
						},
					},
				},
			},
		},
		{
			fd: FuncDirectory{
				map[string]*FuncEntry{
					"getMax@3@33": &FuncEntry{
						"getMax",
						types.NewLambdishType(types.Bool, 0),
						2,
						[]types.LambdishType{
							types.NewLambdishType(types.Bool, 0),
							types.NewLambdishType(types.Bool, 0),
						},
						&VarDirectory{
							map[string]*VarEntry{
								"minVal": &VarEntry{
									"minVal",
									types.NewLambdishType(types.Num, 0),
								},
							},
						},
					},
				},
			},
			e: &FuncEntry{
				"exists",
				types.NewLambdishType(types.Bool, 0),
				1,
				[]types.LambdishType{
					types.NewLambdishType(types.Bool, 0),
				},
				&VarDirectory{
					map[string]*VarEntry{
						"isEnabled": &VarEntry{
							"isEnabled",
							types.NewLambdishType(types.Bool, 0),
						},
					},
				},
			},
			want: FuncDirectory{
				map[string]*FuncEntry{
					"getMax@3@33": &FuncEntry{
						"getMax",
						types.NewLambdishType(types.Bool, 0),
						2,
						[]types.LambdishType{
							types.NewLambdishType(types.Bool, 0),
							types.NewLambdishType(types.Bool, 0),
						},
						&VarDirectory{
							map[string]*VarEntry{
								"minVal": &VarEntry{
									"minVal",
									types.NewLambdishType(types.Num, 0),
								},
							},
						},
					},
					"exists@3@3": &FuncEntry{
						"exists",
						types.NewLambdishType(types.Bool, 0),
						1,
						[]types.LambdishType{
							types.NewLambdishType(types.Bool, 0),
						},
						&VarDirectory{
							map[string]*VarEntry{
								"isEnabled": &VarEntry{
									"isEnabled",
									types.NewLambdishType(types.Bool, 0),
								},
							},
						},
					},
				},
			},
		},
		{
			fd: FuncDirectory{
				map[string]*FuncEntry{
					"getMax@3@33": &FuncEntry{
						"getMax",
						types.NewLambdishType(types.Bool, 0),
						2,
						[]types.LambdishType{
							types.NewLambdishType(types.Bool, 0),
							types.NewLambdishType(types.Bool, 0),
						},
						&VarDirectory{
							map[string]*VarEntry{
								"isEnabled": &VarEntry{
									"isEnabled",
									types.NewLambdishType(types.Bool, 0),
								},
							},
						},
					},
					"exists@3@3": &FuncEntry{
						"exists",
						types.NewLambdishType(types.Bool, 0),
						1,
						[]types.LambdishType{
							types.NewLambdishType(types.Bool, 0),
						},
						&VarDirectory{
							map[string]*VarEntry{
								"minVal": &VarEntry{
									"minVal",
									types.NewLambdishType(types.Num, 0),
								},
							},
						},
					},
				},
			},
			e: &FuncEntry{
				"getEven",
				types.NewLambdishType(types.Num, 1),
				2,
				[]types.LambdishType{
					types.NewLambdishType(types.Num, 1),
					types.NewLambdishType(types.Num, 0),
				},
				&VarDirectory{
					map[string]*VarEntry{
						"arrBool": &VarEntry{
							"arrBool",
							types.NewLambdishType(types.Bool, 10),
						},
					},
				},
			},
			want: FuncDirectory{
				map[string]*FuncEntry{
					"getMax@3@33": &FuncEntry{
						"getMax",
						types.NewLambdishType(types.Bool, 0),
						2,
						[]types.LambdishType{
							types.NewLambdishType(types.Bool, 0),
							types.NewLambdishType(types.Bool, 0),
						},
						&VarDirectory{
							map[string]*VarEntry{
								"isEnabled": &VarEntry{
									"isEnabled",
									types.NewLambdishType(types.Bool, 0),
								},
							},
						},
					},
					"exists@3@3": &FuncEntry{
						"exists",
						types.NewLambdishType(types.Bool, 0),
						1,
						[]types.LambdishType{
							types.NewLambdishType(types.Bool, 0),
						},
						&VarDirectory{
							map[string]*VarEntry{
								"minVal": &VarEntry{
									"minVal",
									types.NewLambdishType(types.Num, 0),
								},
							},
						},
					},
					"getEven@[1]@[1]1": &FuncEntry{
						"getEven",
						types.NewLambdishType(types.Num, 1),
						2,
						[]types.LambdishType{
							types.NewLambdishType(types.Num, 1),
							types.NewLambdishType(types.Num, 0),
						},
						&VarDirectory{
							map[string]*VarEntry{
								"arrBool": &VarEntry{
									"arrBool",
									types.NewLambdishType(types.Bool, 10),
								},
							},
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
		want *FuncEntry
	}{
		{
			key: "getEven@[1]@[1]1",
			fd: FuncDirectory{
				map[string]*FuncEntry{
					"getMax@3@33": &FuncEntry{
						"getMax",
						types.NewLambdishType(types.Bool, 0),
						2,
						[]types.LambdishType{
							types.NewLambdishType(types.Bool, 0),
							types.NewLambdishType(types.Bool, 0),
						},
						&VarDirectory{
							map[string]*VarEntry{
								"isEnabled": &VarEntry{
									"isEnabled",
									types.NewLambdishType(types.Bool, 0),
								},
							},
						},
					},
					"exists@3@3": &FuncEntry{
						"exists",
						types.NewLambdishType(types.Bool, 0),
						1,
						[]types.LambdishType{
							types.NewLambdishType(types.Bool, 0),
						},
						&VarDirectory{
							map[string]*VarEntry{
								"minVal": &VarEntry{
									"minVal",
									types.NewLambdishType(types.Num, 0),
								},
							},
						},
					},
					"getEven@[1]@[1]1": &FuncEntry{
						"getEven",
						types.NewLambdishType(types.Num, 1),
						2,
						[]types.LambdishType{
							types.NewLambdishType(types.Num, 1),
							types.NewLambdishType(types.Num, 0),
						},
						&VarDirectory{
							map[string]*VarEntry{
								"arrBool": &VarEntry{
									"arrBool",
									types.NewLambdishType(types.Bool, 10),
								},
							},
						},
					},
				},
			},
			want: &FuncEntry{
				"getEven",
				types.NewLambdishType(types.Num, 1),
				2,
				[]types.LambdishType{
					types.NewLambdishType(types.Num, 1),
					types.NewLambdishType(types.Num, 0),
				},
				&VarDirectory{
					map[string]*VarEntry{
						"arrBool": &VarEntry{
							"arrBool",
							types.NewLambdishType(types.Bool, 10),
						},
					},
				},
			},
		},
		{
			key: "exists@3@3",
			fd: FuncDirectory{
				map[string]*FuncEntry{
					"getMax@3@33": &FuncEntry{
						"getMax",
						types.NewLambdishType(types.Bool, 0),
						2,
						[]types.LambdishType{
							types.NewLambdishType(types.Bool, 0),
							types.NewLambdishType(types.Bool, 0),
						},
						&VarDirectory{
							map[string]*VarEntry{
								"isEnabled": &VarEntry{
									"isEnabled",
									types.NewLambdishType(types.Bool, 0),
								},
							},
						},
					},
					"exists@3@3": &FuncEntry{
						"exists",
						types.NewLambdishType(types.Bool, 0),
						1,
						[]types.LambdishType{
							types.NewLambdishType(types.Bool, 0),
						},
						&VarDirectory{
							map[string]*VarEntry{
								"minVal": &VarEntry{
									"minVal",
									types.NewLambdishType(types.Num, 0),
								},
							},
						},
					},
					"getEven@[1]@[1]1": &FuncEntry{
						"getEven",
						types.NewLambdishType(types.Num, 1),
						2,
						[]types.LambdishType{
							types.NewLambdishType(types.Num, 1),
							types.NewLambdishType(types.Num, 0),
						},
						&VarDirectory{
							map[string]*VarEntry{
								"arrBool": &VarEntry{
									"arrBool",
									types.NewLambdishType(types.Bool, 10),
								},
							},
						},
					},
				},
			},
			want: &FuncEntry{
				"exists",
				types.NewLambdishType(types.Bool, 0),
				1,
				[]types.LambdishType{
					types.NewLambdishType(types.Bool, 0),
				},
				&VarDirectory{
					map[string]*VarEntry{
						"minVal": &VarEntry{
							"minVal",
							types.NewLambdishType(types.Num, 0),
						},
					},
				},
			},
		},
		{
			key: "exists@3@4",
			fd: FuncDirectory{
				map[string]*FuncEntry{
					"getMax@3@33": &FuncEntry{
						"getMax",
						types.NewLambdishType(types.Bool, 0),
						2,
						[]types.LambdishType{
							types.NewLambdishType(types.Bool, 0),
							types.NewLambdishType(types.Bool, 0),
						},
						&VarDirectory{
							map[string]*VarEntry{
								"isEnabled": &VarEntry{
									"isEnabled",
									types.NewLambdishType(types.Bool, 0),
								},
							},
						},
					},
					"exists@3@3": &FuncEntry{
						"exists",
						types.NewLambdishType(types.Bool, 0),
						1,
						[]types.LambdishType{
							types.NewLambdishType(types.Bool, 0),
						},
						&VarDirectory{
							map[string]*VarEntry{
								"minVal": &VarEntry{
									"minVal",
									types.NewLambdishType(types.Num, 0),
								},
							},
						},
					},
					"getEven@[1]@[1]1": &FuncEntry{
						"getEven",
						types.NewLambdishType(types.Num, 1),
						2,
						[]types.LambdishType{
							types.NewLambdishType(types.Num, 1),
							types.NewLambdishType(types.Num, 0),
						},
						&VarDirectory{
							map[string]*VarEntry{
								"arrBool": &VarEntry{
									"arrBool",
									types.NewLambdishType(types.Bool, 10),
								},
							},
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
				map[string]*FuncEntry{
					"getMax@3@33": &FuncEntry{
						"getMax",
						types.NewLambdishType(types.Bool, 0),
						2,
						[]types.LambdishType{
							types.NewLambdishType(types.Bool, 0),
							types.NewLambdishType(types.Bool, 0),
						},
						&VarDirectory{
							map[string]*VarEntry{
								"isEnabled": &VarEntry{
									"isEnabled",
									types.NewLambdishType(types.Bool, 0),
								},
							},
						},
					},
					"exists@3@3": &FuncEntry{
						"exists",
						types.NewLambdishType(types.Bool, 0),
						1,
						[]types.LambdishType{
							types.NewLambdishType(types.Bool, 0),
						},
						&VarDirectory{
							map[string]*VarEntry{
								"minVal": &VarEntry{
									"minVal",
									types.NewLambdishType(types.Num, 0),
								},
							},
						},
					},
					"getEven@[1]@[1]1": &FuncEntry{
						"getEven",
						types.NewLambdishType(types.Num, 1),
						2,
						[]types.LambdishType{
							types.NewLambdishType(types.Num, 1),
							types.NewLambdishType(types.Num, 0),
						},
						&VarDirectory{
							map[string]*VarEntry{
								"arrBool": &VarEntry{
									"arrBool",
									types.NewLambdishType(types.Bool, 10),
								},
							},
						},
					},
				},
			},
			want: true,
		},
		{
			key: "exists@3@3",
			fd: FuncDirectory{
				map[string]*FuncEntry{
					"getMax@3@33": &FuncEntry{
						"getMax",
						types.NewLambdishType(types.Bool, 0),
						2,
						[]types.LambdishType{
							types.NewLambdishType(types.Bool, 0),
							types.NewLambdishType(types.Bool, 0),
						},
						&VarDirectory{
							map[string]*VarEntry{
								"isEnabled": &VarEntry{
									"isEnabled",
									types.NewLambdishType(types.Bool, 0),
								},
							},
						},
					},
					"exists@3@3": &FuncEntry{
						"exists",
						types.NewLambdishType(types.Bool, 0),
						1,
						[]types.LambdishType{
							types.NewLambdishType(types.Bool, 0),
						},
						&VarDirectory{
							map[string]*VarEntry{
								"minVal": &VarEntry{
									"minVal",
									types.NewLambdishType(types.Num, 0),
								},
							},
						},
					},
					"getEven@[1]@[1]1": &FuncEntry{
						"getEven",
						types.NewLambdishType(types.Num, 1),
						2,
						[]types.LambdishType{
							types.NewLambdishType(types.Num, 1),
							types.NewLambdishType(types.Num, 0),
						},
						&VarDirectory{
							map[string]*VarEntry{
								"arrBool": &VarEntry{
									"arrBool",
									types.NewLambdishType(types.Bool, 10),
								},
							},
						},
					},
				},
			},
			want: true,
		},
		{
			key: "exists@3@4",
			fd: FuncDirectory{
				map[string]*FuncEntry{
					"getMax@3@33": &FuncEntry{
						"getMax",
						types.NewLambdishType(types.Bool, 0),
						2,
						[]types.LambdishType{
							types.NewLambdishType(types.Bool, 0),
							types.NewLambdishType(types.Bool, 0),
						},
						&VarDirectory{
							map[string]*VarEntry{
								"isEnabled": &VarEntry{
									"isEnabled",
									types.NewLambdishType(types.Bool, 0),
								},
							},
						},
					},
					"exists@3@3": &FuncEntry{
						"exists",
						types.NewLambdishType(types.Bool, 0),
						1,
						[]types.LambdishType{
							types.NewLambdishType(types.Bool, 0),
						},
						&VarDirectory{
							map[string]*VarEntry{
								"minVal": &VarEntry{
									"minVal",
									types.NewLambdishType(types.Num, 0),
								},
							},
						},
					},
					"getEven@[1]@[1]1": &FuncEntry{
						"getEven",
						types.NewLambdishType(types.Num, 1),
						2,
						[]types.LambdishType{
							types.NewLambdishType(types.Num, 1),
							types.NewLambdishType(types.Num, 0),
						},
						&VarDirectory{
							map[string]*VarEntry{
								"arrBool": &VarEntry{
									"arrBool",
									types.NewLambdishType(types.Bool, 10),
								},
							},
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
