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
/*

	b, err := r.snapshot.Get(key, r.store.defaultReadOptions)
	if err == leveldb.ErrNotFound {
		return nil, nil
	}
	return b, err
	// TODO Implement retrieving a single key
*/
	c, _ := redis.Dial("tcp","127.0.0.1:5679")
        defer c.Close()

	b, err := redis.String(c.Do("GET",string(key))) //i'm not sure how to call c.Do, not sure if it's correct.
	if err != nil  {
		return nil, nil
	}
	return []byte(b), err
}

// MultiGet retrieves multiple values in one call.
func (r Reader) MultiGet(keys [][]byte) ([][]byte, error) {
//	return store.MultiGet(r, keys)
        c, _ := redis.Dial("tcp","127.0.0.1:5679")
        defer c.Close()

	
	args := make([]interface{}, len(keys))

	for i := 0; i < len(keys); i++ {
		args[i] = string(keys[i])
	}

	redisStrAry, err := redis.String(c.Do("MGET",args...)) //i'm not sure how to call c.Do, not sure if it's correct.
	
        var redisStrByte [][]byte

	for i := 0; i < len(redisStrAry); i++ {
		redisStrByte[i] = []byte(redisStrAry[i])
	}


	return redisStrByte
	// TODO implement retrieving the keys
	//return nil, fmt.Errorf("Not implemented")
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
//I'm not sure how to do this. 
	//defer c.Close() ???

	// TODO Check wether reader must be closed and implement accordingly.
	return fmt.Errorf("Not implemented/checked")
}
