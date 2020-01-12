package main

func tribonacci(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else if n == 2 {
		return 1
	} else {
		answ := make([]int, n+1)
		answ[0] = tribonacci(0)
		answ[1] = tribonacci(1)
		answ[2] = tribonacci(2)
		for i := 3; i < n+1; i++ {
			answ[i] = answ[i-3] + answ[i-2] + answ[i-1]
		}
		return answ[n]
	}
}
