// Copyright (c) 2022-2024 The Focela Authors, All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cache

import "sync"

type adapterMemoryExpireTimes struct {
	mu          sync.RWMutex          // expireTimeMu ensures the concurrent safety of expireTimes map.
	expireTimes map[interface{}]int64 // expireTimes is the expiring key to its timestamp mapping, which is used for quick indexing and deleting.
}

func newAdapterMemoryExpireTimes() *adapterMemoryExpireTimes {
	return &adapterMemoryExpireTimes{
		expireTimes: make(map[interface{}]int64),
	}
}

func (d *adapterMemoryExpireTimes) Get(key interface{}) (value int64) {
	d.mu.RLock()
	value = d.expireTimes[key]
	d.mu.RUnlock()
	return
}

func (d *adapterMemoryExpireTimes) Set(key interface{}, value int64) {
	d.mu.Lock()
	d.expireTimes[key] = value
	d.mu.Unlock()
}

func (d *adapterMemoryExpireTimes) Delete(key interface{}) {
	d.mu.Lock()
	delete(d.expireTimes, key)
	d.mu.Unlock()
}
