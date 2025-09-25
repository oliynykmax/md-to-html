package main

import (
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		fmt.Fprintln(os.Stderr, "Error:", e)
		fatalUsage()
	}
}

func fatalUsage() {
	fmt.Println("Usage: ./md-to-html file.md")
	os.Exit(1)
}

func main() {

	if len(os.Args) != 2 || !strings.HasSuffix(os.Args[1], ".md") {
		fatalUsage()
	}

	dat, err := os.ReadFile(os.Args[1])
	check(err)

	fmt.Printf("File len:%d", len(dat))
}
