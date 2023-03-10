package mergearraysandsort

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

var (
	impls = make([]mergeandsorter, 0)
)

func TestMergeArrays(t *testing.T) {
	impls = append(impls,
		mergeArraysS1{},
	)

	for _, impl := range impls {
		t.Logf("Using: %T", impl)

		td := getTestData()
		bytes, _ := os.ReadFile("problem.md")
		t.Logf("%s [#tests=%d]", bytes, len(td))

		for i, tt := range td {
			actualOpt := impl.mergeArrays(tt.input1, tt.input2)
			var pOrF string
			if assert.Equal(t, tt.expectedOpt, actualOpt) {
				pOrF = "\u2713"
			} else {
				pOrF = "\u2717"
			}

			t.Logf("test #%d, result: %s, input=%v and %v and expected output=%v", i+1, pOrF, tt.input1, tt.input2, tt.expectedOpt)
		}
	}
}

type ArrayTester struct {
	input1      []int
	input2      []int
	expectedOpt []int
}

func getTestData() []ArrayTester {
	return []ArrayTester{
		{[]int{0, 3, 4, 31}, []int{4, 6, 30}, []int{0, 3, 4, 4, 6, 30, 31}},
		{[]int{0, 3, 4, 31}, []int{0, 2, 3, 6, 9, 30}, []int{0, 0, 2, 3, 3, 4, 6, 9, 30, 31}},
		{[]int{0, 3}, []int{0, 2, 3, 6, 9, 30}, []int{0, 0, 2, 3, 3, 6, 9, 30}},
		{[]int{0, 3, 4}, []int{4}, []int{0, 3, 4, 4}},
	}
}
