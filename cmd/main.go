package main

import (
	"bytes"
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

	var processed []byte
	// loop to process each string and the apply the formats in html way
	for part := range bytes.SplitSeq(in, []byte("\n")) {
		trimmed := bytes.TrimSpace(part)

		l := 0
		for ; l < len(trimmed) && l < 6; l++ {
			if trimmed[l] != '#' {
				break
			}
		}

		if l > 0 && len(trimmed) > l && trimmed[l] == ' ' {
			// It's a header, wrap in <h1>..<h6>
			processed = append(processed, []byte("<h"+string('0'+l)+">")...)
			processed = append(processed, trimmed[l+1:]...)
			processed = append(processed, []byte("</h"+string('0'+l)+">")...)
		} else {
			processed = append(processed, part...)
		}
		processed = append(processed, '\n')

		out = append(out, processed...)
	}

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
