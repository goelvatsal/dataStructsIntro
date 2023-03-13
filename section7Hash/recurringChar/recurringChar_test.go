package recurringchar

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

var (
	impls = make([]RecurringChar, 0)
)

func TestMergeArrays(t *testing.T) {
	impls = append(impls,
		recurringCharS1{},
	)

	for _, impl := range impls {
		t.Logf("Using: %T", impl)

		td := getTestData()
		bytes, _ := os.ReadFile("problem.md")
		t.Logf("%s [#tests=%d]", bytes, len(td))

		for i, tt := range td {
			actualOpt := impl.findRecurringChar(tt.input1)
			var pOrF string
			if assert.Equal(t, tt.firstRecurring, actualOpt) {
				pOrF = "\u2713"
			} else {
				pOrF = "\u2717"
			}

			t.Logf("test #%d, result: %s, input=%v and expected output=%v", i+1, pOrF, tt.input1, tt.firstRecurring)
		}
	}
}

type ArrayTester struct {
	input1         []int
	firstRecurring int
}

func getTestData() []ArrayTester {
	return []ArrayTester{
		{[]int{2, 5, 1, 2, 3, 5, 1, 2, 4}, 2},
		{[]int{2, 1, 1, 2, 3, 5, 1, 2, 4}, 1},
		{[]int{8, 3, 6, 7, 4, 9, 8, 2, 5}, 8},
		{[]int{6, 4, 3, 2, 7, 2, 7, 4, 6}, 2},
		{[]int{2, 3, 4, 5}, -1},
		{[]int{7, 2, 3, 9}, -1},
	}
}
