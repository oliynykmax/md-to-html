package main

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

var (
	boldStart = []byte("<strong><b>")
	boldEnd   = []byte("</b></strong>")

	italicStart = []byte("<em><i>")
	italicEnd   = []byte("</i></em>")

	strikeStart = []byte("<del><s>")
	strikeEnd   = []byte("</s></del>")

	underlineStart = []byte("<u>")
	underlineEnd   = []byte("</u>")

	codeStart = []byte("<code>")
	codeEnd   = []byte("</code>")

	codeBlockStart = []byte("<pre><code>")
	codeBlockEnd   = []byte("</code></pre>")

	smallStart = []byte("<small>")
	smallEnd   = []byte("</small>")

	markStart = []byte("<mark>")
	markEnd   = []byte("</mark>")

	header = []byte(
		"<!DOCTYPE html>\n<html>\n<head>\n<meta charset=\"UTF-8\">\n<title>Document</title>\n</head>\n<body>\n",
	)
	footer = []byte("\n</body>\n</html>\n")
)

// utils
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

func processHeader(part []byte) []byte {
	var processed []byte

	n := 0
	for n < 3 && n < len(part) && part[n] == ' ' {
		n++
	}
	trimmed := part[n:]

	l := 0
	for l < len(trimmed) && l < 6 && trimmed[l] == '#' {
		l++
	}

	if l > 0 && len(trimmed) > l && trimmed[l] == ' ' {
		// It's a header, wrap in <h1>..<h6>
		processed = append(processed, []byte("<h"+strconv.Itoa(l)+">")...)
		processed = append(processed, trimmed[l+1:]...)
		processed = append(processed, []byte("</h"+strconv.Itoa(l)+">")...)
	} else {
		processed = append(processed, part...)
	}
	processed = append(processed, '\n')

	return processed
}

type TagPos struct {
	Index int
	Type  string
	Start bool
}

type FormatInfo struct {
	HtmlStart []byte
	HtmlEnd   []byte
	MdLen     int
}

var formatMap = map[string]FormatInfo{
	"bold":   {boldStart, boldEnd, 2},
	"italic": {italicStart, italicEnd, 1},
	"strike": {strikeStart, strikeEnd, 2},
}

func findMarkers(part []byte) ([]int, []int, []int) {
	var bold, italic, strike []int

	for i := 0; i < len(part); i++ {
		switch part[i] {
		case '*':
			if i+1 < len(part) && part[i+1] == '*' {
				bold = append(bold, i)
				i++
			} else {
				italic = append(italic, i)
			}
		case '_':
			if i+1 < len(part) && part[i+1] == '_' {
				bold = append(bold, i)
				i++
			} else {
				italic = append(italic, i)
			}
		case '~':
			if i+1 < len(part) && part[i+1] == '~' {
				strike = append(strike, i)
				i++
			}
		}
	}

	return bold, italic, strike
}

func createCleanTags(bold, italic, strike []int) []TagPos {
	if len(bold)%2 != 0 {
		bold = bold[:len(bold)-1]
	}
	if len(italic)%2 != 0 {
		italic = italic[:len(italic)-1]
	}
	if len(strike)%2 != 0 {
		strike = strike[:len(strike)-1]
	}

	var tags []TagPos
	addPairs := func(arr []int, t string) {
		for i := 0; i < len(arr); i += 2 {
			tags = append(tags, TagPos{Index: arr[i], Type: t, Start: true})
			tags = append(tags, TagPos{Index: arr[i+1], Type: t, Start: false})
		}
	}
	addPairs(bold, "bold")
	addPairs(italic, "italic")
	addPairs(strike, "strike")

	sort.Slice(tags, func(i, j int) bool {
		return tags[i].Index < tags[j].Index
	})

	// remove overlapping invalid pairs (like __Bold**)
	cleaned := []TagPos{}
	stack := make(map[string][]int)
	for _, t := range tags {
		if t.Start {
			stack[t.Type] = append(stack[t.Type], t.Index)
			cleaned = append(cleaned, t)
		} else {
			if len(stack[t.Type]) == 0 {
				continue
			}
			stack[t.Type] = stack[t.Type][:len(stack[t.Type])-1]
			cleaned = append(cleaned, t)
		}
	}

	return cleaned
}

func buildOutput(part []byte, cleaned []TagPos) []byte {
	processed := make([]byte, 0, len(part))

	last := 0

	for _, t := range cleaned {
		processed = append(processed, part[last:t.Index]...)
		var html []byte
		info := formatMap[t.Type]
		if t.Start {
			html = info.HtmlStart
		} else {
			html = info.HtmlEnd
		}
		processed = append(processed, html...)
		last = t.Index + info.MdLen
	}

	processed = append(processed, part[last:]...)

	return processed
}

func processInlineFormatting(part []byte) []byte {
	bold, italic, strike := findMarkers(part)

	cleaned := createCleanTags(bold, italic, strike)

	return buildOutput(part, cleaned)
}

func processCodeBlocks(in []byte) ([]byte, map[string][]byte) {
	lines := bytes.Split(in, []byte("\n"))
	var out []byte
	codeBlocks := make(map[string][]byte)
	counter := 0
	i := 0
	for i < len(lines) {
		line := lines[i]
		if bytes.HasPrefix(line, []byte("```")) {
			lang, _ := bytes.CutPrefix(line, []byte("```"))
			lang = bytes.TrimSpace(lang)
			var codeContent []byte
			i++
			for i < len(lines) && !bytes.HasPrefix(lines[i], []byte("```")) {
				codeContent = append(codeContent, lines[i]...)
				codeContent = append(codeContent, '\n')
				i++
			}
			if i < len(lines) {
				i++ // skip closing ```
			}
			placeholder := fmt.Sprintf("__CODEBLOCK_%d__", counter)
			counter++
			out = append(out, []byte(placeholder)...)
			out = append(out, '\n')
			var html []byte
			html = append(html, codeBlockStart...)
			if len(lang) > 0 {
				html = append(html, []byte(` class="language-`+string(lang)+`"`)...)
			}
			html = append(html, '>')
			html = append(html, codeContent...)
			html = append(html, codeBlockEnd...)
			html = append(html, '\n')
			codeBlocks[placeholder] = html
		} else {
			out = append(out, line...)
			out = append(out, '\n')
			i++
		}
	}
	return out, codeBlocks
}

func convertToHtml(in []byte) (out []byte) {

	out = append(out, header...)

	in, codeBlocks := processCodeBlocks(in)

	// loop to process each string and the apply the formats in html way
	for part := range bytes.SplitSeq(in, []byte("\n")) {
		processed := processInlineFormatting(part)
		processed = processHeader(processed)
		out = append(out, processed...)
	}

	// replace placeholders with code block HTML
	for placeholder, html := range codeBlocks {
		out = bytes.ReplaceAll(out, []byte(placeholder), html)
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

	output := convertToHtml(d)

	_, err = f.Write(output)
	check(err)

	fmt.Println(string(output))
}
