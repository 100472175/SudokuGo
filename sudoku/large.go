package sudoku

import (
	"bufio"
	"fmt"
	"log"

	"os"
)

const (
	ROW_LEN    = 16
	candidates = "GFEDCBA987654321"
)
var loops int

func ParseInput(input string) [ROW_LEN][ROW_LEN]rune {
	file, err := os.Open(input)
	if err != nil {
		log.Fatalf("Error opening the file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	line := 0
	for scanner.Scan() {
		line++
	}
	if line != ROW_LEN {
		fmt.Println("Error: The file does not contain a parseable ROW_LENxROW_LEN sudoku board")
	} else {
		// Parse the file as if was a ROW_LENxROW_LEN sudoku board
		file, err := os.Open(input)
		if err != nil {
			log.Fatalf("Error opening the file: %v", err)
		}
		defer file.Close()

		board := [ROW_LEN][ROW_LEN]rune{}
		scanner = bufio.NewScanner(file)
		for i := 0; i < ROW_LEN; i++ {
			scanner.Scan()
			line := scanner.Text()
			for j, c := range line {
				if c == '.' || c == 'x' || c == 'X' || c == '0' {
					board[i][j] = 0
				} else {
					board[i][j] = c
				}
			}
		}
		return board
	}
	return [ROW_LEN][ROW_LEN]rune{}
}

func PrintBoard(board [ROW_LEN][ROW_LEN]rune) {
	for i, row := range board {
		for j, cell := range row {
			if cell == 0 {
				fmt.Print("0")
			} else {
				fmt.Print(string(cell))
			}
			if j == 3 || j == 7 || j == 11 {
				fmt.Print(" | ")
			}
		}
		fmt.Println()
		if i == 3 || i == 7 || i == 11 {
			fmt.Println("-----+------+------+------+")
		}
	}
}

func isBoardValid(board *[ROW_LEN][ROW_LEN]rune) bool {

	//check duplicates by row
	for row := 0; row < ROW_LEN; row++ {
		// Counter is a map, with the key being a rune and the value being the number of times that rune appears in the row
		counter := make(map[rune]int)
		for col := 0; col < ROW_LEN; col++ {
			counter[board[row][col]]++
		}
		if hasDuplicates(counter) {
			return false
		}
	}

	//check duplicates by column
	for row := 0; row < ROW_LEN; row++ {
		counter := make(map[rune]int)
		for col := 0; col < ROW_LEN; col++ {
			counter[board[col][row]]++
		}
		if hasDuplicates(counter) {
			return false
		}
	}

	//check 3x3 section
	for i := 0; i < ROW_LEN; i += 4 {
		for j := 0; j < ROW_LEN; j += 4 {
			counter := make(map[rune]int)
			for row := i; row < i+4; row++ {
				for col := j; col < j+4; col++ {
					counter[board[row][col]]++
				}
				if hasDuplicates(counter) {
					return false
				}
			}
		}
	}

	return true
}

func hasDuplicates(counter map[rune]int) bool {
	for i, count := range counter {
		if i == 0 {
			continue
		}
		if count > 1 {
			return true
		}
	}
	return false
}

func hasEmptyCell(board *[ROW_LEN][ROW_LEN]rune) bool {
	for i := 0; i < ROW_LEN; i++ {
		for j := 0; j < ROW_LEN; j++ {
			if board[i][j] == 0 {
				return true
			}
		}
	}
	return false
}

func Backtrack(board *[ROW_LEN][ROW_LEN]rune) bool {
	loops++
	if loops % 100_000 == 0{
		fmt.Println("Llevo ", loops, " iteraciones")
	}
	if !hasEmptyCell(board) {
		return true
	}
	for i := 0; i < ROW_LEN; i++ {
		for j := 0; j < ROW_LEN; j++ {
			if board[i][j] == 0 {
				for _, candidate := range candidates {
					board[i][j] = candidate
					if isBoardValid(board) {
						if Backtrack(board) {
							return true
						}
						board[i][j] = 0
					} else {
						board[i][j] = 0
					}
				}
				return false
			}
		}
	}
	return false
}
