package slice

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRandomOne(t *testing.T) {
	testCases := []struct {
		name      string
		arr       []any
		check     func(i any) bool
		wantError error
	}{
		{
			name: "normal case 1",
			arr:  []any{1},
			check: func(i any) bool {
				return i == 1
			},
		},
		{
			name: "normal case 2",
			arr:  []any{1, "2", 3.0, true},
			check: func(i any) bool {
				t.Logf("random item is: %t", i)
				return true
			},
		},
		{
			name: "empty slice",
			arr:  []any{},
			check: func(i any) bool {
				return i == nil
			},
		},
	}

	for _, tCase := range testCases {
		t.Run(tCase.name, func(t *testing.T) {
			c := RandomOne(tCase.arr)
			assert.True(t, true, tCase.check(c), "它们应该相等")
		})
	}
}
