package funcdir

import (
	"github.com/Loptt/lambdish-compiler/types"
	"testing"
)

func TestString(t *testing.T) {
	tests := []struct {
		e    entry
		want string
	}{
		{
			e: entry{
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
			e: entry{
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
			e: entry{
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
			e: entry{
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

func TestAdd(t *testing.T){
	tests := []struct {
		fd Funcdirectory
		want Funcdirectory
	}
}
