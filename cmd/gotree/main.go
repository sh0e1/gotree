package main

import (
	"bufio"
	"flag"
	"fmt"
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
	dirnames := flag.Args()
	if len(dirnames) == 0 {
		dirnames = []string{"."}
	}

	opt := &gotree.Option{
		IsDisplayAllFiles: *a,
		Level:             *l,
	}
	w := bufio.NewWriter(os.Stdout)
	for _, dir := range dirnames {
		fmt.Fprintf(w, "\x1b[34m%s\x1b[0m\n", dir)
		if err := gotree.Tree(w, dir, "", 0, opt); err != nil {
			log.Fatal(err)
		}
	}
	w.Flush()
}
