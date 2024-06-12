// Copyright (c) 2022-2024 The Focela Authors, All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package types provides high performance and concurrent-safe basic variable types.
package types

// New is alias of NewAny.
// See NewAny, NewInterface.
func New(value ...interface{}) *Any {
	return NewAny(value...)
}
