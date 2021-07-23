package main

func findCircleNum(isConnected [][]int) int {
	visited := make([]int, len(isConnected))
	count := 0
	for i := 0; i < len(isConnected); i++ {
		// have we visited the province yet?
		if visited[i] == 0 {
			dfs(&isConnected, &visited, i)
			// mark it when we do
			count++
		}
	}

	return count
}

func dfs(M *[][]int, visited *[]int, i int) {
	for j := 0; j < len(*M); j++ {
		if (*M)[i][j] == 1 && (*visited)[j] == 0 {
			(*visited)[j] = 1
			dfs(M, visited, j)
		}
	}
}

// adjacency matrix
// each row indicates a node
// the first column of the first row, indicates a connection to row 0
// so in an adjacency matrix, self is always connected
// the second column indicates a connection to the 1st row, and so on
