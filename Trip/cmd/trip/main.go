package main

import "fmt"

// Given a list of source and destinations make a function that returns true if there is a route from a source and destination and the number
// connections to get from the source to destination.

// For example
// ["san franscico", "san diego"]
// ["san diego", "las vegas"]
// ["san diego", "los angeles"]
// ["las vegas", "salt lake city"]
// ["las vegas", "phoenix"]
// ["phoenix", "san antonio"]
// ["salt lake city", "denver"]
// ["denver", "indianapolis"]
// ["las vegas", "boise"]
// ["boise", "alberta"]

// Where the source = "san franscico" and the destination = "denver" should return
// ["san franscico", "san diego"]
// ["san diego", "las vegas"]
// ["las vegas", "salt lake city"]
// ["salt lake city", "denver"]
// with 4 connections

type Node struct {
	Visited     bool
	Destination string
}

type AdjacencyList struct {
	list map[string][]Node
}

func NewAdjacencyList() *AdjacencyList {
	l := make(map[string][]Node)
	return &AdjacencyList{
		list: l,
	}
}

func (a *AdjacencyList) AddEdge(source, destination string) {
	a.list[source] = append(a.list[source], Node{Destination: destination})
}

func (a *AdjacencyList) dfs(node Node, destination string) (bool, int) {
	if node.Visited {
		return false, 0
	}

	node.Visited = true
	fmt.Println(node.Destination)

	if node.Destination == destination {
		return true, 1
	}

	neighbors, _ := a.list[node.Destination]
	for _, n := range neighbors {
		found, count := a.dfs(n, destination)
		if found {
			return true, count + 1
		}
	}
	return false, 0
}

func (a *AdjacencyList) FindTrip(source, destination string) (found bool, numberConnections int) {
	fmt.Println(source)
	neighbors, _ := a.list[source]
	for _, n := range neighbors {
		found, count := a.dfs(n, destination)
		if found {
			return true, count
		}
	}
	return false, 0
}

func main() {

	al := NewAdjacencyList()

	al.AddEdge("san franscico", "san diego")
	al.AddEdge("san franscico", "las angeles")
	al.AddEdge("san diego", "las vegas")
	al.AddEdge("san diego", "los angeles")
	al.AddEdge("las vegas", "salt lake city")
	al.AddEdge("las vegas", "phoenix")
	al.AddEdge("phoenix", "san antonio")
	al.AddEdge("salt lake city", "denver")
	al.AddEdge("denver", "indianapolis")
	al.AddEdge("las vegas", "boise")
	al.AddEdge("boise", "alberta")

	found, count := al.FindTrip("san franscico", "indianapolis")
	fmt.Println(found)
	fmt.Println(count)
}
