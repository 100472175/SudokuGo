package main

import (
	large "Sudoku/sudoku/large"
	medium "Sudoku/sudoku/medium"
	small "Sudoku/sudoku/small"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <input_file>")
		os.Exit(1)
	}
	now := time.Now()
	defer func() {
		fmt.Println("Time taken: ", time.Since(now))
	}()

	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatalf("Error opening the file: %v", err)
	}
	defer file.Close()

	lines, err := lineCounter(file)
	if err != nil {
		log.Fatalf("Error counting the lines: %v", err)
	}

	if lines == 16 {
		board := large.ParseInput(os.Args[1])
		time_now := time.Now()
		large.PrintBoard(board)
		num0 := []int{-1, 0}
		large.PreprocessBoard(&board)
		num0 = append(num0, large.Count0(&board))
		var candidates *[16][16][]int
		for {
			candidates = large.PreprocessBoard(&board)
			if num0[len(num0)-1] == num0[len(num0)-2] {
				break
			} else {
				num0 = append(num0, large.Count0(&board))
			}
			fmt.Println("Number of unknowns: ", num0[len(num0)-1])
		}
		fmt.Println("Time taken for preprocessing: ", time.Since(time_now))
		fmt.Println("Board after preprocessing: ", num0[len(num0)-1])
		large.PrintBoard(board)

		if large.Backtrack(&board, candidates) {
			fmt.Println("The 16x16 Sudoku has been solved.")
			large.PrintBoard(board)
		} else {
			fmt.Println("The 16x16 Sudoku can't be solved.")
		}
		os.Exit(0)
	} else if lines == 9 {
		board := medium.ParseInput(os.Args[1])
		candidates := medium.PreprocessBoard(&board)
		if medium.Backtrack(&board, candidates) {
			fmt.Println("The 9x9 Sudoku has been solved.")
			medium.PrintBoard(board)
		} else {
			fmt.Printf("The 9x9 Sudoku can't be solved.")
		}
	} else if lines == 4 {
		board := small.ParseInput(os.Args[1])
		if small.Backtrack(&board) {
			fmt.Println("The 4x4 Sudoku has been solved.")
			small.PrintBoard(board)
		} else {
			fmt.Println("The 4x4 Sudoku can't be solved.")
		}
	} else {
		fmt.Println("Invalid number of columns. Please enter 4, 9, or 16 in the file.")
	}

}

func lineCounter(r io.Reader) (int, error) {
	buf := make([]byte, 32*1024)
	count := 0
	lineSep := []byte{'\n'}

	for {
		c, err := r.Read(buf)
		count += bytes.Count(buf[:c], lineSep)

		switch {
		case err == io.EOF:
			return count + 1, nil

		case err != nil:
			return count, err
		}
	}
}
