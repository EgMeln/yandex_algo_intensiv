package main

type node struct {
	Val int
	R   *node
	L   *node
}

func dfs(graph *node, sum *int) {
	if graph == nil {
		return
	}
	*sum += graph.Val

	dfs(graph.R, sum)
	dfs(graph.L, sum)
}
