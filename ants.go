package main

import (
	"bufio"
	"fmt"
	"strconv"
)

type AntSimulation struct {
	antCount             int
	end                  *Node
	allPaths             [][]*Node
	bestPaths            [][]interface{}
	antPaths             [][]*Node
	antPositions         []int
	isNodeFull           map[*Node]bool
	antSteps             []int
	finishedAnts         []bool
	tour                 int
	startNodeConnections int
}

func NewAntSimulation(antCount int, end *Node, allPaths [][]*Node, bestChoosedPaths [][]interface{}) *AntSimulation {
	sim := &AntSimulation{
		antCount:     antCount,
		end:          end,
		allPaths:     allPaths,
		bestPaths:    bestChoosedPaths,
		antPaths:     make([][]*Node, antCount),
		antPositions: make([]int, antCount),
		isNodeFull:   make(map[*Node]bool),
		antSteps:     make([]int, antCount),
		finishedAnts: make([]bool, antCount),
		tour:         1,
	}

	findPathNode := interfaceToNode(bestChoosedPaths)
	for i := 0; i < antCount; i++ {
		sim.antPaths[i] = findPathNode[i%len(findPathNode)]
		sim.antPositions[i] = 1
		sim.antSteps[i] = 1
	}
	sim.antPaths[antCount-1] = interfaceToNode(bestChoosedPaths)[0]
	sim.startNodeConnections = len(findPathNode)
	
	return sim
}

func (sim *AntSimulation) Simulate() {
	if len(sim.allPaths) == 0 {
		fmt.Println("Başlangıç ve bitiş noktası arasında yol bulunamadı.")
		return
	}

	for {
		allAntsFinished := true
		maxPathLength := 0
		antsMovingFromStart := 0

		for i := 0; i < sim.antCount; i++ {
			if sim.antPositions[i] >= len(sim.antPaths[i]) || sim.finishedAnts[i] {
				continue
			}

			if len(sim.antPaths[i]) > maxPathLength {
				maxPathLength = len(sim.antPaths[i])
			}

			if sim.antSteps[i] < maxPathLength {
				nextNode := sim.antPaths[i][sim.antPositions[i]]

				if sim.antPositions[i] > 1 && sim.antPositions[i]-1 < len(sim.antPaths[i]) {
					sim.isNodeFull[sim.antPaths[i][sim.antPositions[i]-1]] = false
				}

				if sim.antPositions[i] == 1 {
					if antsMovingFromStart >= sim.startNodeConnections {
						continue
					}
					antsMovingFromStart++
				}

				if !sim.isNodeFull[nextNode] || nextNode == sim.antPaths[i][len(sim.antPaths[i])-1] {
					fmt.Printf("L%d-%s ", i+1, nextNode.Value)
					sim.isNodeFull[nextNode] = true
					sim.antPositions[i]++
					sim.antSteps[i]++

					if nextNode == sim.end {
						sim.finishedAnts[i] = true
					}
				}

			}

			if sim.antPositions[i] < len(sim.antPaths[i]) {
				allAntsFinished = false
			}
		}

		fmt.Println()

		if allAntsFinished {
			break
		}
		sim.tour++
	}
	fmt.Printf("Tour : %d\n", sim.tour)
}

func interfaceToNode(paths [][]interface{}) [][]*Node {
	result := make([][]*Node, len(paths))

	for i, path := range paths {
		result[i] = make([]*Node, len(path))
		for j, value := range path {
			result[i][j] = &Node{Value: value}
		}

		// Komşu düğümleri bağla (ilk ve son düğüm hariç)
		for j := 1; j < len(path)-1; j++ {
			result[i][j].Neighbors = []*Node{result[i][j-1], result[i][j+1]}
		}

		// İlk düğümün sadece bir komşusu var (sonraki düğüm)
		result[i][0].Neighbors = []*Node{result[i][1]}

		// Son düğümün sadece bir komşusu var (önceki düğüm)
		result[i][len(path)-1].Neighbors = []*Node{result[i][len(path)-2]}
	}

	return result
}

func findAntCount(scanner bufio.Scanner) int {
	scanner.Scan()
	antCount, _ := strconv.Atoi(scanner.Text())
	return antCount
}
