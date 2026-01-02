// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build wasip3

package runtime

func netpollinit() {
	throw("todo: netpollinit")
}

func netpollIsPollDescriptor(fd uintptr) bool {
	return false
}

func netpollopen(fd uintptr, pd *pollDesc) int32 {
	throw("todo: netpollopen")
	return 0
}

func netpollarm(pd *pollDesc, mode int) {
	throw("todo: netpollarm")
}

func netpolldisarm(pd *pollDesc, mode int32) {
	throw("todo: netpolldisarm")
}

func netpollclose(fd uintptr) int32 {
	throw("todo: netpollclose")
	return 0
}

func netpollBreak() {}

func netpoll(delay int64) (gList, int32) {
	throw("todo: netpoll")
	return gList{}, 0
}
