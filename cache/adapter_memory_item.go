// Copyright (c) 2022-2024 The Focela Authors, All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cache

import "github.com/gocarp/go/times"

// IsExpired checks whether `item` is expired.
func (item *adapterMemoryItem) IsExpired() bool {
	// Note that it should use greater than or equal judgement here
	// imagining that the cache time is only 1 millisecond.

	return item.e < times.TimestampMilli()
}
