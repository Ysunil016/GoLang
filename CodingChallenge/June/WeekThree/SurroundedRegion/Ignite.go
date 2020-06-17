package main

import "fmt"

func main() {
	Arr := [][]byte{{'X', 'X', 'X', 'X'}, {'X', 'O', 'O', 'X'}, {'X', 'X', 'O', 'X'}, {'X', 'O', 'X', 'X'}}
	fmt.Println(Arr)
	solve(Arr)
	fmt.Println(Arr)
}
func solve(board [][]byte) {
	Row := len(board)
	if Row == 0 {
		return
	}
	Col := len(board[0])
	if Col == 0 {
		return
	}
	for I := 0; I < Row; I++ {
		if board[I][0] == 79 {
			DFS(board, I, 0)
		}
		if board[I][Col-1] == 79 {
			DFS(board, I, Col-1)
		}
	}
	for I := 0; I < Col; I++ {
		if board[0][I] == 79 {
			DFS(board, 0, I)
		}
		if board[Row-1][I] == 79 {
			DFS(board, Row-1, I)
		}
	}
	for I := 0; I < Row; I++ {
		for J := 0; J < Col; J++ {
			if board[I][J] == 79 {
				board[I][J] = 88
			} else if board[I][J] == 90 {
				board[I][J] = 79
			}
		}
	}

}

// DFS ...
func DFS(data [][]byte, I int, J int) {
	if I > len(data) || I < 0 || J > len(data[0]) || J < 0 {
		return
	}
	if data[I][J] == 79 {
		data[I][J] = 90
	}
	if I > 0 && data[I-1][J] == 79 {
		DFS(data, I-1, J)
	}
	if J > 0 && data[I][J-1] == 79 {
		DFS(data, I, J-1)
	}
	if I < len(data)-1 && data[I+1][J] == 79 {
		DFS(data, I+1, J)
	}
	if J < len(data[0])-1 && data[I][J+1] == 79 {
		DFS(data, I, J+1)
	}
}
