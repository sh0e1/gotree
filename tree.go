package gotree

import (
	"fmt"
	"io"
	"io/ioutil"
	"strings"
)

// Option ...
type Option struct {
	IsDisplayAllFiles bool
	Level             int
}

// Tree ...
func Tree(w io.Writer, dir, stem string, level int, opt *Option) error {
	if opt.Level == level {
		return nil
	}
	level++

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return err
	}

	for i, f := range files {
		filename := f.Name()
		if !opt.IsDisplayAllFiles && strings.HasPrefix(filename, ".") {
			continue
		}

		var (
			branch   = "├──"
			addition = "│  "
		)
		if len(files)-1 == i {
			branch = "└──"
			addition = "   "
		}

		if f.IsDir() {
			fmt.Fprintf(w, "%s%s \x1b[34m%s\x1b[0m\n", stem, branch, filename)
			dirname := fmt.Sprintf("%s/%s", dir, filename)
			if err := Tree(w, dirname, stem+addition, level, opt); err != nil {
				return err
			}
		} else {
			fmt.Fprintf(w, "%s%s %s\n", stem, branch, filename)
		}
	}
	return nil
}
