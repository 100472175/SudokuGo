# Sudoku Solver in GO
This is a simple sudoku solver written in GO. It uses a backtracking algorithm to solve the sudoku puzzle. The program reads the puzzle from a file and writes the solution to another file.

This implementation can't solve sudoku puzzles of any size, only for 4x4, 9x9 and 16x16 puzzles. It has a specific implementation for each size, a go package for each size (small: 4x4, medium: 9x9, large: 16x16).

## How to run
1. Clone the repository
2. Run the following command to build the program
```bash
go build main.go
```
3. Run the program with the following command
```bash
./main <input_file> 
```
The input file should contain the sudoku puzzle.

This file must have the following format:
- n lines with n numbers each, separated by spaces
- Use 0/x/X to represent empty cells
- The number of lines must be equal to the number of columns
- The number of columns must be a perfect square (4, 9, 16)

Note: The program when solving a 16x16 sudoku puzzle can take a long time to finish, specially if there is not much information in the input file. (It can take up to 2 hours to solve a 16x16 sudoku puzzle with only a few numbers in the input file)

## Example
You can run the program with the provided examples in the "input" folder.
