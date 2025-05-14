package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/ilbagatto/vsop87-go/internal/vsop87"
)

const modulePath = "github.com/ilbagatto/vsop87-go"

func main() {
	in := flag.String("in", "data/vsop87d.yaml", "path to vsop87 YAML input")
	out := flag.String("out", "internal/vsop87/generated", "output directory for generated code")
	flag.Parse()

	if err := vsop87.GenerateSources(modulePath, *in, *out); err != nil {
		fmt.Fprintf(os.Stderr, "Error generating sources: %v", err)
		os.Exit(1)
	}
	fmt.Println("VSOP87 sources generated in ", *out)
}
