// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build wasip3

package syscall

type timestamp = uint64
type dircookie = uint64

type Stat_t struct {
	Dev      uint64
	Ino      uint64
	Filetype uint8
	Nlink    uint64
	Size     uint64
	Atime    uint64
	Mtime    uint64
	Ctime    uint64

	Mode int

	// Uid and Gid are always zero on wasip3 platforms
	Uid uint32
	Gid uint32
}

func Open(path string, openmode int, perm uint32) (int, error) {
	panic("todo: Open")
	return 0, nil
}

func Read(fd int, b []byte) (int, error) {
	panic("todo: Read")
	return 0, nil
}

func Write(fd int, b []byte) (int, error) {
	panic("todo: Write")
	return 0, nil
}

func Pread(fd int, b []byte, offset int64) (int, error) {
	panic("todo: Pread")
	return 0, nil
}

func Pwrite(fd int, b []byte, offset int64) (int, error) {
	panic("todo: Pwrite")
	return 0, nil
}

func Seek(fd int, offset int64, whence int) (int64, error) {
	panic("todo: Seek")
	return 0, nil
}

func Close(fd int) error {
	panic("todo: Close")
	return nil
}

func Fsync(fd int) error {
	panic("todo: Fsync")
	return nil
}

func Ftruncate(fd int, length int64) error {
	panic("todo: Ftruncate")
	return nil
}

func Fchmod(fd int, mode uint32) error {
	panic("todo: Fchmod")
	return nil
}

func Fstat(fd int, st *Stat_t) error {
	panic("todo: Fstat")
	return nil
}

func Fchown(fd int, uid, gid int) error {
	return ENOSYS
}

func Chdir(path string) error {
	panic("todo: Chdir")
	return nil
}

func ReadDir(fd int, buf []byte, cookie dircookie) (int, error) {
	panic("todo: ReadDir")
	return 0, nil
}

func CloseOnExec(fd int) {
	// nothing to do - no exec
}

func Mkdir(path string, perm uint32) error {
	panic("todo: Mkdir")
	return nil
}

func Chmod(path string, mode uint32) error {
	panic("todo: Chmod")
	return nil
}

func Chown(path string, uid, gid int) error {
	return ENOSYS
}

func Lchown(path string, uid, gid int) error {
	return ENOSYS
}

func UtimesNano(path string, ts []Timespec) error {
	panic("todo: UtimesNano")
	return nil
}

func Rename(from, to string) error {
	panic("todo: Rename")
	return nil
}

func Truncate(path string, length int64) error {
	panic("todo: Truncate")
	return nil
}

func Unlink(path string) error {
	panic("todo: Unlink")
	return nil
}

func Rmdir(path string) error {
	panic("todo: Rmdir")
	return nil
}

func Link(path, link string) error {
	panic("todo: Link")
	return nil
}

func Symlink(path, link string) error {
	panic("todo: Symlink")
	return nil
}

func Readlink(path string, buf []byte) (n int, err error) {
	panic("todo: Readlink")
	return 0, nil
}

const ImplementsGetwd = true

func Getwd() (string, error) {
	panic("todo: Getwd")
	return "", nil
}

func Stat(path string, st *Stat_t) error {
	panic("todo: Stat")
	return nil
}

func Lstat(path string, st *Stat_t) error {
	panic("todo: Lstat")
	return nil
}
