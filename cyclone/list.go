package cyclone

import (
	"strconv"

	"github.com/mediocregopher/radix/v3"
)

type List struct {
	cyclone *Cyclone
	key     string
}

//TODO: BLPOP
//TODO: BRPOP
//TODO: BRPOPLPUSH

// Index (LINDEX) Returns the element at index index in the list stored at key.
// The index is zero-based, so 0 means the first element, 1 the second element
// and so on. Negative indices can be used to designate elements starting at the
// tail of the list. Here, -1 means the last element, -2 means the penultimate
// and so forth.
// https://redis.io/commands/lindex
//
// Time complexity: O(N) where N is the number of elements to traverse to get to the
//                  element at index. This makes asking for the first or the last
//                  element of the list O(1).
func (l *List) Index(index int) (elem string) {
	l.cyclone.Raw.Do(radix.Cmd(&elem, "LINDEX", l.key, strconv.Itoa(index)))
	return
}

//TODO: LINSERT

// Len returns the length of the list stored at key. If key does not exist,
// it is interpreted as an empty list and 0 is returned.
// https://redis.io/commands/llen
//
// Time complexity: O(1)
func (l *List) Len() (lenOfList int) {
	l.cyclone.Raw.Do(radix.Cmd(&lenOfList, "LLEN", l.key))
	return
}

// Pop (LPOP) Removes and returns the first element of the list stored at key.
// https://redis.io/commands/lpop
//
// Time complexity: O(1)
func (l *List) Pop() (elem string) {
	l.cyclone.Raw.Do(radix.Cmd(&elem, "LPOP", l.key))
	return
}

//TODO: LPOS

// Push (LPUSH) Inserts all the specified values at the head of the list stored at key.
// If key does not exist, it is created as empty list before performing the push
// operations. When key holds a value that is not a list, an error is returned.
// https://redis.io/commands/lpush
//
// Time complexity: O(1) for each element added, so O(N) to add N
//                  elements when the command is called with multiple arguments.
func (l *List) Push(elems ...interface{}) (lenAfterPush int) {
	l.cyclone.Raw.Do(radix.FlatCmd(
		&lenAfterPush,
		"LPUSH",
		l.key,
		elems...,
	))
	return
}

// PushX (LPUSHX) Inserts specified values at the head of the list stored at key,
// only if key already exists and holds a list. In contrary to LPUSH, no operation
// will be performed when key does not yet exist.
// https://redis.io/commands/lpushx
//
// Time complexity: O(1) for each element added, so O(N) to add N elements
//                  when the command is called with multiple arguments.
func (l *List) PushX(elems ...interface{}) (lenAfterPush int) {
	l.cyclone.Raw.Do(radix.FlatCmd(
		&lenAfterPush,
		"LPUSHX",
		l.key,
		elems...,
	))
	return
}

// Range (LRANGE) Returns the specified elements of the list stored at key.
// The offsets start and stop are zero-based indexes, with 0 being the first
// element of the list (the head of the list), 1 being the next element and so on.
//
// These offsets can also be negative numbers indicating offsets starting at the
// end of the list. For example, -1 is the last element of the list, -2 the
// penultimate, and so on.
//
// https://redis.io/commands/lrange
//
// Time complexity: O(S+N) where S is the distance of start offset from HEAD for small
//                  lists, from nearest end (HEAD or TAIL) for large lists; and N is
//                  the number of elements in the specified range.
func (l *List) Range(start, stop int) (elems []string) {
	l.cyclone.Raw.Do(radix.Cmd(
		&elems,
		"LRANGE",
		l.key,
		strconv.Itoa(start),
		strconv.Itoa(stop),
	))
	return
}

// Rem (LREM) Removes the first count occurrences of elements equal to element
// from the list stored at key. The count argument influences the operation
// in the following ways:
// - count > 0: Remove elements equal to element moving from head to tail.
// - count < 0: Remove elements equal to element moving from tail to head.
// - count = 0: Remove all elements equal to element.
// https://redis.io/commands/lrem
//
// Time complexity: O(N+M) where N is the length of the list and M is the
//                  number of elements removed.
func (l *List) Rem(count int, elem string) (removedElems int) {
	l.cyclone.Raw.Do(radix.Cmd(
		&removedElems,
		"LREM",
		l.key,
		strconv.Itoa(count),
		elem,
	))
	return
}

// Set sets the list element at index to element.
// https://redis.io/commands/lset
//
// Time complexity: O(N) where N is the length of the list. Setting either
//                  the first or the last element of the list is O(1).
func (l *List) Set(index int, elem string) bool {
	err := l.cyclone.Raw.Do(radix.Cmd(
		nil,
		"LSET",
		l.key,
		strconv.Itoa(index),
		elem,
	))
	return err == nil
}

// Trim (LTRIM) Trim an existing list so that it will contain only the specified
// range of elements specified. Both start and stop are zero-based indexes,
// where 0 is the first element of the list (the head), 1 the next element and so on.
// https://redis.io/commands/ltrim
//
// Time complexity: O(N) where N is the number of elements to be removed by the operation.
func (l *List) Trim(start, stop int) bool {
	err := l.cyclone.Raw.Do(radix.Cmd(
		nil,
		"LTRIM",
		l.key,
		strconv.Itoa(start),
		strconv.Itoa(stop),
	))
	return err == nil
}

// RPop removes and returns the last element of the list stored at key.
// https://redis.io/commands/rpop
//
// Time complexity: O(1)
func (l *List) RPop() (elem string) {
	l.cyclone.Raw.Do(radix.Cmd(&elem, "RPOP", l.key))
	return
}

//TODO: RPOPLPUSH

// RPush inserts all the specified values at the tail of the list stored at key.
// If key does not exist, it is created as empty list before performing the push operation.
// https://redis.io/commands/rpush
//
// Time complexity: O(1) for each element added, so O(N) to add N elements when
//                  the command is called with multiple arguments.
func (l *List) RPush(elems ...interface{}) (lenAfterPush int) {
	l.cyclone.Raw.Do(radix.FlatCmd(
		&lenAfterPush,
		"RPUSH",
		l.key,
		elems...,
	))
	return
}

// RPushX inserts specified values at the tail of the list stored at key, only
// if key already exists and holds a list. In contrary to RPUSH, no operation
// will be performed when key does not yet exist.
// https://redis.io/commands/rpushx
//
// Time complexity: O(1) for each element added, so O(N) to add N elements
//                  when the command is called with multiple arguments.
func (l *List) RPushX(elems ...interface{}) (lenAfterPush int) {
	l.cyclone.Raw.Do(radix.FlatCmd(
		&lenAfterPush,
		"RPUSHX",
		l.key,
		elems...,
	))
	return
}
