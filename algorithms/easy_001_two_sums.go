package algorithms

import "sort"

func twoSum(nums []int, target int) []int {
	numSorted := make([]int, len(nums))
	for i, num := range nums {
		numSorted[i] = num
	}
	sort.Ints(numSorted)

	numFirst, numSec := 0, 0
	found := false
	for i, num := range numSorted {
		if num < 0 {
			if numSorted[len(numSorted)-1] > target {
				numTarget := target - num
				for iNext := len(numSorted) - 1; iNext >= 0; iNext-- {
					numNext := numSorted[iNext]
					if numNext < numTarget {
						break
					} else if numNext == numTarget {
						found = true
						numFirst = num
						numSec = numNext
						break
					}
				}
			}
		} else {
			numTarget := target - num
			for _, numNext := range numSorted[i+1:] {
				if numNext > numTarget {
					break
				} else if numNext == numTarget {
					found = true
					numFirst = num
					numSec = numNext
					break
				}
			}
		}

		if found {
			break
		}
	}

	iFirst, iSec := -1, -1
	for i, num := range nums {
		if iFirst == -1 && num == numFirst {
			iFirst = i
			continue
		}
		if iSec == -1 && num == numSec {
			iSec = i
			continue
		}
		if iFirst != -1 && iSec != -1 {
			break
		}
	}
	if iFirst > iSec {
		iFirst, iSec = iSec, iFirst
	}

	return []int{iFirst, iSec}
}
