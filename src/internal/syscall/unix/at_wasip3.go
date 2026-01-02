// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build wasip3

package unix

import (
	"syscall"
)

// The values of these constants are not part of the WASI API.
const (
	// UTIME_OMIT is the sentinel value to indicate that a time value should not
	// be changed. It is useful for example to indicate for example with UtimesNano
	// to avoid changing AccessTime or ModifiedTime.
	// Its value must match syscall/fs_wasip3.go
	UTIME_OMIT = -0x2

	AT_REMOVEDIR        = 0x200
	AT_SYMLINK_NOFOLLOW = 0x100
)

func Unlinkat(dirfd int, path string, flags int) error {
	panic("todo: Unlinkat")
	return nil
}

func Openat(dirfd int, path string, flags int, perm uint32) (int, error) {
	panic("todo: Openat")
	return 0, nil
}

func Fstatat(dirfd int, path string, stat *syscall.Stat_t, flags int) error {
	panic("todo: Fstatat")
	return nil
}

func Readlinkat(dirfd int, path string, buf []byte) (int, error) {
	panic("todo: Readlinkat")
	return 0, nil
}

func Mkdirat(dirfd int, path string, mode uint32) error {
	panic("todo: Mkdirat")
	return nil
}

func Fchmodat(dirfd int, path string, mode uint32, flags int) error {
	panic("todo: Fchmodat")
	return nil
}

func Fchownat(dirfd int, path string, uid, gid int, flags int) error {
	panic("todo: Fchownat")
	return nil
}

func Renameat(olddirfd int, oldpath string, newdirfd int, newpath string) error {
	panic("todo: Renameat")
	return nil
}

func Linkat(olddirfd int, oldpath string, newdirfd int, newpath string, flag int) error {
	panic("todo: Linkat")
	return nil
}

func Symlinkat(oldpath string, newdirfd int, newpath string) error {
	panic("todo: Symlinkat")
	return nil
}
