package main

import (
	"fmt"

	"github.com/hexops/gotextdiff"
	"github.com/hexops/gotextdiff/myers"
	"github.com/hexops/gotextdiff/span"
)

func main() {

	edits := myers.ComputeEdits(span.URIFromPath("a.txt"), aString, bString)
	diff := fmt.Sprint(gotextdiff.ToUnified("a.txt", "b.txt", aString, edits))

}
