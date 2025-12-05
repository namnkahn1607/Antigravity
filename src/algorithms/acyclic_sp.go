/* algorithms: Acyclic Shortest Path */
// a shortest path algorithm (Diraph).
package main

import (
	"fmt"
	"math"
)

/*
	Analysis on Complexity

- Time: O(V + E) average/worst case.
- Space: O(V + E).
*/
func acyclicSP() {
	acyclicSP := &AcyclicSP{}
	V := 5
	edges := [][]int{
		{0, 2, 4}, {1, 2, 2}, {1, 0, 3}, {1, 3, 5},
		{2, 4, 2}, {2, 3, -3}, {4, 3, 2},
	}

	ans := acyclicSP.shortestPath(V, edges)

	for i := range V {
		fmt.Printf("Distance to vertex %d: %.2f", i, ans[i])
	}
}

func (asp *AcyclicSP) shortestPath(V int, edges [][]int) []float64 {
	adj := make([][][2]int, V)

	for _, edge := range edges {
		u, v, w := edge[0], edge[1], edge[2]
		adj[u] = append(adj[u], [2]int{v, w})
	}

	// Get vertices in topological order.
	topo := &Topological{}
	order := topo.topoOrder(V, edges)

	distTo := make([]float64, V)
	distTo[order[0]] = 0.0
	posInf := math.Inf(1)

	// Initialize all paths positive infinity cost.
	for i := 1; i < V; i++ {
		distTo[i] = posInf
	}

	for _, u := range order {
		if distTo[u] == posInf {
			continue
		}

		for _, neighbor := range adj[u] {
			v, w := neighbor[0], float64(neighbor[1])
			newW := distTo[u] + w
			distTo[v] = math.Min(distTo[v], newW) // relax edge
		}
	}

	return distTo
}

type AcyclicSP struct{}
