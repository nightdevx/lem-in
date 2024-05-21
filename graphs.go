package main

import (
	"container/list"
	"fmt"
	"sort"
)

type Node struct {
	Value     interface{}
	Neighbors []*Node
	ID        int
}

type Graph struct {
	Nodes []*Node
}

var startNode, endNode *Node

// NewGraph fonksiyonu ile yeni bir Graph oluşturulur.
func NewGraph() *Graph {
	return &Graph{Nodes: []*Node{}} // Boş bir düğüm listesi ile başlatılır
}

// Graph'a node ekler.
func (g *Graph) AddNode(value interface{}) *Node {
	node := &Node{Value: value}
	g.Nodes = append(g.Nodes, node)
	return node
}

// Graph'taki iki node'u birbirine bağlar.
func (g *Graph) AddEdge(node1, node2 *Node) {
	node1.Neighbors = append(node1.Neighbors, node2)
	node2.Neighbors = append(node2.Neighbors, node1)
}

// Graph oluşturur.
func createGraph(rooms, links []string, startRoom, endRoom string) *Graph {
	graph := NewGraph()

	// Rooms (odalar) eklenir
	roomNodes := make(map[string]*Node) // Oda isimlerini düğümlere eşlemek için bir map
	for _, room := range rooms {
		if room == startRoom {
			startNode = graph.AddNode(startRoom)
			roomNodes[room] = startNode
		} else if room == endRoom {
			endNode = graph.AddNode(endRoom)
			roomNodes[room] = endNode
		} else {
			node := graph.AddNode(room)
			roomNodes[room] = node
		}
	}

	// Graph içerisindeki bağlantıları oluşturur
	for i := 0; i < len(links); i += 2 {
		from, to := links[i], links[i+1]

		// Oda isimlerinin geçerli olup olmadığı kontrol edilir
		fromNode, fromExists := roomNodes[from]
		toNode, toExists := roomNodes[to]

		if !fromExists || !toExists {
			panic("Invalid link: room not found") // Geçersiz oda ismi
		}
		if fromNode == toNode {
			panic("Invalid link:room cannot be related to itself") // Kendisi ile bağlantılı oda hatası
		}

		graph.AddEdge(fromNode, toNode)
	}

	return graph
}

// Graph ı yazdırmak için kullanılır
func printGraph(graph *Graph) {
	for _, node := range graph.Nodes {
		fmt.Printf("Node: %v, Neighbors: ", node.Value)
		for _, neighbor := range node.Neighbors {
			fmt.Printf("%v ", neighbor.Value)
		}
		fmt.Println()
	}
}

// Başlangıçtan sona giden tüm yolları bulmak için BFS kullanılır.
func findAllPathsBFS(start, end *Node) [][]interface{} {
	queue := list.New()
	queue.PushBack([]*Node{start})
	var result [][]interface{}

	for queue.Len() > 0 {
		path := queue.Remove(queue.Front()).([]*Node)
		node := path[len(path)-1]

		if node == end {
			var newPath []interface{}
			for _, n := range path {
				newPath = append(newPath, n.Value)
			}
			result = append(result, newPath)
			continue
		}

		for _, neighbor := range node.Neighbors {
			if !contains(path, neighbor) {
				newPath := make([]*Node, len(path))
				copy(newPath, path)
				newPath = append(newPath, neighbor)
				queue.PushBack(newPath)
			}
		}
	}
	return result
}

// Yolun verilen node u içerip içermediğini kontrol eder
func contains(path []*Node, node *Node) bool {
	for _, n := range path {
		if n == node {
			return true
		}
	}
	return false
}

// En fazla sayıda birbiriyle çakışmayan dizi eleman sayısı en fazla ve altdizi eleman sayısı en az olan kombinasyonu bulur
func findMaxNonOverlappingPaths(start, end *Node) [][]interface{} {
	allPaths := findAllPathsBFS(start, end)
	sort.Slice(allPaths, func(i, j int) bool {
		return len(allPaths[i]) < len(allPaths[j])
	})

	return maxNonOverlappingPaths(allPaths)
}

func maxNonOverlappingPaths(paths [][]interface{}) [][]interface{} {
	var maxSet [][]interface{}
	for _, path := range paths {
		if !overlapsWithExistingSet(path, maxSet) {
			maxSet = append(maxSet, path)
			if len(maxSet) == len(paths) {
				return maxSet // Early termination
			}
		}
	}
	return maxSet
}

func overlapsWithExistingSet(path []interface{}, existingSet [][]interface{}) bool {
	for _, existingPath := range existingSet {
		for i := 1; i < len(path)-1; i++ {
			for j := 1; j < len(existingPath)-1; j++ {
				if path[i] == existingPath[j] {
					return true
				}
			}
		}
	}
	return false
}
