package main

import (
	"Sudoku/sudoku/large"
	"Sudoku/sudoku/medium"
	"Sudoku/sudoku/small"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

func main() {
	if len(os.Args) < 2 || len(os.Args) > 3 {
		fmt.Println("Usage: go run main.go <input_file> [output_file]")
		os.Exit(0)
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

	switch lines {
	case 16:
		board := large.ParseInput(os.Args[1])
		candidates := large.Preprocessing(&board)
		if large.Backtrack(&board, candidates) {
			fmt.Println("The 16x16 Sudoku has been solved. This is the result ->")
			large.PrintBoard(board)
		} else {
			fmt.Println("The 16x16 Sudoku can't be solved.")
		}
	case 9:
		board := medium.ParseInput(os.Args[1])
		candidates := medium.Preprocessing(&board)
		if medium.Backtrack(&board, candidates) {
			fmt.Println("The 9x9 Sudoku has been solved. This is the result ->")
			medium.PrintBoard(board)
		} else {
			fmt.Printf("The 9x9 Sudoku can't be solved.")
		}
	case 4:
		board := small.ParseInput(os.Args[1])
		if small.Backtrack(&board) {
			fmt.Println("The 4x4 Sudoku has been solved. This is the result ->")
			small.PrintBoard(board)
		} else {
			fmt.Println("The 4x4 Sudoku can't be solved.")
		}
	default:
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
