package main

import (
	"github.com/egotist0/egoDB"
	"io/ioutil"
)

func main() {
	// open a db with default options.
	path, _ := ioutil.TempDir("", "egoDB")
	// you must specify a db path.
	opts := egoDB.DefaultOptions(path)
	db, err := egoDB.Open(opts)
	defer func() {
		_ = db.Close()
	}()
	if err != nil {
		panic(err)
	}

	cfOpts := egoDB.DefaultColumnFamilyOptions("a-new-cf")
	cfOpts.DirPath = "/tmp"
	cf, err := db.OpenColumnFamily(cfOpts)
	if err != nil {
		panic(err)
	}

	// the same with db.Put
	err = cf.Put([]byte("name"), []byte("egoDB"))
	if err != nil {
		// ...
	}
}
