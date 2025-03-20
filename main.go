package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"

	"github.com/gitleaks/go-gitdiff/gitdiff"
)

func main() {
	// open original file
	originalFile, err := os.Open("a.txt")
	if err != nil {
		log.Fatal(err)
	}

	// read patch
	patchFile, err := os.Open("changes.patch")
	if err != nil {
		log.Fatal(err)
	}

	// parse patch
	patchChan, err := gitdiff.Parse(patchFile)
	if err != nil {
		log.Fatal(err)
	}

	// get patch from channel
	patch := <-patchChan

	// apply diff
	var output bytes.Buffer
	if err := gitdiff.Apply(&output, originalFile, patch); err != nil {
		log.Fatal(err)
	}

	// create new file
	modifiedFile, err := os.Create("b.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer modifiedFile.Close()

	// write to new file
	_, err = output.WriteTo(modifiedFile)
	if err != nil {
		log.Fatal(err)
	}

	// open new file
	openedModifiedFile, err := os.Open("b.txt")
	if err != nil {
		log.Fatal(err)
	}

	// read and print new file content
	scanner := bufio.NewScanner(openedModifiedFile)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

}
