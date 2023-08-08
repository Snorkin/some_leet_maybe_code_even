package valid_sudoku

func isValidSudoku(board [][]byte) bool {

	var row, col, block [10][10]bool

	for i := range board {
		for j, c := range board[i] {
			if c == '.' {
				continue
			}
			// fmt.Println(c)
			// fmt.Println(string(c))
			// fmt.Println(c - '0')
			if col[j][c-'0'] || row[i][c-'0'] || block[i/3+j/3*3][c-'0'] {
				return false
			}
			col[j][c-'0'] = true
			row[i][c-'0'] = true
			block[i/3+j/3*3][c-'0'] = true
		}
	}
	return true
}

// --- initial solution
// func isValidSudoku(board [][]byte) bool {

// 	for i := 0; i < 9; i++ {
// 		for j := 0; j < 9; j++ {
// 			if board[i][j] == '.' {
// 				continue
// 			}

// 			if checkRows(i, j, board) || checkCols(i, j, board) || checkGrid(i, j, board) {
// 				return false
// 			}
// 		}
// 	}

// 	return true
// }

// func checkRows(x int, y int, board [][]byte) bool { // - cheking row by moving columns
// 	for i := x; i < 9; i++ {
// 		// fmt.Println("i - ", i, " y - ", y, " board xy : "+string(board[x][y])+" board xi : "+string(board[x][i]))
// 		// fmt.Println("x - ", x)
// 		if x == i {
// 			continue
// 		}
// 		if board[x][y] == board[i][y] {
// 			return true
// 		}
// 	}
// 	return false
// }

// func checkCols(x int, y int, board [][]byte) bool { // - cheking columns by moving rows
// 	for i := y; i < 9; i++ {
// 		if y == i {
// 			continue
// 		}
// 		if board[x][y] == board[x][i] {
// 			return true
// 		}
// 	}
// 	return false
// }

// func checkGrid(x int, y int, board [][]byte) bool { // - cheking grid 3 by 3 by moving both
// 	for i := 0; i < 3; i++ {
// 		for j := 0; j < 3; j++ {
// 			a := (x/3)*3 + i
// 			b := (y/3)*3 + j
// 			// fmt.Println("i - ", i, " j - ", j, " board xy : "+string(board[x][y])+" board xi : "+string(board[x][i]))
// 			// fmt.Println("x - ", x, " y - ", y, " a - ", a, " b - ", b)
// 			if x == a && y == b {
// 				continue
// 			}
// 			if board[x][y] == board[a][b] {
// 				return true
// 			}
// 		}
// 	}
// 	return false
// }
