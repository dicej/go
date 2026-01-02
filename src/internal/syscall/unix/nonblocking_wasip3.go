// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build wasip3

package unix

func IsNonblock(fd int) (nonblocking bool, err error) {
	panic("todo: IsNonblock")
	return true, nil
}

func HasNonblockFlag(flag int) bool {
	panic("todo: HasNonblock")
	return true
}
