package exercicios

import "fmt"

type File []bool

type Chessboard map[string]File

func CountInFile(cb Chessboard, file string) int {
	count := 0
	for _, val := range cb[file] {
		if val {
			count++
		}
	}
	return count
}

func CountInRank(cb Chessboard, rank int) int {
	if rank < 1 || rank > 8 {
		return 0
	}
	count := 0
	for _, file := range cb {
		if len(file) >= rank && file[rank-1] {
			count++
		}
	}
	return count
}

func CountAll(cb Chessboard) int {
	count := 0
	for _, file := range cb {
		count += len(file)
	}
	return count
}

func CountOccupied(cb Chessboard) int {
	count := 0
	for _, file := range cb {
		for _, val := range file {
			if val {
				count++
			}
		}
	}
	return count
}

func mais() {
	chessboard := Chessboard{
		"A": File{true, true, false, true, false, true, false, true},
		"B": File{false, false, false, false, true, false, false, false},
		"C": File{false, false, false, false, false, true, true, true},
		"D": File{true, false, false, false, false, false, true, true},
		"E": File{false, true, false, false, false, false, false, false},
		"F": File{true, true, true, false, false, false, false, true},
		"G": File{false, false, false, false, false, false, false, true},
		"H": File{true, true, true, true, true, true, true, true},
	}

	fmt.Println("Count in File A:", CountInFile(chessboard, "A"))
	fmt.Println("Count in Rank 2:", CountInRank(chessboard, 2))
	fmt.Println("Count All:", CountAll(chessboard))
	fmt.Println("Count Occupied:", CountOccupied(chessboard))
}