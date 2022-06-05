package main

import (
	"fmt"
)

func merge(listA, listB []int) []int {
	mergedList := make([]int, 0)
	idxA, idxB := 0, 0

	for i := 0; i < len(listA)+len(listB); i++ {
		// avoid indexing list if it is at the end.
		if idxA == len(listA) {
			mergedList = append(mergedList, listB[idxB:]...)
			return mergedList
		} else if idxB == len(listB) {
			mergedList = append(mergedList, listA[idxA:]...)
			return mergedList
		}

		fmt.Println(listA[idxA], listB[idxB])

		if listA[idxA] < listB[idxB] {
			mergedList = append(mergedList, listA[idxA])
			idxA += 1
		} else {
			mergedList = append(mergedList, listB[idxB])
			idxB += 1
		}
		fmt.Println("merged list:", mergedList)
	}

	return mergedList
}

func main() {
	fmt.Println(merge([]int{0, 3, 4, 31}, []int{4, 6, 30}))
}
