package main

import (
	"fmt"
	"os"

	"github.com/hexops/gotextdiff"
	"github.com/hexops/gotextdiff/myers"
	"github.com/hexops/gotextdiff/span"
)

func main() { // Read the contents of the two files.
	file1Path := "a.txt"
	file2Path := "b.txt"

	file1Bytes, err := os.ReadFile(file1Path)
	if err != nil {
		fmt.Printf("Error reading file1: %v\n", err)
		return
	}
	file2Bytes, err := os.ReadFile(file2Path)
	if err != nil {
		fmt.Printf("Error reading file2: %v\n", err)
		return
	}

	file1 := string(file1Bytes)
	file2 := string(file2Bytes)

	// Generate the diff.
	diff := myers.ComputeEdits(span.URIFromPath(file1Path), file1, file2)
	fmt.Print(diff)

	// Print the diff.
	fmt.Println(fmt.Sprint(gotextdiff.ToUnified(file1Path, file2Path, file1, diff)))

	// Apply the diff to file1.
	applied := gotextdiff.ApplyEdits(file1, diff)

	// Print the result.
	fmt.Println(applied)
}
