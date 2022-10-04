package gutil

import (
	"testing"

	. "github.com/stretchr/testify/require"
)

func TestToMap(t *testing.T) {

	type Name struct {
		ID        int
		FirstName string
		LastName  string
	}

	type args[T any, V comparable] struct {
		slice []T
		key   func(*T) V
	}

	type test[T any, V comparable] struct {
		name string
		args args[T, V]
	}
	tests := []test[Name, int]{
		{
			name: "Array to map test",
			args: args[Name, int]{
				slice: []Name{
					{1, "Shahriar", "Ahmed"},
					{2, "Alamin", "Ahmed"},
					{3, "Shorif", "Ahmed"},
					{4, "Saim", "Ahmed"},
				},
				key: func(n *Name) int { return n.ID },
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ToMap(tt.args.slice, tt.args.key)
			for _, v := range tt.args.slice {
				Equal(t, v.FirstName, got[v.ID].FirstName)
				Equal(t, v.LastName, got[v.ID].LastName)
			}

			tt.args.slice[0].FirstName = "Sazid"
			Equal(t, got[1].FirstName, tt.args.slice[0].FirstName)
		})
	}
}
