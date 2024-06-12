// Copyright (c) 2022-2024 The Focela Authors, All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package types

// Any is a struct for concurrent-safe operation for type any.
type Any = Interface

// NewAny creates and returns a concurrent-safe object for any type,
// with given initial value `value`.
func NewAny(value ...any) *Any {
	t := &Any{}
	if len(value) > 0 && value[0] != nil {
		t.value.Store(value[0])
	}
	return t
}
