package redis

import (
	"fmt"

	"github.com/mwmahlberg/bleve/index/store"
)

// Reader implements the KVReader interface for a Redis backend.
type Reader struct {
	store *Store
}

// Get returns the value associated with the key
// If the key does not exist, nil is returned.
// The caller owns the bytes returned.
func (r Reader) Get(key []byte) ([]byte, error) {

	// TODO Implement retrieving a single key
//	c, _ := redis.Dial("tcp","127.0.0.1:5679")
// defer c.Close()

b, err := redis.String(r.store.conn.Do("GET",string(key)))
if err != nil  {
	return nil, nil
}
return []byte(b), err

	//is this correct? the problem is I missed the part where redis connection value is happening.

//where do i put the 127.0.0.1:5679?
	
	return nil, fmt.Errorf("Not implemented")
}

// MultiGet retrieves multiple values in one call.
func (r Reader) MultiGet(keys [][]byte) ([][]byte, error) {

	// TODO implement retrieving the keys
	return nil, fmt.Errorf("Not implemented")
}

// PrefixIterator returns a KVIterator that will
// visit all K/V pairs with the provided prefix
func (r Reader) PrefixIterator(prefix []byte) store.KVIterator {
	return Iterator{store: r.store, prefix: prefix}
}

// RangeIterator returns a KVIterator that will
// visit all K/V pairs >= start AND < end
func (r Reader) RangeIterator(start, end []byte) store.KVIterator {
	return Iterator{store: r.store, start: start, end: end}
}

// Close closes the Reader.
func (r Reader) Close() error {

	// TODO Check wether reader must be closed and implement accordingly.
	return fmt.Errorf("Not implemented/checked")
}
