package main

func bfs(queue []*node, res []int) []int {
	if len(queue) == 0 {
		return res
	}
	res = append(res, queue[0].Val)
	if queue[0].R != nil {
		queue = append(queue, queue[0].R)
	}
	if queue[0].L != nil {
		queue = append(queue, queue[0].L)
	}
	return bfs(queue[1:], res)
}
