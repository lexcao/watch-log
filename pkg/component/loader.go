package component

import "os"

type IteratorLine interface {
	HasNext() bool
	Next() string
}

type Loader interface {
	Load(file *os.File) IteratorLine
}
