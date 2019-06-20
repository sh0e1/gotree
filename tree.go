package gotree

import (
	"fmt"
	"io"
	"io/ioutil"
	"strings"
)

// Execute ...
func Execute(w io.Writer, dirs []string, opt *Option) error {
	if len(dirs) == 0 {
		dirs = []string{"."}
	}

	if opt == nil {
		opt = defaultOption
	}

	for _, dir := range dirs {
		fmt.Fprintf(w, "%s\n", dir)
		if err := tree(w, dir, "", 0, opt); err != nil {
			return err
		}
	}
	return nil
}

func tree(w io.Writer, dir, stem string, level int, opt *Option) error {
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
		fmt.Fprintf(w, "%s%s %s\n", stem, branch, filename)
		if f.IsDir() {
			dirname := fmt.Sprintf("%s/%s", dir, filename)
			if err := tree(w, dirname, stem+addition, level, opt); err != nil {
				return err
			}
		}
	}
	return nil
}

// Option ...
type Option struct {
	IsDisplayAllFiles bool
	Level             int
}

var defaultOption = &Option{
	IsDisplayAllFiles: false,
	Level:             -1,
}
