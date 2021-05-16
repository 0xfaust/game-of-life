package main

import (
	"fmt"
	"math/rand"
	"time"
)

// cell struct with location and state (dead or alive)
type Cell struct {
	x     int
	y     int
	state int
}
// TODO: Refactor this out to match Go paradigms? 

// grid size and number of generations to simulate
const GRID_X, GRID_Y int = 10, 10
const GEN int = 10

// generate starting grid (using random seed)
func populateSeed() (grid [GRID_X][GRID_Y]int, err error) {
	// generate random seed
	rand.Seed(time.Now().UTC().UnixNano())

	// check grid size is valid
    if GRID_X < 0 {
		return grid, fmt.Errorf("grid dimensions must be a positive number")
    }

	// populate grid with random states
	for i := 0; i < GRID_X; i++ {
		for j := 0; j < GRID_Y; j++ {
			grid[i][j] = rand.Intn(2)
		}
	}

	return grid, err
}

// evaluate if a specified cell lives or dies
func lifeEvaluation(grid [GRID_X][GRID_Y]int, cell Cell) Cell {
	var alive int
	
	// count living neighbours
	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			if grid[cell.x-i][cell.y-j] == 1 {
				alive++
			}
		}
	}

	// any live cell with two or three live neighbours survives.
	if cell.state == 1 && (alive-cell.state == 2 || alive-cell.state == 3) {
		cell.state = 1
	// any dead cell with three live neighbours becomes a live cell.
	} else if cell.state == 0 && alive-cell.state == 3 {
		cell.state = 1
	// all other live cells die in the next generation. Similarly, all other dead cells stay dead.
	} else if cell.state == 1 {
		cell.state = 0
	}
	return cell
}

// populates a new grid with next generation states (ignores edges)
func nextGeneration(current_grid [GRID_X][GRID_Y]int) [GRID_X][GRID_Y]int {
	// new grid declared with each call (inefficient)
	var next_grid [GRID_X][GRID_Y]int
	var current Cell
	// populate new grid
	for i := 1; i < GRID_X-1; i++ {
		for j := 1; j < GRID_Y-1; j++ {
			current.state = current_grid[i][j]
			current.x = i
			current.y = j
			next_grid[i][j] = lifeEvaluation(current_grid, current).state
		}
	}
	return next_grid
	// TODO: declare two grids and swap between them for memory efficinecy
	// TODO: better method for dealing with grid edges
}

// prints grid along with generation count, x == alive, - == dead
func printGrid(grid [GRID_X][GRID_Y]int, gen int) {
	// print generation
	fmt.Println("generation", gen)
	
	// print grid
	for i := 0; i < GRID_X; i++ {
		for j := 0; j < GRID_Y; j++ {
			if grid[i][j] == 1 {
				// alive
				fmt.Print(" x ")
			} else {
				// dead
				fmt.Print(" - ")
			}
		}
		fmt.Println("")
	}
}

// run the simulation and print the results for specified gen count
func simulate(grid [GRID_X][GRID_Y]int, cell Cell) {
	var err error

	// generate random array
	grid, err = populateSeed()
	if err != nil {
		fmt.Println(err)
		return 
	}
	// print generation 0 grid
	printGrid(grid, 0)

	// simulate for gen generations
	for i := 1; i <= GEN; i++ {
		grid = nextGeneration(grid)
		printGrid(grid, i)
	}
}

func main() {
	// declare grid and cell
	var grid [GRID_X][GRID_Y]int
	var cell Cell

	// simulate game of life
	simulate(grid, cell)
}