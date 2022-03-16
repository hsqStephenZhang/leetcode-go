package main

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func trapRainWater(heightMap [][]int) int {
	m := len(heightMap)
	if m == 0 {
		return 0
	}
	n := len(heightMap[0])
	aroundHeight := make([][]int, m+1)
	searched := make([][]bool, m+1)
	for i := 0; i <= len(aroundHeight); i++ {
		aroundHeight[i] = make([]int, n+1)
		searched[i] = make([]bool, n+1)
	}
	for i := 0; i < m; i++ {
		searched[i][0] = true
		searched[i][n-1] = true
	}
	for j := 0; j < n; j++ {
		searched[0][j] = true
		searched[m-1][j] = true
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if !searched[i][j] {
				aroundHeight[i][j] = dfs(heightMap, aroundHeight, searched, i, j)
			}
		}
	}

	return 0
}

func dfs(heightMap [][]int, aroundHeight [][]int, searched [][]bool, i, j int) int {
	if searched[i][j] && aroundHeight[i][j] != 0 {
		return aroundHeight[i][j]
	}
	searched[i][j] = true
	aroundHeight[i][j] = heightMap[i][j]
	maxHeight := heightMap[i][j]
	if i > 1 {
		maxHeight = max(maxHeight, dfs(heightMap, aroundHeight, searched, i-1, j))
	}
	if j > 1 {
		maxHeight = max(maxHeight, dfs(heightMap, aroundHeight, searched, i, j-1))
	}
	if i < len(heightMap) {
		maxHeight = max(maxHeight, dfs(heightMap, aroundHeight, searched, i+1, j))
	}
	if j < len(heightMap[0]) {
		maxHeight = max(maxHeight, dfs(heightMap, aroundHeight, searched, i, j+1))
	}
	aroundHeight[i][j] = maxHeight - heightMap[i][j]
	return maxHeight
}

func main() {

}
