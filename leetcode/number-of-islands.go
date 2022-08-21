/*
- Problem: https://leetcode.com/problems/number-of-islands/
*/

type IslandFinder struct {
	grid         [][]byte
	totalIslands int
}

func numIslands(grid [][]byte) int {

	var isf IslandFinder = IslandFinder{
		grid:         grid,
		totalIslands: 0,
	}
  
        // Iterating over the grid to find a land
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == byte('1') {
                                // If an unmarked land if found then starting it as an island and marking all the lands of this island through DFS.
				isf.totalIslands++
				dfs(&isf, []int{i, j})
			} else {
				continue
			}
		}
	}

	return isf.totalIslands
}

func dfs(isf *IslandFinder, pos []int) {
	if isValidPos(isf, pos) {
                // Marking the land as visited
		isf.grid[pos[0]][pos[1]] = byte('2')

		dfs(isf, []int{pos[0], pos[1] + 1})
		dfs(isf, []int{pos[0], pos[1] - 1})
		dfs(isf, []int{pos[0] + 1, pos[1]})
		dfs(isf, []int{pos[0] - 1, pos[1]})
	}
}

func isValidPos(isf *IslandFinder, pos []int) bool {
        // Only adding an unmarked land postion to DFS
	if pos[0] >= 0 && pos[0] < len(isf.grid) && pos[1] >= 0 && pos[1] < len(isf.grid[0]) && isf.grid[pos[0]][pos[1]] == byte('1') {
		return true
	}
	return false
}
