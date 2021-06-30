package main

import (
	"bufio"
	"go-cmake-colors/gtestcoloring"
	"os"
	"log"
	"fmt"
)

func main() {
  scanner := bufio.NewScanner(os.Stdin)
	colorer := gtestcoloring.CreateGTestColorer()

  for scanner.Scan() {
		in := scanner.Text()
		fmt.Println(colorer.Colorize(in))
	}
	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
}
