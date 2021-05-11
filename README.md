# game-of-life

![Release Badge](https://img.shields.io/github/v/tag/0xfaust/game-of-life?color=red&label=release&sort=semver) 
![Site Status Badge](https://img.shields.io/website?down_message=down&label=api+status&up_message=up&url=http://0xfaust.io:8000/game-of-life)  
   
A distributed implementation of [Conway's Game of Life](https://en.wikipedia.org/wiki/Conway's_Game_of_Life) written in Go. The Game of Life is a cellular automaton devised by the mathematician John Horton Conway in 1970. It is a zero-player game, meaning that its evolution is determined by its initial state, requiring no further input. One interacts with the Game of Life by creating an initial configuration and observing how it evolves. It is Turing complete and can simulate a universal constructor or any other Turing machine.

## Rules

1. Any live cell with fewer than two live neighbours dies, as if by **underpopulation**.
2. Any live cell with two or three live neighbours lives on to the next generation.
3. Any live cell with more than three live neighbours dies, as if by **overpopulation**.
4. Any dead cell with exactly three live neighbours becomes a live cell, as if by **reproduction**.

## Versioning

This project is currently in the early development stage. A basic [SemVer](http://semver.org/) system is used for versioning. For the versions available, see the [tags and releases](https://github.com/0xfaust/game-of-life/releases) on this repository.
