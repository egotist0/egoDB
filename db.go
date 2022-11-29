package egoDB

import (
	"errors"
	"os"
	"sync"

	"github.com/egotist0/egoDB/util"
)

var (
	// ErrDefaultCfNil default column family is nil.
	ErrDefaultCfNil = errors.New("default column family is nil")
)

// egoDB provide basic operations for a persistent kv store.
// It`s methods(Put Get Delete) are self explanatory, and executed in default ColumnFamily.
// You can create a custom ColumnFamily by calling method OpenColumnFamily.
type egoDB struct {
	// all column families.
	cfs  map[string]*ColumnFamily
	opts Options
	mu   sync.RWMutex
}

// Open a new egoDB instance, actually will just open the default column family.
func Open(opt Options) (*egoDB, error) {
	if !util.PathExist(opt.DBPath) {
		if err := os.MkdirAll(opt.DBPath, os.ModePerm); err != nil {
			return nil, err
		}
	}

	db := &egoDB{opts: opt, cfs: make(map[string]*ColumnFamily)}
	// load default column family.
	if opt.CfOpts.CfName == "" {
		opt.CfOpts.CfName = DefaultColumnFamilyName
	}
	if _, err := db.OpenColumnFamily(opt.CfOpts); err != nil {
		return nil, err
	}
	return db, nil
}

// Close close database.
func (db *egoDB) Close() error {
	for _, cf := range db.cfs {
		if err := cf.Close(); err != nil {
			return err
		}
	}
	return nil
}

// Sync syncs the content of all column families to disk.
func (db *egoDB) Sync() error {
	for _, cf := range db.cfs {
		if err := cf.Sync(); err != nil {
			return err
		}
	}
	return nil
}

// Put put to default column family.
func (db *egoDB) Put(key, value []byte) error {
	return db.PutWithOptions(key, value, nil)
}

// PutWithOptions put to default column family with options.
func (db *egoDB) PutWithOptions(key, value []byte, opt *WriteOptions) error {
	columnFamily := db.getColumnFamily(DefaultColumnFamilyName)
	if columnFamily == nil {
		return ErrDefaultCfNil
	}
	return columnFamily.PutWithOptions(key, value, opt)
}

// Get get from default column family.
func (db *egoDB) Get(key []byte) ([]byte, error) {
	columnFamily := db.getColumnFamily(DefaultColumnFamilyName)
	if columnFamily == nil {
		return nil, ErrDefaultCfNil
	}
	return columnFamily.Get(key)
}

// Delete delete from default column family.
func (db *egoDB) Delete(key []byte) error {
	return db.DeleteWithOptions(key, nil)
}

// DeleteWithOptions delete from default column family with options.
func (db *egoDB) DeleteWithOptions(key []byte, opt *WriteOptions) error {
	columnFamily := db.getColumnFamily(DefaultColumnFamilyName)
	if columnFamily == nil {
		return ErrDefaultCfNil
	}
	return columnFamily.DeleteWithOptions(key, opt)
}

func (db *egoDB) getColumnFamily(cfName string) *ColumnFamily {
	db.mu.RLock()
	defer db.mu.RUnlock()
	return db.cfs[cfName]
}
