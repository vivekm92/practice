package graphs

/*
	problem : https://www.interviewbit.com/problems/path-in-directed-graph/
*/

func dfs(u int, graph [][]int, visited []bool) {
	if visited[u] {
		return
	}

	visited[u] = true
	for _, v := range graph[u] {
		if !visited[v] {
			dfs(v, graph, visited)
		}
	}
}

// T(n) : O(|V| + |E|), S(n) : O(|V|)
func PathInDirectedGraph(A int, B [][]int) int {

	graph := make([][]int, A+1)
	for _, v := range B {
		graph[v[0]] = append(graph[v[0]], v[1])
	}

	visited := make([]bool, A+1)
	dfs(1, graph, visited)
	if visited[A] {
		return 1
	}

	return 0
}

func bfs(u int, graph [][]int, visited []bool) {

	q := make([]int, len(graph)+1)
	q = append(q, u)
	for len(q) > 0 {
		u = q[0]
		q = q[1:]
		visited[u] = true
		for _, v := range graph[u] {
			if !visited[v] {
				q = append(q, v)
			}
		}
	}
}

// T(n) : O(|V| + |E|), S(n) : O(|V|)
func PathInDirectedGraph1(A int, B [][]int) int {

	graph := make([][]int, A+1)
	for _, v := range B {
		graph[v[0]] = append(graph[v[0]], v[1])
	}

	visited := make([]bool, A+1)
	bfs(1, graph, visited)
	if visited[A] {
		return 1
	}

	return 0
}
