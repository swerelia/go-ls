package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"./printcols"
)

func main() {
	showHidden := flag.Bool("a", false, "display hidden entries")
	directoryTrailingSlash := flag.Bool("F", false, "add '/' char at the end of each directory name")
	flag.Parse()

	arguments := flag.Args()
	path := arguments[0]

	entries, err := os.ReadDir(path)
	if err != nil {
		fmt.Println("Could not find directory '%s'", path)
		return
	}

	if !*showHidden {
		entries = filterHidden(entries)
	}

	names := getEntryNames(entries, *directoryTrailingSlash)

	// fmt.Printf("%v", names)
	printcols.PrintColumns(&names, 4);
}

func filterHidden(entries []os.DirEntry) []os.DirEntry {
	var filteredEntries []os.DirEntry
	for _, entry := range entries {
		if !strings.HasPrefix(entry.Name(), ".") {
			filteredEntries = append(filteredEntries, entry)
		}
	}
	return filteredEntries
}

func getEntryNames(entries []os.DirEntry, directoryTrailingSlash bool) []string {
	var names []string
	for _, entry := range entries {
		names = append(names, getEntryName(entry, directoryTrailingSlash))
	}
	return names
}

func getEntryName(entry os.DirEntry, directoryTrailingSlash bool) string {
	if entry.IsDir() && directoryTrailingSlash {
		return entry.Name() + "/"
	}
	return entry.Name()
}