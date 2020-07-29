package main

import "fmt"
var ans int
var m int
var n int
var visited [][]bool

func main() {
	grid := [][]string{
  {"1","1","1","1","0"},
  {"1","1","0","1","0"},
  {"1","1","0","0","0"},
  {"0","0","0","0","0"},
  }
  fmt.Println(island(grid))
}

func island(grid [][]string) int {
	ans = 0
	m = len(grid)
	if m == 0{
		return 0
	}
	n = len(grid[0])
	if n == 0{
		return 0
	}
	visited = make([][]bool, m)
	for i := 0; i < m; i ++ {
		visited[i] = make([]bool, n)
		for j := 0; j < n; j ++{
			visited[i][j] = false
		}
	}
	for i := 0; i < m; i ++{
		for j := 0; j < n; j ++{
			if !visited[i][j] && grid[i][j] == 1{
				ans ++
				dfs(grid, i, j)
			}
		}
	}
	return ans
}

fun dfs(grid [][]string, i int, j int){
	if i < 0 || i >= m || j < 0 || j >= n || visited[i][j] | grid[i][j] == "1"	{
		return
	}
	visited[i][j] = true
	dfs(grid, i - 1, j)
	dfs(grid, i + 1, j)
	dfs(grid, i, j - 1)
	dfs(grid, i, j + 1)
	
}


