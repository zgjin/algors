package algors

func minDistance(word1 string, word2 string) int {
	var cache [][]int
	for i := 0; i < len(word1); i++ {
		rows := make([]int, len(word2))
		for j := 0; j < len(word2); j++ {
			rows[j] = 9999
		}
		cache = append(cache, rows)
	}
	return minDistancedp(cache, word1, word2, len(word1)-1, len(word2)-1)
}

func minDistancedp(cache [][]int, word1, word2 string, i, j int) int {
	if i < 0 {
		return j + 1
	}
	if j < 0 {
		return i + 1
	}

	if cache[i][j] != 9999 {
		return cache[i][j]
	}

	if word1[i] == word2[j] {
		return minDistancedp(cache, word1, word2, i-1, j-1)
	}

	cache[i][j] = minof3(
		minDistancedp(cache, word1, word2, i, j-1)+1,
		minDistancedp(cache, word1, word2, i-1, j)+1,
		minDistancedp(cache, word1, word2, i-1, j-1)+1,
	)
	return cache[i][j]
}
