package algorithms

import (
	"reflect"
	"testing"
)

func TestTwoSums(t *testing.T) {
	cases := []struct {
		inNums   []int
		inTarget int
		want     []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
	}
	for _, c := range cases {
		got := twoSum(c.inNums, c.inTarget)
		if !reflect.DeepEqual(c.want, got) {
			t.Errorf("twoSum(%v) == %v, want %v", c.inNums, got, c.want)
		}
	}
}
