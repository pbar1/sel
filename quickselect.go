package sel

import "reflect"

// Quickselect finds the kth smallest element in an unordered list of ints.
// Also known as FIND, or Hoare's selection algorithm.
func Quickselect(list Interface, k int) interface{} {
	resultIndex := quickselectIndex(list, 0, list.Len()-1, k, partitionHoare)
	switch reflect.TypeOf(list).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(list)
		return s.Index(resultIndex).Interface()
	}
	return nil
}

func QuickselectInts(list []int, k int) int {
	result := Quickselect(IntList(list), k)
	return result.(int)
}

func QuickselectStrings(list []string, k int) string {
	result := Quickselect(StringList(list), k)
	return result.(string)
}

// quickselectIndex performs quickselect on the list, mutating it in the process,
// and returns the resulting index at which the kth smallest element resides.
func quickselectIndex(
	list Interface,
	left, right, k int,
	partition func(list Interface, left, right, pivotIndex int) int,
) int {
	for {
		if left == right {
			return k
		}

		// Choose a random pivot between left and right
		pivotIndex := RNG.Intn(right-left) + left

		// pivotIndex = partitionLomuto(list, left, right, pivotIndex)
		pivotIndex = partition(list, left, right, pivotIndex)

		if k == pivotIndex {
			return k
		} else if k < pivotIndex {
			right = pivotIndex - 1
		} else {
			left = pivotIndex + 1
		}
	}
}

// partitionHoare is Hoare's original partition scheme.
// See: https://en.wikipedia.org/wiki/Quicksort#Hoare_partition_scheme
func partitionHoare(list Interface, left, right, pivotIndex int) int {
	for i, j := left-1, right+1; ; {

		// Find leftmost element greater than or equal to pivot
		i++
		for list.Less(i, pivotIndex) {
			i++
		}

		// Find rightmost element less than or equal to pivot
		j--
		for list.Greater(j, pivotIndex) {
			j--
		}

		// If pointers meet
		if i >= j {
			return j
		}

		// Swap the values at each pointer
		list.Swap(i, j)
	}
}

// partitionLomuto is Lomuto's simplified partition scheme.
func partitionLomuto(list Interface, left, right, pivotIndex int) int {

	// Swap pivot to the end
	list.Swap(pivotIndex, right)

	storeIndex := left
	for i := left; i <= right-1; i++ {
		if list.Less(i, pivotIndex) {
			list.Swap(storeIndex, i)
			storeIndex++
		}
	}

	// Swap pivot into its final position
	list.Swap(right, storeIndex)

	return storeIndex
}
