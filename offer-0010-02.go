package main

func numWays(n int) int {

	f := make([]int, 2)
	f[0] = 1
	f[1] = 1

	if n < 2 {
		return f[n]
	}

	f1 := 1
	f2 := 1
	fn := f1 + f2

	for i := 2; i <= n; i++ {
		fn = f1 + f2
		if fn > 1000000007 {
			fn -= 1000000007
		}
		f1 = f2
		f2 = fn
	}

	return fn

}
