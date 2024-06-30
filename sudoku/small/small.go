package small

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const ROW_NUM = 4

func PrintBoard(board [ROW_NUM][ROW_NUM]int) {
	fmt.Println("+-----+-----+")
	for row := 0; row < ROW_NUM; row++ {
		fmt.Print("| ")
		for col := 0; col < ROW_NUM; col++ {
			if col == 2 {
				fmt.Print("| ")
			}
			fmt.Printf("%d ", board[row][col])
			if col == 3 {
				fmt.Print("|")
			}
		}
		if row == 1 || row == 3{
			fmt.Println("\n+-----+-----+")
		} else {
			fmt.Println()
		}
	}
}

func ParseInput(input string) [ROW_NUM][ROW_NUM]int {
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
	if line != ROW_NUM {
		content, err := os.ReadFile(input)
		if err != nil {
			fmt.Println("Error reading the file")
			os.Exit(1)
		}

		board := [ROW_NUM][ROW_NUM]int{}
		for i, c := range string(content) {
			row := i / ROW_NUM
			col := i % ROW_NUM
			if c == '.' {
				board[row][col] = 0
			} else {
				board[row][col] = int(c - '0')
			}
		}
		return board
	} else {
		// Parse the file as if was a ROW_NUMxROW_NUM sudoku board
		file, err := os.Open(input)
		if err != nil {
			log.Fatalf("Error opening the file: %v", err)
		}
		defer file.Close()

		board := [ROW_NUM][ROW_NUM]int{}
		scanner = bufio.NewScanner(file)
		for i := 0; i < ROW_NUM; i++ {
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

func HasDuplicates(counter [5]int) bool {
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

func IsBoardValid(board *[ROW_NUM][ROW_NUM]int) bool {

	//check duplicates by row
	for row := 0; row < ROW_NUM; row++ {
		counter := [5]int{}
		for col := 0; col < ROW_NUM; col++ {
			counter[board[row][col]]++
		}
		if HasDuplicates(counter) {
			return false
		}
	}

	//check duplicates by column
	for row := 0; row < ROW_NUM; row++ {
		counter := [5]int{}
		for col := 0; col < ROW_NUM; col++ {
			counter[board[col][row]]++
		}
		if HasDuplicates(counter) {
			return false
		}
	}

	//check 3x3 section
	for i := 0; i < ROW_NUM; i += 2 {
		for j := 0; j < ROW_NUM; j += 2 {
			counter := [5]int{}
			for row := i; row < i+2; row++ {
				for col := j; col < j+2; col++ {
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

func HasEmptyCell(board *[ROW_NUM][ROW_NUM]int) bool {
	for i := 0; i < ROW_NUM; i++ {
		for j := 0; j < ROW_NUM; j++ {
			if board[i][j] == 0 {
				return true
			}
		}
	}
	return false
}

func Backtrack(board *[ROW_NUM][ROW_NUM]int) bool {
	if !HasEmptyCell(board) {
		return true
	}
	for i := 0; i < ROW_NUM; i++ {
		for j := 0; j < ROW_NUM; j++ {
			if board[i][j] == 0 {
				for candidate := ROW_NUM; candidate >= 1; candidate-- {
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
