// Copyright (c) 2022-2024 The Focela Authors, All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package types

import (
	"strconv"
	"sync/atomic"

	"github.com/gocarp/utils/conv"
)

// Byte is a struct for concurrent-safe operation for type byte.
type Byte struct {
	value int32
}

// NewByte creates and returns a concurrent-safe object for byte type,
// with given initial value `value`.
func NewByte(value ...byte) *Byte {
	if len(value) > 0 {
		return &Byte{
			value: int32(value[0]),
		}
	}
	return &Byte{}
}

// Clone clones and returns a new concurrent-safe object for byte type.
func (v *Byte) Clone() *Byte {
	return NewByte(v.Val())
}

// Set atomically stores `value` into t.value and returns the previous value of t.value.
func (v *Byte) Set(value byte) (old byte) {
	return byte(atomic.SwapInt32(&v.value, int32(value)))
}

// Val atomically loads and returns t.value.
func (v *Byte) Val() byte {
	return byte(atomic.LoadInt32(&v.value))
}

// Add atomically adds `delta` to t.value and returns the new value.
func (v *Byte) Add(delta byte) (new byte) {
	return byte(atomic.AddInt32(&v.value, int32(delta)))
}

// Cas executes the compare-and-swap operation for value.
func (v *Byte) Cas(old, new byte) (swapped bool) {
	return atomic.CompareAndSwapInt32(&v.value, int32(old), int32(new))
}

// String implements String interface for string printing.
func (v *Byte) String() string {
	return strconv.FormatUint(uint64(v.Val()), 10)
}

// MarshalJSON implements the interface MarshalJSON for json.Marshal.
func (v Byte) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatUint(uint64(v.Val()), 10)), nil
}

// UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.
func (v *Byte) UnmarshalJSON(b []byte) error {
	v.Set(conv.Uint8(string(b)))
	return nil
}

// UnmarshalValue is an interface implement which sets any type of value for `v`.
func (v *Byte) UnmarshalValue(value interface{}) error {
	v.Set(conv.Byte(value))
	return nil
}

// DeepCopy implements interface for deep copy of current type.
func (v *Byte) DeepCopy() interface{} {
	if v == nil {
		return nil
	}
	return NewByte(v.Val())
}
