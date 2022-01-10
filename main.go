package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	showHidden := flag.Bool("a", false, "display hidden files")
	flag.Parse()

	arguments := flag.Args()
	path := arguments[0]

	files, err := getDirectoryContents(path, *showHidden)
	if err != nil {
		println("No such file or directory")
		return
	}
	fmt.Printf("%v", files)
}

func getDirectoryContents(path string, showHidden bool) ([]string, error) {
	entries, err := os.ReadDir(path)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	var files []string
	for _, entry := range entries {
		if showHidden == true || !strings.HasPrefix(entry.Name(), ".") {
			files = append(files, entry.Name())
		}
	}
	return files, nil
}
