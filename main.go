package main // import "markspin"

//go:generate exec/build-resources

import (
	"markspin/view"
)

func main() {
	view.Main()
}
