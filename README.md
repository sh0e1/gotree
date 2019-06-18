# gotree

tree command by Go

## Requirement

- [Go 1.12](https://golang.org/doc/go1.12)

## Install

```bash
go get github.com/sh0e1/gotree
go install $GOPATH/src/github.com/sh0e1/gotree/cmd/gotree
```

## Usage

```bash
gotree -h
Usage of gotree:
  -L int
        Descend only level directories deep. (default -1)
  -a    All files are listed.
```

```bash
gotree
.
├── LICENSE
├── Makefile
├── README.md
├── cmd
│  └── gotree
│     └── main.go
├── go.mod
├── testdata
│  ├── file
│  └── sub1
│     ├── file
│     └── sub2
│        └── file
└── tree.go

```

## Licence

[MIT](https://github.com/sh0e1/gotree/blob/master/LICENSE)

## Author

[sh0e1](https://github.com/sh0e1)
