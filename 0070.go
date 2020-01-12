package main

func climbStairs(n int) int {
	if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	}

	types := make([]int,n)
	types[0] = 1
	types[1] = 2

	for m :=2; m<n ;m++  {
		types[m] = types[m-1]+types[m-2]
	}

	return types[n-1]
}
