package main

import (
	"bufio"
	"flag"
	"log"
	"os"

	"github.com/sh0e1/gotree"
)

func main() {
	var (
		a = flag.Bool("a", false, "All files are listed.")
		l = flag.Int("L", -1, "Descend only level directories deep.")
	)

	flag.Parse()
	opt := &gotree.Option{
		IsDisplayAllFiles: *a,
		Level:             *l,
	}
	w := bufio.NewWriter(os.Stdout)
	if err := gotree.Execute(w, flag.Args(), opt); err != nil {
		log.Fatal(err)
	}
	w.Flush()
}
