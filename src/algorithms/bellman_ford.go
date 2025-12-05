/* algorithms: Bellman-Ford */
// a shortest path algorithm (Graph).
package main

import (
	"errors"
	"fmt"
	"math"
)

/*
	Analysis on Complexity

- Time: O(V.E) average/worst case.
- Space: O(V + E).
*/
func bellmanFord() {
	bellmanFord := &BellmanFord{}
	V := 5
	edges := [][]int{
		{0, 1, 6}, {0, 2, 5}, {1, 3, -1}, {2, 1, -2},
		{2, 3, 4}, {2, 4, 3}, {3, 4, 3},
	}
	src := 1

	ans, err := bellmanFord.shortestPath(V, edges, src)
	fmt.Printf("Shortest distances from source vertex %d:\n", src)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	for i := range V {
		if math.IsInf(ans[i], 1) {
			fmt.Printf("Distance to vertex %d: Unreachable\n", i)
		} else {
			fmt.Printf("Distance to vertex %d: %.2f\n", i, ans[i])
		}
	}
}

func (bf *BellmanFord) shortestPath(V int, edges [][]int, src int) ([]float64, error) {
	distTo := make([]float64, V)
	distTo[src] = 0.0
	posInf := math.Inf(1)

	for i := 1; i < V; i++ {
		distTo[i] = posInf
	}

	for i := 1; i < V; i++ {
		// Early termination flag
		hasChanged := false

		for _, edge := range edges {
			u, v, w := edge[0], edge[1], edge[2]

			// Only continue computing on computed vertices
			if distTo[u] != posInf {
				newW := distTo[u] + float64(w)

				if newW < distTo[v] {
					distTo[v] = newW // relax edge
					hasChanged = true
				}
			}
		}

		// No more updates to distTo[] -> Finish!
		if !hasChanged {
			return distTo, nil
		}
	}

	for _, edge := range edges {
		u, v, w := edge[0], edge[1], float64(edge[2])

		if distTo[u] != posInf && distTo[u]+w < distTo[v] {
			return nil, errors.New("Exist negative cycle(s) in input Graph.")
		}
	}

	return distTo, nil
}

type BellmanFord struct{}
