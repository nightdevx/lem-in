package main

import (
	"bufio"
	"strings"
)

var startRoom, endRoom string

func findAllRooms(scanner bufio.Scanner) []string {
	var rooms []string
	for scanner.Scan() {
		line := scanner.Text()
		if line == "##start" {
			scanner.Scan()
			line = strings.Split(scanner.Text(), " ")[0]
			startRoom = line
			rooms = append(rooms, line)
		} else if line == "##end" {
			scanner.Scan()
			line = strings.Split(scanner.Text(), " ")[0]
			endRoom = line
			rooms = append(rooms, line)
		} else if strings.Contains(line, " ") {
			splitLine := strings.Split(line, " ")
			if len(splitLine) == 3 {
				line = strings.Split(scanner.Text(), " ")[0]
				rooms = append(rooms, line)
			}
		} 
	}
	return rooms
}
