// Copyright (c) 2022-2024 The Focela Authors, All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cache

import (
	"context"

	"github.com/gocarp/go/container/list"
	"github.com/gocarp/go/container/maps"
	"github.com/gocarp/go/container/types"
	"github.com/gocarp/go/timer"
)

// LRU cache object.
// It uses list.List from stdlib for its underlying doubly linked list.
type adapterMemoryLru struct {
	cache   *AdapterMemory // Parent cache object.
	data    *maps.Map      // Key mapping to the item of the list.
	list    *list.List     // Key list.
	rawList *list.List     // History for key adding.
	closed  *types.Bool    // Closed or not.
}

// newMemCacheLru creates and returns a new LRU object.
func newMemCacheLru(cache *AdapterMemory) *adapterMemoryLru {
	lru := &adapterMemoryLru{
		cache:   cache,
		data:    maps.New(true),
		list:    list.New(true),
		rawList: list.New(true),
		closed:  types.NewBool(),
	}
	return lru
}

// Close closes the LRU object.
func (lru *adapterMemoryLru) Close() {
	lru.closed.Set(true)
}

// Remove deletes the `key` FROM `lru`.
func (lru *adapterMemoryLru) Remove(key interface{}) {
	if v := lru.data.Get(key); v != nil {
		lru.data.Remove(key)
		lru.list.Remove(v.(*list.Element))
	}
}

// Size returns the size of `lru`.
func (lru *adapterMemoryLru) Size() int {
	return lru.data.Size()
}

// Push pushes `key` to the tail of `lru`.
func (lru *adapterMemoryLru) Push(key interface{}) {
	lru.rawList.PushBack(key)
}

// Pop deletes and returns the key from tail of `lru`.
func (lru *adapterMemoryLru) Pop() interface{} {
	if v := lru.list.PopBack(); v != nil {
		lru.data.Remove(v)
		return v
	}
	return nil
}

// SyncAndClear synchronizes the keys from `rawList` to `list` and `data`
// using Least Recently Used algorithm.
func (lru *adapterMemoryLru) SyncAndClear(ctx context.Context) {
	if lru.closed.Val() {
		timer.Exit()
		return
	}
	// Data synchronization.
	var alreadyExistItem interface{}
	for {
		if rawListItem := lru.rawList.PopFront(); rawListItem != nil {
			// Deleting the key from list.
			if alreadyExistItem = lru.data.Get(rawListItem); alreadyExistItem != nil {
				lru.list.Remove(alreadyExistItem.(*list.Element))
			}
			// Pushing key to the head of the list
			// and setting its list item to hash table for quick indexing.
			lru.data.Set(rawListItem, lru.list.PushFront(rawListItem))
		} else {
			break
		}
	}
	// Data cleaning up.
	for clearLength := lru.Size() - lru.cache.cap; clearLength > 0; clearLength-- {
		if topKey := lru.Pop(); topKey != nil {
			lru.cache.clearByKey(topKey, true)
		}
	}
}
