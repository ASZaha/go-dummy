package main

import (
	"bytes"
	"fmt"
	"log"
	"os"

	"github.com/gitleaks/go-gitdiff/gitdiff"
)

func main() {
	patch, err := os.Open("changes.patch")
	if err != nil {
		log.Fatal(err)
	}

	// files is a slice of *gitdiff.File describing the files changed in the patch
	// preamble is a string of the content of the patch before the first file
	files, err := gitdiff.Parse(patch)
	if err != nil {
		log.Fatal(err)
	}

	code, err := os.Open("a.txt")
	if err != nil {
		log.Fatal(err)
	}
	file := <-files

	// apply the changes in the patch to a source file
	var output bytes.Buffer
	if err := gitdiff.Apply(&output, code, file); err != nil {
		log.Fatal(err)
	}

	fmt.Println(output.String())
}
