package main

//逆向双指针
func mergeSortArray(A []int, m int, B []int, n int) {
	m = m - 1
	n = n - 1
	for m > -1 || n > -1 {

		if m < 0 {
			A[m+n+1] = B[n]
			n--
		} else if n < 0 {
			return
		} else if A[m] > B[n] {
			A[m+n+1] = A[m]
			m--
		} else {
			A[m+n+1] = B[n]
			n--
		}
	}
}
