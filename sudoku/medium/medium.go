package medium

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func PrintBoard(board [9][9]int) {
	fmt.Println("+-------+-------+-------+")
	for row := 0; row < 9; row++ {
		fmt.Print("| ")
		for col := 0; col < 9; col++ {
			if col == 3 || col == 6 {
				fmt.Print("| ")
			}
			fmt.Printf("%d ", board[row][col])
			if col == 8 {
				fmt.Print("|")
			}
		}
		if row == 2 || row == 5 || row == 8 {
			fmt.Println("\n+-------+-------+-------+")
		} else {
			fmt.Println()
		}
	}
}

func ParseInput(input string) [9][9]int {
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
	if line != 9 {
		content, err := os.ReadFile(input)
		if err != nil {
			fmt.Println("Error reading the file")
			os.Exit(1)
		}

		board := [9][9]int{}
		for i, c := range string(content) {
			row := i / 9
			col := i % 9
			if c == '.' {
				board[row][col] = 0
			} else {
				board[row][col] = int(c - '0')
			}
		}
		return board
	} else {
		// Parse the file as if was a 9x9 sudoku board
		file, err := os.Open(input)
		if err != nil {
			log.Fatalf("Error opening the file: %v", err)
		}
		defer file.Close()

		board := [9][9]int{}
		scanner = bufio.NewScanner(file)
		for i := 0; i < 9; i++ {
			scanner.Scan()
			line := scanner.Text()
			for j, c := range line {
				if c == '.' || c == 'x' || c == 'X' || c == '0' {
					board[i][j] = 0
				} else {
					board[i][j] = int(c - '0')
				}
			}
		}
		return board

	}
}

func HasDuplicates(counter [10]int) bool {
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

func IsBoardValid(board *[9][9]int) bool {

	//check duplicates by row
	for row := 0; row < 9; row++ {
		counter := [10]int{}
		for col := 0; col < 9; col++ {
			counter[board[row][col]]++
		}
		if HasDuplicates(counter) {
			return false
		}
	}

	//check duplicates by column
	for row := 0; row < 9; row++ {
		counter := [10]int{}
		for col := 0; col < 9; col++ {
			counter[board[col][row]]++
		}
		if HasDuplicates(counter) {
			return false
		}
	}

	//check 3x3 section
	for i := 0; i < 9; i += 3 {
		for j := 0; j < 9; j += 3 {
			counter := [10]int{}
			for row := i; row < i+3; row++ {
				for col := j; col < j+3; col++ {
					counter[board[row][col]]++
				}
				if HasDuplicates(counter) {
					return false
				}
			}
		}
	}

	return true
}

func HasEmptyCell(board *[9][9]int) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == 0 {
				return true
			}
		}
	}
	return false
}

func Backtrack(board *[9][9]int) bool {
	if !HasEmptyCell(board) {
		return true
	}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == 0 {
				for candidate := 9; candidate >= 1; candidate-- {
					board[i][j] = candidate
					if IsBoardValid(board) {
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
