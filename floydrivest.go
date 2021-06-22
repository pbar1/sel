package sel

import (
	"math"
)

// FloydRivest finds the kth smallest element in an unordered list of ints.
// Also known as **SELECT**.
func FloydRivest(list []int, k int) []int {
	return floydrivest(list, 0, len(list)-1, k)
}

// See: https://stackoverflow.com/questions/29592546/floyd-rivest-vs-introselect-algorithm-performance
// Also: https://core.ac.uk/download/pdf/82672439.pdf
// Also: https://github.com/gliese1337/floyd-rivest/blob/master/src/floyd-rivest.ts
// TODO: according to the above stackoverflow post, the Wikipedia algorithm isn't optimal
func floydrivest(list []int, left, right, k int) []int {
	for right > left {
		if right-left > 600 {
			n := right - left + 1
			i := k - left + 1

			n_f := float64(n) // for easier calculation
			i_f := float64(i)
			k_f := float64(k)

			z := math.Log(n_f)
			s := 0.5 * math.Exp(2*z/3)
			sd := 0.5 * math.Sqrt(z*s*(n_f-s)/n_f) * float64(sign(i-n/2))

			newLeft := int(math.Max(float64(left), math.Floor(k_f-i_f*s/n_f+sd)))
			newRight := int(math.Min(float64(right), math.Floor(k_f+(n_f-i_f)*s/n_f+sd)))
			floydrivest(list, newLeft, newRight, k)
		}

		t := list[k]
		i := left
		j := right

		list[left], list[k] = list[k], list[left]
		if list[right] > t {
			list[right], list[left] = list[left], list[right]
		}
		for i < j {
			list[i], list[j] = list[j], list[i]
			i++
			j--

			for list[i] < t {
				i++
			}
			for list[j] > t {
				j--
			}
		}

		if list[left] == t {
			list[left], list[j] = list[j], list[left]
		} else {
			j++
			list[j], list[right] = list[right], list[j]
		}

		if j <= k {
			left = j + 1
		}
		if k <= j {
			right = j - 1
		}
	}

	return list[:k+1]
}

func sign(x int) int {
	if x < 0 {
		return -1
	}
	if x > 0 {
		return 1
	}
	return 0
}
