package main

import (
	"Sudoku/sudoku"
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run main.go <input_file> <number-columns>")
		os.Exit(1)
	}
	now := time.Now()
	defer func() {
		fmt.Println("Time taken: ", time.Since(now))
	}()

	if os.Args[2] != "9" {
		board := sudoku.ParseInput(os.Args[1])
		sudoku.PrintBoard(board)

		if sudoku.Backtrack(&board) {
			fmt.Println("The 16x16 Sudoku has been solved.")
			sudoku.PrintBoard(board)
		} else {
			fmt.Println("The 16x16 Sudoku can't be solved.")
		}

		os.Exit(0)
	}
	board := parseInput(os.Args[1])
	if backtrack(&board) {
		fmt.Println("The 9x9 Sudoku has been solved.")
		printBoard(board)
	} else {
		fmt.Printf("The 9zx9 Sudoku can't be solved.")
	}

}

func printBoard(board [9][9]int) {
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

func parseInput(input string) [9][9]int {
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

func hasDuplicates(counter [10]int) bool {
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

func isBoardValid(board *[9][9]int) bool {

	//check duplicates by row
	for row := 0; row < 9; row++ {
		counter := [10]int{}
		for col := 0; col < 9; col++ {
			counter[board[row][col]]++
		}
		if hasDuplicates(counter) {
			return false
		}
	}

	//check duplicates by column
	for row := 0; row < 9; row++ {
		counter := [10]int{}
		for col := 0; col < 9; col++ {
			counter[board[col][row]]++
		}
		if hasDuplicates(counter) {
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
				if hasDuplicates(counter) {
					return false
				}
			}
		}
	}

	return true
}

func hasEmptyCell(board *[9][9]int) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == 0 {
				return true
			}
		}
	}
	return false
}

func backtrack(board *[9][9]int) bool {
	if !hasEmptyCell(board) {
		return true
	}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == 0 {
				for candidate := 9; candidate >= 1; candidate-- {
					board[i][j] = candidate
					if isBoardValid(board) {
						if backtrack(board) {
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
