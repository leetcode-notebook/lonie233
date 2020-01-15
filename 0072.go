package main

func minDistance(word1 string, word2 string) int {

	var steps [][]int
	for i := 1; i <= len(word1)+1; i++ {
		steps = append(steps, make([]int, len(word2)+1))
	}

	for i := 1; i < len(word1)+1; i++ {
		steps[i][0] = i
	}

	for i := 1; i < len(word2)+1; i++ {
		steps[0][i] = i
	}

	for i := 1; i <= len(word1); i++ {
		for i2 := 1; i2 <= len(word2); i2++ {
			cas := 0
			if word1[i-1] == word2[i2-1] {
				cas = 0
			} else {
				cas = 1
			}

			rep := getTwoMinNumber(steps[i-1][i2]+1, steps[i][i2-1]+1)
			steps[i][i2] = getTwoMinNumber(steps[i-1][i2-1]+cas, rep)
		}
	}

	return steps[len(word1)][len(word2)]
}
