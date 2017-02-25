package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/pkg/errors"
)

func main() {
	yamlInputFile := flag.String("yaml", "irssi.yaml", "Input Yaml File")
	flag.Parse()

	configMapping, err := ParseInputFile(*yamlInputFile)
	if err != nil {
		fmt.Print(errors.Wrap(err, "could not parse input file"))
		os.Exit(1)
	}

	err = WriteIrssFile(*configMapping)
	if err != nil {
		fmt.Print(errors.Wrap(err, "could not write config file"))
		os.Exit(1)
	}
}
