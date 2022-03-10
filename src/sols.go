package src

import (
	"sort"
)

func First(A []int, k int) []int {

	var res []int = A
	for i := 0; i < k; i++ {
		res = FirstSol(res)
	}
	return res

}

func FirstSol(arr []int) []int {
	l := len(arr)
	var b []int = make([]int, l)
	for i := 0; i < l; i++ {
		if i == (l - 1) {
			b[0] = arr[i]
		} else {
			b[i+1] = arr[i]
		}
	}
	return b
}

func Second(B []int) int {
	return SecondSol(B)
}

func SecondSol(A []int) int {
	len := len(A)
	var res int
	var z int
	var arr = make([]int, len)
	for i := 0; i < len; i++ {
		z = 0
		for j := 0; j < len; j++ {
			if A[i] == A[j] {
				z = z + 1
			}
		}
		arr[i] = z
	}
	for y, x := range arr {
		if x%2 == 1 {
			res = A[y]
		}
	}
	return res
}

func Third(A []int) int {
	return ThirdSol(A)
}

func ThirdSol(A []int) int {

	sort.Ints(A)
	for i := 0; i < len(A)-1; i++ {
		if A[i] != (A[i+1] - 1) {
			return 0
		}
	}
	return 1
}

func Fourth(A []int) int {
	return FourthSol(A)
}

func FourthSol(A []int) int {
	sort.Ints(A)
	var res int
	for i := 0; i < len(A)-1; i++ {
		if A[i] != (A[i+1] - 1) {
			res = A[i] + 1
		}
	}
	return res
}
