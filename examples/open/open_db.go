package main

import (
	"github.com/egotist0/egoDB"
	"io/ioutil"
	"os"
)

func main() {
	// open a db with default options.
	path, _ := ioutil.TempDir("", "egoDB")
	// you must specify a db path.
	opts := egoDB.DefaultOptions(path)
	db, err := egoDB.Open(opts)
	defer func() {
		_ = db.Close()
		_ = os.RemoveAll(path)
	}()

	if err != nil {
		panic(err)
	}
}
