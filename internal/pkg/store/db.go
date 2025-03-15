package store

import "sync"

type DB struct {
	records map[string]string
	mu      sync.Mutex
}

func NewDB() *DB {
	return &DB{
		records: make(map[string]string),
	}
}

func (d *DB) Get(key string) (string, bool) {
	d.mu.Lock()
	defer d.mu.Unlock()

	v, ok := d.records[key]
	return v, ok
}

func (d *DB) Set(key, value string) {
	d.mu.Lock()
	defer d.mu.Unlock()

	d.records[key] = value
}
