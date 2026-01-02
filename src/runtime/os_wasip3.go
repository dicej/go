// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build wasip3

package runtime

import (
	"unsafe"
)

func exit(code int32) {
	throw("todo")
}

func write1(fd uintptr, p unsafe.Pointer, n int32) int32 {
	throw("todo: write1")
	return 0
}

func usleep(usec uint32) {
	throw("todo: usleep")
}

func readRandom(r []byte) int {
	throw("todo: readRandom")
	return 0
}

func goenvs() {
	throw("todo: goenvs")
}

func walltime() (sec int64, nsec int32) {
	return walltime1()
}

func walltime1() (sec int64, nsec int32) {
	throw("todo: walltime1")
	return 0, 0
}

func nanotime1() int64 {
	throw("todo: nanotime1")
	return 0
}
