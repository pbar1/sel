package sel

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Quickselect finds the kth smallest element in an unordered list of ints.
// Also known as **FIND** and **Hoare's selection algorithm**.
func Quickselect(list []int, k int) []int {
	return quickselect(list, 0, len(list)-1, k, partitionHoare)
}

func quickselect(
	list []int,
	left, right, k int,
	partition func(list []int, left, right, pivotIndex int) int,
) []int {
	for {
		if left == right {
			return list[:left+1]
		}

		// Choose a random pivot between left and right
		pivotIndex := rand.Intn(right-left) + left

		// pivotIndex = partitionLomuto(list, left, right, pivotIndex)
		pivotIndex = partition(list, left, right, pivotIndex)

		if k == pivotIndex {
			return list[:k+1]
		} else if k < pivotIndex {
			right = pivotIndex - 1
		} else {
			left = pivotIndex + 1
		}
	}
}

// partitionLomuto is Lomuto's simplified partition scheme.
func partitionLomuto(list []int, left, right, pivotIndex int) int {
	pivot := list[pivotIndex]

	// Swap pivot to the end
	list[pivotIndex], list[right] = list[right], list[pivotIndex]

	storeIndex := left
	for i := left; i <= right-1; i++ {
		if list[i] < pivot {
			list[storeIndex], list[i] = list[i], list[storeIndex]
			storeIndex++
		}
	}

	// Swap pivot into its final position
	list[right], list[storeIndex] = list[storeIndex], list[right]

	return storeIndex
}

// partitionHoare is Hoare's original partition scheme.
// See: https://en.wikipedia.org/wiki/Quicksort#Hoare_partition_scheme
func partitionHoare(list []int, left, right, pivotIndex int) int {
	pivot := list[pivotIndex]

	for i, j := left-1, right+1; ; {

		// Find leftmost element greater than or equal to pivot
		i++
		for list[i] < pivot {
			i++
		}

		// Find rightmost element less than or equal to pivot
		j--
		for list[j] > pivot {
			j--
		}

		// If pointers meet
		if i >= j {
			return j
		}

		// Swap the values at each pointer
		list[i], list[j] = list[j], list[i]
	}
}
