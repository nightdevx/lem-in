package main

import (
	"bufio"
	"fmt"
	"os"
)

// Dosyayı açmak için
func fileOpener() *os.File {
	openFile, err := os.Open("examples/" + os.Args[1])
	if err != nil {
		fmt.Println("Dosya açılamadı!")
		os.Exit(1)
	}
	return openFile
}

// Dosyayı taramak için
func fileScanner(dosya *os.File) *bufio.Scanner {
	if dosya != nil {
		return bufio.NewScanner(dosya)
	}
	return nil
}
