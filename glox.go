package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		runfile()
	} else {
		runprompt()
	}
}

func runfile() {
	// TODO
}

func runprompt() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(" > ")
		line, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		s := newscanner(line)

		for _, t := range s.scantokens() {
			t.debug()
		}
	}
}
