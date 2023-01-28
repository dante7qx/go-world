package _algorithm

import (
	"fmt"
	"sort"
)

func BinarySearch(sourceData []int, searchVal int) (index int) {
	sort.Ints(sourceData)
	fmt.Println("先排序：", sourceData)
	var mid int
	start := 0
	lenArr := len(sourceData) - 1
	end := lenArr
	for start <= end {
		mid = (start + end) / 2
		if searchVal == sourceData[mid] {
			//fmt.Printf("%d found at location %d.\n", searchVal, mid+1)
			break
		} else if searchVal > sourceData[mid] {
			start = mid + 1
		} else if searchVal < sourceData[mid] {
			end = mid - 1
		}
	}
	if start > end {
		fmt.Printf("Not found! %d isn't present in the list.\n", searchVal)
	}
	return mid + 1
}
