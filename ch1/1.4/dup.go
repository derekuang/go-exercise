package main

import (
	"bufio"
	"fmt"
	"os"
)

type lineStat struct {
	count int
	files []string
}

func main() {
	counts := make(map[string]lineStat)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, stat := range counts {
		if stat.count > 1 {
			fmt.Printf("%d\t%s\t", stat.count, line)
			for _, file := range stat.files {
				fmt.Printf("%s\t", file)
			}
			fmt.Println()
		}
	}
}

func countLines(f *os.File, counts map[string]lineStat) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		if entry, ok := counts[input.Text()]; ok {
			entry.count++
			if !isMember(entry.files, f.Name()) {
				entry.files = append(entry.files, f.Name())
			}
			counts[input.Text()] = entry
		} else {
			counts[input.Text()] = lineStat{
				count: 1,
				files: []string{f.Name()},
			}
		}
	}
	// NOTE: ignoring potential errors from input.Err()
}

func isMember(arr []string, elem string) bool {
	for _, e := range arr {
		if e == elem {
			return true
		}
	}
	return false
}
