/*
- Problem: https://leetcode.com/problems/flood-fill/
*/

type FloodFill struct {
	image      [][]int
	queue      [][]int
	hashSet    map[string]struct{}
	startColor int
}

func floodFill(image [][]int, sr int, sc int, color int) [][]int {

	var ff FloodFill = FloodFill{
		image:   image,
		queue:   [][]int{},
		hashSet: map[string]struct{}{},
	}

	if image[sr][sc] == color {
		return image
	}

	ff.startColor = image[sr][sc]
	ff.queue = append(ff.queue, []int{sr, sc})

	doBfs(&ff, color)

	return ff.image

}

func doBfs(ff *FloodFill, color int) {
	if len(ff.queue) > 0 {
		pos := ff.queue[0]
		ff.queue = ff.queue[1:]

		// Changing color and adding the pixel to hash set to save it as done
		ff.image[pos[0]][pos[1]] = color
		ff.hashSet[fmt.Sprintf("%d-%d", pos[0], pos[1])] = struct{}{}

		// adding 4 directional pixels to the bfs queue
		addToQueue(ff, []int{pos[0], pos[1] + 1}, color)
		addToQueue(ff, []int{pos[0], pos[1] - 1}, color)
		addToQueue(ff, []int{pos[0] - 1, pos[1]}, color)
		addToQueue(ff, []int{pos[0] + 1, pos[1]}, color)

		doBfs(ff, color)
	}
}

func addToQueue(ff *FloodFill, pos []int, color int) {
	// checking if the pixel is already colored or eligible to be colored
	if _, ok := ff.hashSet[fmt.Sprintf("%d-%d", pos[0], pos[1])]; !ok {
		if pos[0] >= 0 && pos[1] >= 0 && pos[0] < len(ff.image) && pos[1] < len(ff.image[0]) && (ff.image[pos[0]][pos[1]] == ff.startColor) {
			ff.queue = append(ff.queue, pos)
		}
	}
}
