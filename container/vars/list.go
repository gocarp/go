// Copyright (c) 2022-2024 The Focela Authors, All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package vars

import "github.com/gocarp/utils"

// ListItemValues retrieves and returns the elements of all item struct/map with key `key`.
// Note that the parameter `list` should be type of slice which contains elements of map or struct,
// or else it returns an empty slice.
func (v *Var) ListItemValues(key interface{}) (values []interface{}) {
	return utils.ListItemValues(v.Val(), key)
}

// ListItemValuesUnique retrieves and returns the unique elements of all struct/map with key `key`.
// Note that the parameter `list` should be type of slice which contains elements of map or struct,
// or else it returns an empty slice.
func (v *Var) ListItemValuesUnique(key string) []interface{} {
	return utils.ListItemValuesUnique(v.Val(), key)
}
