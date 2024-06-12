// Copyright (c) 2022-2024 The Focela Authors, All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package vars

import "github.com/gocarp/utils/conv"

// Scan automatically checks the type of `pointer` and converts `params` to `pointer`. It supports `pointer`
// with type of `*map/*[]map/*[]*map/*struct/**struct/*[]struct/*[]*struct` for converting.
//
// See conv.Scan.
func (v *Var) Scan(pointer interface{}, mapping ...map[string]string) error {
	return conv.Scan(v.Val(), pointer, mapping...)
}
