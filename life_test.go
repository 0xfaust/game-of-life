package main

import (
	"testing"
)

func TestPopulateSeed(t *testing.T) {
    grid, err := populateSeed()
    for i := 0; i < GRID_X; i++ {
		for j := 0; j < GRID_Y; j++ {
			if (grid[i][j] != 0 && grid[i][j] != 1) ||err != nil {
                t.Errorf("state should be 0 or 1, got: %d.", grid[i][j])
            }
		}
	}
}

func TestFirstRule(t *testing.T) {
	// any live cell with two or three live neighbours survives.
	grid := [GRID_X][GRID_Y]int {
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 1, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 1, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 1, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	cell := Cell{
    	x: 2,
		y: 2,
	}
	if lifeEvaluation(grid, cell).state != 1 {
		t.Errorf("state should be 1, got: %d.", lifeEvaluation(grid, cell).state)
	}
}

func TestSecondRule(t *testing.T) {
	// any dead cell with three live neighbours becomes a live cell.
	grid := [GRID_X][GRID_Y]int {
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 1, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 1, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 1, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	cell := Cell{
    	x: 2,
		y: 3,
	}
	if lifeEvaluation(grid, cell).state != 1 {
		t.Errorf("state should be 1, got: %d.", lifeEvaluation(grid, cell).state)
	}
}

func TestThirdRule(t *testing.T) {
	// all other live cells die in the next generation. Similarly, all other dead cells stay dead.
	grid := [GRID_X][GRID_Y]int {
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 1, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 1, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 1, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	cell := Cell{
    	x: 1,
		y: 2,
	}
	if lifeEvaluation(grid, cell).state != 0 {
		t.Errorf("state should be 1, got: %d.", lifeEvaluation(grid, cell).state)
	}
}
