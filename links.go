package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func findLinks(scanner bufio.Scanner) []string {
	var linkedRooms []string
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), "-") {
			rooms := strings.Split(scanner.Text(), "-")
			if rooms[0] == rooms[1] {
				fmt.Println("ERROR: invalid data format, The room cannot be linked to itself")
				os.Exit(0)
			}
			if len(rooms) == 2 {
				linkedRooms = append(linkedRooms, rooms...)
			}
		}
	}
	return linkedRooms
}
