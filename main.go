package main

import (
	"fmt"
	"os"
	"time"
)

var (
	antCount           int
	rooms, linkedRooms []string
)

func main() {
	// Zaman ölçümü başlat
	startTime := time.Now()
	fileProcessing()

	graph := createGraph(rooms, linkedRooms, startRoom, endRoom)
	if startNode == nil {
		fmt.Println("ERROR: invalid data format, no start room found")
		os.Exit(0)
	}
	paths := findAllPathsBFS(startNode, endNode)
	bestCombinatedPath := findMaxNonOverlappingPaths(startNode, endNode)
	intToNodePaths := interfaceToNode(paths)
	simulation := NewAntSimulation(antCount, endNode, intToNodePaths, bestCombinatedPath)
	simulation.Simulate()

	fmt.Println("----------------------------")
	fmt.Println("BAŞLANGIÇ VE BİTİŞ ODALARI: "+startRoom, endRoom)
	fmt.Print("TÜM ODALAR: ")
	fmt.Println(rooms)
	fmt.Println("BAĞLANTILI ODALAR: ")
	for i := 0; i < len(linkedRooms); i += 2 {
		fmt.Printf("%s - %s\n", linkedRooms[i], linkedRooms[i+1])
	}
	fmt.Print("KARINCA SAYISI: ")
	fmt.Println(antCount)
	fmt.Println("----------------------------")
	fmt.Println("GRAF YAPISI:")
	printGraph(graph)
	fmt.Println("----------------------------")
	fmt.Println("OPTİMUM SEÇİLEN YOLLAR:")
	for _, cpath := range bestCombinatedPath {
		fmt.Println(cpath)
	}
	fmt.Println("----------------------------")

	// Zaman ölçümü bitir
	elapsed := time.Since(startTime)
	fmt.Printf("Kodun çalışması %.8f saniye sürdü.\n", elapsed.Seconds())
}

func fileProcessing() {
	file := fileOpener()
	scanner := fileScanner(file)
	antCount = findAntCount(*scanner)
	if antCount <= 0 {
		fmt.Println("ERROR: invalid data format, invalid number of Ants")
		os.Exit(0)
	}
	file.Seek(0, 0)
	rooms = findAllRooms(*scanner)

	file.Seek(0, 0)
	linkedRooms = findLinks(*scanner)
}
