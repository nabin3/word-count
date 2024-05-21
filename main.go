package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	// Defining options
	countBytes := flag.Bool("c", false, "count bytes in a file")
	countLines := flag.Bool("l", false, "count lines in a file")
	countWords := flag.Bool("w", false, "count words in a file")
	countChars := flag.Bool("m", false, "count charachters in a file")

	// Custom Usage function
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "usage: %s [option] [filename]\n", os.Args[0])
		flag.PrintDefaults()
	}
	// Parses all flags
	flag.Parse()

	// Checking if user has given more than one non-flag argument
	if flag.NArg() != 1 {
		flag.Usage()
		os.Exit(1)
	}
	// Getting the filename
	filename := flag.Arg(0)

	// Reading from given file
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	if *countBytes {
		log.Printf("%d %s\n", len(data), filename)
	} else if *countLines {
		log.Printf("%d %s\n", len(strings.Split(string(data), "\n")), filename)
	} else if *countWords {
		log.Printf("%d %s\n", len(strings.Fields(string(data))), filename)
	} else if *countChars {
		log.Printf("%d %s\n", len([]rune(string(data))), filename)
	} else {
		fmt.Printf("%d  %d  %d %d %s\n", len(data), len(strings.Split(string(data), "\n")), len(strings.Fields(string(data))), len([]rune(string(data))), filename)
		os.Exit(0)
	}
}
