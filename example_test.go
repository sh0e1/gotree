package gotree_test

import (
	"bufio"
	"os"

	"github.com/sh0e1/gotree"
)

func ExampleExecute() {
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	_ = gotree.Execute(w, []string{"./testdata"}, nil)
	// Output:
	// ./testdata
	// ├── file
	// └── sub1
	//    ├── file
	//    └── sub2
	//       └── file
}

func ExampleExecute_displayAll() {
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	opt := &gotree.Option{
		IsDisplayAllFiles: true,
		// Output all directories and files recursively with level -1
		Level: -1,
	}
	_ = gotree.Execute(w, []string{"./testdata"}, opt)
	// Output:
	// ./testdata
	// ├── .invisible
	// ├── file
	// └── sub1
	//    ├── file
	//    └── sub2
	//       └── file
}
func ExampleExecute_setLevel() {
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	opt := &gotree.Option{
		Level: 1,
	}
	_ = gotree.Execute(w, []string{"./testdata"}, opt)
	// Output:
	// ./testdata
	// ├── file
	// └── sub1
}
