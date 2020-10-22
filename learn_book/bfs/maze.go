package main

import (
	"fmt"
	"os"
)

func readMaze(filename string) [][]int {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	var row, col int
	fmt.Fscanf(file, "%d %d", &row, &col)

	maze := make([][]int, row)
	for i := range maze {
		maze[i] = make([]int, col)
		for j := range maze[i] {
			fmt.Fscanf(file, "%d", &maze[i][j])
		}
	}
	return maze
}

type point struct {
	i, j int
}

var dirs = [4]point{
	{-1,0}, {0,-1}, {1,0}, {0,1},
}

func (p point) add(d point) point {
	return point{p.i + d.i, p.j + d.j}
}

// 取matrix p点的值
func (p point) at(matrix [][]int) (int, bool) {
	if p.i < 0 || p.i >= len(matrix) {
		return 0, false
	}

	if p.j < 0 || p.j >= len(matrix[p.i]) {
		return 0, false
	}

	return matrix[p.i][p.j], true
}

func walk(maze [][]int, start, end point) ([][]int, bool) {
	steps := make([][]int, len(maze))
	for i := range steps {
		steps[i] = make([]int, len(maze[i]))
	}

	Q := []point{start}

	for len(Q) > 0 {
		cur := Q[0]
		Q = Q[1:]

		if cur == end {
			return steps, true
		}

		for _, dir := range dirs {
			next := cur.add(dir)

			// 判断next是否有效

			if next == start {
				continue
			}

			val, ok := next.at(maze)
			if !ok || val == 1 {
				continue
			}

			val, ok = next.at(steps)
			if !ok || val != 0 {
				continue
			}

			steps[next.i][next.j] = steps[cur.i][cur.j] + 1
			Q = append(Q, next)
		}
	}

	return steps, false
}

func printMatrix(maze [][]int) {
	for i := 0; i < len(maze); i++ {
		for j := 0; j < len(maze[i]); j++ {
			fmt.Printf("%3d", maze[i][j])
		}
		fmt.Println()
	}
}

func main() {
	maze := readMaze("go_learning/src/gopl.io/learn_book/bfs/maze.in")
	start, end := point{0, 0}, point{len(maze)-1, len(maze[0])-1}
	steps, ok := walk(maze, start, end)
	fmt.Println(ok)
	printMatrix(steps)

	cur := end
	path := []point{cur}

	for cur != start {
		val, _ := cur.at(steps)
		for _, dir := range dirs {
			prev := cur.add(dir)
			pval, ok := prev.at(steps)
			if !ok {
				continue
			}

			if pval == val - 1 {
				cur = prev
				path = append(path, prev)
				break
			}
		}
	}

	for i := len(path)-1; i >= 0; i-- {
		fmt.Printf("%d ", path[i])
	}
}
