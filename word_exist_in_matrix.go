package algors

func exist(board [][]byte, word string) bool {
	pathIndex := 0
	var exists [][]bool
	for _, cols := range board {
		exists = append(exists, make([]bool, len(cols)))
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if existCore(board, exists, word, &pathIndex, i, j) {
				return true
			}
		}
	}
	return false
}

func existCore(board [][]byte,
	exists [][]bool,
	word string,
	pathIndex *int, col, row int) (ok bool) {
	if *pathIndex == len(word) {
		return true
	}

	if col >= 0 && col < len(board) &&
		row >= 0 && row < len(board[0]) &&
		board[col][row] == word[*pathIndex] &&
		!exists[col][row] {
		exists[col][row] = true
		*pathIndex++

		ok = existCore(board, exists, word, pathIndex, col, row-1) ||
			existCore(board, exists, word, pathIndex, col, row+1) ||
			existCore(board, exists, word, pathIndex, col-1, row) ||
			existCore(board, exists, word, pathIndex, col+1, row)
		if !ok {
			*pathIndex--
			exists[col][row] = false
			return false
		}
		return true
	}

	return false
}
