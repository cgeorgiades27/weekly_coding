/*
### Week 34, 2021 - (Medium)

This problem was asked by Spotify.

You have access to ranked lists of songs for various users. Each song is represented as an integer, and more preferred songs appear earlier in each list. For example, the list `[4, 1, 7]` indicates that a user likes song `4` the best, followed by songs `1` and `7`.

Given a set of these ranked lists, interleave them to create a playlist that satisfies everyone's priorities.

For example, suppose your input is `{[1, 7, 3], [2, 1, 6, 7, 9], [3, 9, 5]}`. In this case a satisfactory playlist could be `[2, 1, 6, 7, 3, 9, 5]`.
*/

package main

import "fmt"

// Graph data structure
type Graph map[int][]int

func main() {
	playlists := [][]int{{1, 7, 3}, {2, 1, 6, 7, 9}, {3, 9, 5}}
	graph := make(Graph)
	visited := make(map[int]bool)
	worstPositions := make(map[int]int)
	var result []int

	// build the graph
	for _, playlist := range playlists {
		for j, track := range playlist {
			if j+1 < len(playlist) {
				graph[track] = append(graph[track], playlist[j+1])
			}
			visited[track] = false
			// find worst positions
			if val, exists := worstPositions[track]; exists && val < j {
				worstPositions[track] = j
			} else {
				worstPositions[track] = j
			}
		}
	}
	/*
	* find place to start dfs
	* starting point will be the key that has the worst position of 0
	 */
	var startNum int
	for key, val := range worstPositions {
		if val == 0 {
			startNum = key
		}
	}
	fmt.Println("DFS(", startNum, "):", graph)
	dfs(graph, startNum, &visited, &result)
	fmt.Println(result)
}

// Recursive Depth First Search
func dfs(g Graph, s int, vPtr *map[int]bool, result *[]int) {

	// dfs() for every neighbor of s
	(*vPtr)[s] = true
	for _, i := range g[s] {
		if !(*vPtr)[i] {
			dfs(g, i, vPtr, result)
		}
	}
	// append to results array in post-order
	*result = append([]int{s}, *result...)
}
