package main

import (
	"bufio"
	"go-cmake-colors/gtestcoloring"
	"os"
	"log"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	colorer := gtestcoloring.CreateGTestColorer()
	
	defer writer.Flush()
	defer os.Stdout.Close()
	defer os.Stdin.Close()

	for scanner.Scan() {
		in := scanner.Text()
		colorer.Colorize(in, writer)
		if err := writer.Flush(); err != nil {
			log.Println(err)
		}
	}
	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
}

