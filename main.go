package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
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

	opt := &option{
		isDisplayAllFiles: *a,
		level: *l,
	}
	w := bufio.NewWriter(os.Stdout)
	for _, dir := range dirnames {
		fmt.Fprintf(w, "\x1b[34m%s\x1b[0m\n", dir)
		tree(w, dir,"", 0, opt)
	}
	w.Flush()
}

type option struct {
	isDisplayAllFiles bool
	level int
}

func tree(w io.Writer, dir, stem string, level int, opt *option) {
	if opt.level == level {
		return
	}
	level++

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatalln(err)
	}

	for i, f := range files {
		filename := f.Name()
		if !opt.isDisplayAllFiles && strings.HasPrefix(filename, ".") {
			continue
		}

		var (
			branch = "├──"
			addition = "│  "
		)
		if len(files)-1 == i {
			branch = "└──"
			addition = "   "
		}

		if f.IsDir() {
			fmt.Fprintf(w, "%s%s \x1b[34m%s\x1b[0m\n", stem, branch, filename)
			dirname := fmt.Sprintf("%s/%s", dir, filename)
			tree(w, dirname, stem+addition, level, opt)
		} else {
			fmt.Fprintf(w, "%s%s %s\n", stem, branch, filename)
		}
	}
}
