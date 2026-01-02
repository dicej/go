// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build wasip3

package os

func direntIno(buf []byte) (uint64, bool) {
	panic("todo: direntIno")
	return 0, true
}

func direntReclen(buf []byte) (uint64, bool) {
	panic("todo: direntReclen")
	return 0, true
}

func direntNamlen(buf []byte) (uint64, bool) {
	panic("todo: direntNamlen")
	return 0, true
}

func direntType(buf []byte) FileMode {
	panic("todo: direntType")
	return 0
}
