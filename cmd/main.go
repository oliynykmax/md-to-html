package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func check(e error) {
	if e != nil {
		fmt.Fprintln(os.Stderr, "Error:", e)
		os.Exit(1)
	}
}

func fatalUsage() {
	fmt.Println("Usage: ./md-to-html file.md")
	os.Exit(1)
}

func convertToHtml(in []byte) (out []byte, err error) {
	err = nil
	header := []byte(
		"<!DOCTYPE html>\n<html>\n<head>\n<meta charset=\"UTF-8\">\n<title>Document</title>\n</head>\n<body>\n",
	)
	footer := []byte("\n</body>\n</html>\n")

	out = append(out, header...)
	out = append(out, in...)
	out = append(out, footer...)

	return
}

func main() {

	if len(os.Args) != 2 || !strings.HasSuffix(os.Args[1], ".md") {
		fatalUsage()
	}

	d, err := os.ReadFile(os.Args[1])
	check(err)

	base := filepath.Base(os.Args[1])
	f, err := os.Create(strings.Replace(base, ".md", ".html", 1))
	check(err)
	defer f.Close()

	output, err := convertToHtml(d)
	check(err)

	_, err = f.Write(output)
	check(err)

	fmt.Println(string(output))
}
