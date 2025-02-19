package main

import "fmt"

func main() {
	fmt.Println("Game of Life")
	// initialize the game

	game := [][]int{
		{1, 0, 1},
		{0, 1, 0},
		{0, 0, 1},
	}
	fmt.Println("initial")
	printGame(game)

	for i := 0; i < 4; i++ {
		fmt.Println("generation", i+1)
		game = nextGeneration(game)
		printGame(game)
	}
}

func nextGeneration(game [][]int) [][]int {
	newGame := make([][]int, len(game))

	for i := range game {
		newGame[i] = make([]int, len(game[i]))
	}

	for i := range game {
		for j := range game[i] {
			newGame[i][j] = isAlive(game, i, j)
		}
	}

	return newGame
}

func isAlive(game [][]int, i, j int) int {
	aliveNeighbours := 0

	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			if x == 0 && y == 0 {
				continue
			}

			if i+x < 0 || i+x >= len(game) {
				continue
			}

			if j+y < 0 || j+y >= len(game[i]) {
				continue
			}

			if game[i+x][j+y] == 1 {
				aliveNeighbours++
			}
		}
	}

	if game[i][j] == 1 {
		if aliveNeighbours == 2 || aliveNeighbours == 3 {
			return 1
		}
		return 0
	}

	if aliveNeighbours == 3 {
		return 1
	}

	return 0
}

// function to print the game
func printGame(game [][]int) {
	for i := range game {
		for j := range game[i] {
			fmt.Print(game[i][j], " ")
		}
		fmt.Println()
	}
}

func DynamicSlice(row, col int) {
	fmt.Println("Dynamic Slice")

	cell := make([][]int, row)
	for i := range cell {
		cell[i] = make([]int, col)
	}
}
