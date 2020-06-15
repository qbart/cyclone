package cyclone

import (
	"strconv"

	"github.com/mediocregopher/radix/v3"
)

// Hash wraps redis hash operations.
type Hash struct {
	cyclone *Cyclone
	key     string
}

// HashScanIterator allows for channel based iteration.
type HashScanIterator struct {
	hash    *Hash
	cursor  string
	count   int
	pattern string
}

// HashField is used in ChanKV iterator as a channel type.
type HashField struct {
	Key string
	Val string
}

// Del removes the specified fields from the hash stored at key. Specified
// fields that do not exist within this hash are ignored. If key does
// not exist, it is treated as an empty hash and this command returns 0.
// https://redis.io/commands/hdel
//
// Time complexity: O(N) where N is the number of fields to be removed.
func (l *Hash) Del(fields ...interface{}) (deletedKeys int) {
	l.cyclone.Raw.Do(radix.FlatCmd(
		&deletedKeys,
		"HDEL",
		l.key,
		fields...,
	))
	return
}

// Exists returns if field is an existing field in the hash stored at key.
// https://redis.io/commands/hexists
//
// Time complexity: O(1)
func (l *Hash) Exists(field string) bool {
	var exists int
	l.cyclone.Raw.Do(radix.Cmd(&exists, "HEXISTS", l.key, field))
	return exists == 1
}

// Get returns the value associated with field in the hash stored at key.
// https://redis.io/commands/hget
//
// Time complexity: O(1)
func (l *Hash) Get(field string) (value string) {
	l.cyclone.Raw.Do(radix.Cmd(&value, "HGET", l.key, field))
	return
}

// GetAll returns all fields and values of the hash stored at key. In the returned
// value, every field name is followed by its value, so the length of the
// reply is twice the size of the hash.
// https://redis.io/commands/hgetall
//
// Time complexity: O(N) where N is the size of the hash.
func (l *Hash) GetAll() (all map[string]string) {
	l.cyclone.Raw.Do(radix.Cmd(&all, "HGETALL", l.key))
	return
}

// Incr increments the number stored at field in the hash stored at key by increment.
// If key does not exist, a new key holding a hash is created.
// If field does not exist the value is set to 0 before the operation is performed.
// https://redis.io/commands/hincrby
//
// Time complexity: O(1)
func (l *Hash) Incr(field string, by int) (valAfterIncr int) {
	l.cyclone.Raw.Do(radix.Cmd(
		&valAfterIncr,
		"HINCRBY",
		l.key,
		field,
		strconv.FormatInt(int64(by), 10),
	))
	return
}

// IncrFloat increments the specified field of a hash stored at key, and representing a
// floating point number, by the specified increment. If the increment value
// is negative, the result is to have the hash field value decremented.
// If the field does not exist, it is set to 0 before performing the operation.
// https://redis.io/commands/hincrbyfloat
//
// Time complexity: O(1)
func (l *Hash) IncrFloat(field string, by float64) (valAfterIncr float64) {
	l.cyclone.Raw.Do(radix.Cmd(
		&valAfterIncr,
		"HINCRBYFLOAT",
		l.key,
		field,
		strconv.FormatFloat(by, 'E', -1, 64),
	))
	return
}

// Keys returns all field names in the hash stored at key.
// https://redis.io/commands/hkeys
//
// Time complexity: O(N) where N is the size of the hash.
func (l *Hash) Keys() (keys []string) {
	l.cyclone.Raw.Do(radix.Cmd(&keys, "HKEYS", l.key))
	return
}

// Len returns the number of fields contained in the hash stored at key.
// https://redis.io/commands/hlen
//
// Time complexity: O(1)
func (l *Hash) Len() (keyCount int) {
	l.cyclone.Raw.Do(radix.Cmd(&keyCount, "HLEN", l.key))
	return
}

// MGet returns the values associated with the specified fields in the hash
// stored at key. For every field that does not exist in the hash, a nil
// value is returned. Because non-existing keys are treated as empty hashes,
// running HMGET against a non-existing key will return a list of nil values.
// https://redis.io/commands/hmget
//
// Time complexity: O(N) where N is the number of fields being requested.
func (l *Hash) MGet(fields ...interface{}) (values []string) {
	l.cyclone.Raw.Do(radix.FlatCmd(
		&values,
		"HMGET",
		l.key,
		fields...,
	))
	return
}

// Scan iterates fields of Hash types and their associated values.
// https://redis.io/commands/hscan
// https://redis.io/commands/scan
//
// Time complexity: O(1) for every call. O(N) for a complete iteration, including
//                  enough command calls for the cursor to return back to 0.
//                  N is the number of elements inside the collection.
func (l *Hash) Scan() *HashScanIterator {
	return &HashScanIterator{hash: l}
}

// Set sets field in the hash stored at key to value. If key does not exist,
// a new key holding a hash is created. If field already exists in the hash,
// it is overwritten.
// https://redis.io/commands/hset
//
// Time complexity: O(1) for each field/value pair added, so O(N) to add N
//                  field/value pairs when the command is called with multiple
//                  field/value pairs.
func (l *Hash) Set(kvpairs ...interface{}) (addedFields int) {
	l.cyclone.Raw.Do(radix.FlatCmd(
		&addedFields,
		"HSET",
		l.key,
		kvpairs...,
	))
	return
}

// SetNX sets field in the hash stored at key to value, only if field does not yet exist.
// If key does not exist, a new key holding a hash is created. If field already exists,
// this operation has no effect.
// https://redis.io/commands/hsetnx
//
// Time complexity: O(1)
func (l *Hash) SetNX(k, v string) bool {
	var wasSet int
	l.cyclone.Raw.Do(radix.Cmd(&wasSet, "HSETNX", l.key, k, v))
	return wasSet == 1
}

// StrLen returns the string length of the value associated with field in the hash stored at key.
// If the key or the field do not exist, 0 is returned.
// https://redis.io/commands/hstrlen
//
// Time complexity: O(1)
func (l *Hash) StrLen(field string) (length int) {
	l.cyclone.Raw.Do(radix.Cmd(&length, "HSTRLEN", l.key, field))
	return
}

// Vals returns all values in the hash stored at key.
// https://redis.io/commands/hvals
//
// Time complexity: O(N) where N is the size of the hash.
func (l *Hash) Vals() (values []string) {
	l.cyclone.Raw.Do(radix.Cmd(&values, "HVALS", l.key))
	return
}

// Count sets count hint for iterator. Redis default hint is 10 when not specified.
// https://redis.io/commands/scan#the-count-option
//
func (i *HashScanIterator) Count(count int) *HashScanIterator {
	i.count = count
	return i
}

// Match sets match pattern for iterator.
// https://redis.io/commands/scan#the-match-option
//
func (i *HashScanIterator) Match(pattern string) *HashScanIterator {
	i.pattern = pattern
	return i
}

// Chan returns channel and starts iteration.
// It will send Key/Values separately.
func (i *HashScanIterator) Chan(bufferSize int) <-chan string {
	ch := make(chan string, bufferSize)

	go func() {
		opts := radix.ScanOpts{
			Command: "HSCAN",
			Key:     i.hash.key,
		}
		if i.pattern != "" {
			opts.Pattern = i.pattern
		}
		if i.count != 0 {
			opts.Count = i.count
		}
		scanner := radix.NewScanner(i.hash.cyclone.Raw, opts)
		var key string
		for scanner.Next(&key) {
			ch <- key
		}
		close(ch)
		if err := scanner.Close(); err != nil {
			panic(err)
		}
	}()
	return ch
}

// ChanKV returns channel and starts iteration.
// It will send HashField struct containing Key and Val.
func (i *HashScanIterator) ChanKV(bufferSize int) <-chan HashField {
	ch := make(chan HashField, bufferSize)

	go func() {
		var (
			field   HashField
			hasNext bool
		)
		opts := radix.ScanOpts{
			Command: "HSCAN",
			Key:     i.hash.key,
		}
		if i.pattern != "" {
			opts.Pattern = i.pattern
		}
		if i.count != 0 {
			opts.Count = i.count
		}
		scanner := radix.NewScanner(i.hash.cyclone.Raw, opts)

		toggle := true
		for {
			if toggle {
				hasNext = scanner.Next(&field.Key)
			} else {
				hasNext = scanner.Next(&field.Val)
				ch <- field
			}
			if !hasNext {
				break
			}
			toggle = !toggle
		}
		close(ch)
		if err := scanner.Close(); err != nil {
			panic(err)
		}
	}()
	return ch
}
