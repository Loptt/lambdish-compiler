package dir

import (
	"reflect"
	"testing"

	"github.com/Loptt/lambdish-compiler/types"
)

func TestStringVar(t *testing.T) {
	tests := []struct {
		e    *VarEntry
		want string
	}{
		{
			e: &VarEntry{
				"minVal",
				types.NewDataLambdishType(types.Num, 0),
			},
			want: "minVal",
		},
		{
			e: &VarEntry{
				"isEnabled",
				types.NewDataLambdishType(types.Bool, 0),
			},
			want: "isEnabled",
		},
		{
			e: &VarEntry{
				"maxValue",
				types.NewDataLambdishType(types.Num, 0),
			},
			want: "maxValue",
		},
		{
			e: &VarEntry{
				"arrBool",
				types.NewDataLambdishType(types.Bool, 10),
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
		e    *VarEntry
		vd   VarDirectory
		want VarDirectory
	}{
		{
			vd: VarDirectory{
				make(map[string]*VarEntry),
			},
			e: &VarEntry{
				"minVal",
				types.NewDataLambdishType(types.Num, 0),
			},
			want: VarDirectory{
				map[string]*VarEntry{
					"minVal": &VarEntry{
						"minVal",
						types.NewDataLambdishType(types.Num, 0),
					},
				},
			},
		},
		{
			vd: VarDirectory{
				map[string]*VarEntry{
					"isEnabled": &VarEntry{
						"isEnabled",
						types.NewDataLambdishType(types.Bool, 0),
					},
				},
			},
			e: &VarEntry{
				"minVal",
				types.NewDataLambdishType(types.Num, 0),
			},
			want: VarDirectory{
				map[string]*VarEntry{
					"isEnabled": &VarEntry{
						"isEnabled",
						types.NewDataLambdishType(types.Bool, 0),
					},
					"minVal": &VarEntry{
						"minVal",
						types.NewDataLambdishType(types.Num, 0),
					},
				},
			},
		},
		{
			vd: VarDirectory{
				map[string]*VarEntry{
					"isEnabled": &VarEntry{
						"isEnabled",
						types.NewDataLambdishType(types.Bool, 0),
					},
					"minVal": &VarEntry{
						"minVal",
						types.NewDataLambdishType(types.Num, 0),
					},
				},
			},
			e: &VarEntry{
				"arrBool",
				types.NewDataLambdishType(types.Bool, 10),
			},
			want: VarDirectory{
				map[string]*VarEntry{
					"isEnabled": &VarEntry{
						"isEnabled",
						types.NewDataLambdishType(types.Bool, 0),
					},
					"minVal": &VarEntry{
						"minVal",
						types.NewDataLambdishType(types.Num, 0),
					},
					"arrBool": &VarEntry{
						"arrBool",
						types.NewDataLambdishType(types.Bool, 10),
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
		want *VarEntry
	}{
		{
			key: "minVal",
			vd: VarDirectory{
				map[string]*VarEntry{
					"isEnabled": &VarEntry{
						"isEnabled",
						types.NewDataLambdishType(types.Bool, 0),
					},
					"minVal": &VarEntry{
						"minVal",
						types.NewDataLambdishType(types.Num, 0),
					},
					"arrBool": &VarEntry{
						"arrBool",
						types.NewDataLambdishType(types.Bool, 10),
					},
				},
			},
			want: &VarEntry{
				"minVal",
				types.NewDataLambdishType(types.Num, 0),
			},
		},
		{
			key: "x",
			vd: VarDirectory{
				map[string]*VarEntry{
					"isEnabled": &VarEntry{
						"isEnabled",
						types.NewDataLambdishType(types.Bool, 0),
					},
					"minVal": &VarEntry{
						"minVal",
						types.NewDataLambdishType(types.Num, 0),
					},
					"arrBool": &VarEntry{
						"arrBool",
						types.NewDataLambdishType(types.Bool, 10),
					},
				},
			},
			want: nil,
		},
		{
			key: "ARRBOOL",
			vd: VarDirectory{
				map[string]*VarEntry{
					"isEnabled": &VarEntry{
						"isEnabled",
						types.NewDataLambdishType(types.Bool, 0),
					},
					"minVal": &VarEntry{
						"minVal",
						types.NewDataLambdishType(types.Num, 0),
					},
					"arrBool": &VarEntry{
						"arrBool",
						types.NewDataLambdishType(types.Bool, 10),
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
				map[string]*VarEntry{
					"isEnabled": &VarEntry{
						"isEnabled",
						types.NewDataLambdishType(types.Bool, 0),
					},
					"minVal": &VarEntry{
						"minVal",
						types.NewDataLambdishType(types.Num, 0),
					},
					"arrBool": &VarEntry{
						"arrBool",
						types.NewDataLambdishType(types.Bool, 10),
					},
				},
			},
			want: true,
		},
		{
			key: "x",
			vd: VarDirectory{
				map[string]*VarEntry{
					"isEnabled": &VarEntry{
						"isEnabled",
						types.NewDataLambdishType(types.Bool, 0),
					},
					"minVal": &VarEntry{
						"minVal",
						types.NewDataLambdishType(types.Num, 0),
					},
					"arrBool": &VarEntry{
						"arrBool",
						types.NewDataLambdishType(types.Bool, 10),
					},
				},
			},
			want: false,
		},
		{
			key: "ARRBOOL",
			vd: VarDirectory{
				map[string]*VarEntry{
					"isEnabled": &VarEntry{
						"isEnabled",
						types.NewDataLambdishType(types.Bool, 0),
					},
					"minVal": &VarEntry{
						"minVal",
						types.NewDataLambdishType(types.Num, 0),
					},
					"arrBool": &VarEntry{
						"arrBool",
						types.NewDataLambdishType(types.Bool, 10),
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
