package main

import (
	"github.com/egotist0/egoDB"
	"io/ioutil"
	"time"
)

// basic operations for egoDB:
// put
// put with options
// get
// delete
// delete with options
func main() {
	path, _ := ioutil.TempDir("", "egoDB")
	opts := egoDB.DefaultOptions(path)
	db, err := egoDB.Open(opts)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// 1.----put----
	key1 := []byte("name")
	err = db.Put(key1, []byte("egoDB"))
	if err != nil {
		// ...
	}

	key2 := []byte("feature")
	// 2.----put with options----
	writeOpts := &egoDB.WriteOptions{
		Sync:      true,
		ExpiredAt: time.Now().Add(time.Second * 100).Unix(),
	}
	err = db.PutWithOptions(key2, []byte("store data"), writeOpts)
	if err != nil {
		// ...
	}

	// 3.----get----
	val, err := db.Get(key1)
	if err != nil {
		// ...
	}
	if len(val) > 0 {
		// ...
	}

	// 4.----delete----
	err = db.Delete(key1)
	if err != nil {
		// ...
	}

	// 5.----delete with options----
	deleteOpts := &egoDB.WriteOptions{
		Sync:       false,
		DisableWal: true,
	}
	err = db.DeleteWithOptions([]byte("dummy key"), deleteOpts)
	if err != nil {
		// ...
	}
}
