package large

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

const (
	ROW_LEN = 16
)

var loops int

func ParseInput(input string) [ROW_LEN][ROW_LEN]int {
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

		board := [ROW_LEN][ROW_LEN]int{}
		scanner = bufio.NewScanner(file)
		for i := 0; i < ROW_LEN; i++ {
			scanner.Scan()
			line := scanner.Text()
			for j, c := range line {
				if c == '.' || c == 'x' || c == 'X' || c == '0' || c == ' ' {
					board[i][j] = 0
				} else {
					switch c {
					case 'A':
						board[i][j] = 10
					case 'B':
						board[i][j] = 11
					case 'C':
						board[i][j] = 12
					case 'D':
						board[i][j] = 13
					case 'E':
						board[i][j] = 14
					case 'F':
						board[i][j] = 15
					case 'G':
						board[i][j] = 16
					default:
						board[i][j] = int(c - '0')
					}
				}
			}
		}
		return board
	}
	return [ROW_LEN][ROW_LEN]int{}
}

func PrintBoard(board [ROW_LEN][ROW_LEN]int) {
	fmt.Println("+---------+---------+---------+---------+")
	for i, row := range board {
		fmt.Print("| ")
		for j, cell := range row {
			if cell == 0 {
				fmt.Print("0")
			} else {
				switch cell {
				case 10:
					fmt.Print("A")
				case 11:
					fmt.Print("B")
				case 12:
					fmt.Print("C")
				case 13:
					fmt.Print("D")
				case 14:
					fmt.Print("E")
				case 15:
					fmt.Print("F")
				case 16:
					fmt.Print("G")
				default:
					fmt.Print(string(cell + '0'))
				}
			}
			if j == 3 || j == 7 || j == 11 {
				fmt.Print(" |")
			}
			if j == 15 {
				fmt.Print(" |")
			} else {
				fmt.Print(" ")
			}

		}
		fmt.Println()
		if i == 3 || i == 7 || i == 11 || i == 15 {
			fmt.Println("+---------+---------+---------+---------+")
		}
	}
}

func isBoardValid(board *[ROW_LEN][ROW_LEN]int) bool {

	//check duplicates by row
	for row := 0; row < ROW_LEN; row++ {
		// Counter is a map, with the key being a rune and the value being the number of times that rune appears in the row
		counter := [ROW_LEN + 1]int{}
		for col := 0; col < ROW_LEN; col++ {
			counter[board[row][col]]++
		}
		if hasDuplicates(counter) {
			return false
		}
	}

	//check duplicates by column
	for row := 0; row < ROW_LEN; row++ {
		counter := [ROW_LEN + 1]int{}
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
			counter := [ROW_LEN + 1]int{}
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

func hasDuplicates(counter [ROW_LEN + 1]int) bool {
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

func hasEmptyCell(board *[ROW_LEN][ROW_LEN]int) bool {
	for i := 0; i < ROW_LEN; i++ {
		for j := 0; j < ROW_LEN; j++ {
			if board[i][j] == 0 {
				return true
			}
		}
	}
	return false
}

func Backtrack(board *[ROW_LEN][ROW_LEN]int, candidates *[ROW_LEN][ROW_LEN][]int) bool {
	loops++
	if loops%100_000 == 0 {
		fmt.Println("Llevo ", loops, " iteraciones")
	}
	if !hasEmptyCell(board) {
		return true
	}
	for i := 0; i < ROW_LEN; i++ {
		for j := 0; j < ROW_LEN; j++ {
			if board[i][j] == 0 {
				for _, candidate := range candidates[i][j] {
					board[i][j] = candidate
					if isBoardValid(board) {
						if Backtrack(board, candidates) {
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

func PreprocessBoard(board *[ROW_LEN][ROW_LEN]int) *[ROW_LEN][ROW_LEN][]int {
	candidatesBoard := [ROW_LEN][ROW_LEN][]int{}
	for i := 0; i < ROW_LEN; i++ {
		for j := 0; j < ROW_LEN; j++ {
			if board[i][j] != 0 {
				continue
			}
			candidates := calculateCandidates(board, i, j)
			for index, value := range candidates {
				if value == 0 {
					candidates = append(candidates, index)
				}
			}
			if len(candidates) == 1 {
				board[i][j] = candidates[0]
			}
			candidatesBoard[i][j] = candidates
		}
	}
	return &candidatesBoard
}

func calculateCandidates(board *[ROW_LEN][ROW_LEN]int, row, col int) []int {
	if board[row][col] != 0 {
		return []int{}
	}
	candidates := [17]int{}
	for i := 0; i < ROW_LEN; i++ {
		candidates[board[row][i]] = 1
		candidates[board[i][col]] = 1
	}
	startRow, startCol := row-row%4, col-col%4

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			candidates[board[startRow+i][startCol+j]] = 1
		}
	}
	result := []int{}
	for i := 1; i < 17; i++ {
		if candidates[i] == 0 {
			result = append(result, i)
		}
	}
	return result
}

func Count0(board *[ROW_LEN][ROW_LEN]int) int {
	count := 0
	for i := 0; i < ROW_LEN; i++ {
		for j := 0; j < ROW_LEN; j++ {
			if board[i][j] == 0 {
				count++
			}
		}
	}
	return count
}

func Preprocessing(board *[16][16]int) *[16][16][]int {
	time_now := time.Now()
	PrintBoard(*board)
	num0 := []int{-1, 0}
	PreprocessBoard(board)
	num0 = append(num0, Count0(board))
	var candidates *[16][16][]int
	for {
		candidates = PreprocessBoard(board)
		if num0[len(num0)-1] == num0[len(num0)-2] {
			break
		} else {
			num0 = append(num0, Count0(board))
		}
		fmt.Println("Number of unknowns: ", num0[len(num0)-1])
	}
	fmt.Println("Time taken for preprocessing: ", time.Since(time_now))
	fmt.Println("Board after preprocessing: ", num0[len(num0)-1])
	PrintBoard(*board)
	return candidates
}
