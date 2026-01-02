// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build wasip3

package syscall

func Socket(proto, sotype, unused int) (fd int, err error) {
	panic("todo: Socket")
	return 0, nil
}

func Bind(fd int, sa Sockaddr) error {
	panic("todo: Bind")
	return nil
}

func StopIO(fd int) error {
	panic("todo: StopIO")
	return nil
}

func Listen(fd int, backlog int) error {
	panic("todo: Listen")
	return nil
}

func Accept(fd int) (int, Sockaddr, error) {
	panic("todo: Accept")
	return 0, nil, nil
}

func Connect(fd int, sa Sockaddr) error {
	panic("todo: Connect")
	return nil
}

func Recvfrom(fd int, p []byte, flags int) (n int, from Sockaddr, err error) {
	panic("todo: Recvfrom")
	return 0, nil, nil
}

func Sendto(fd int, p []byte, flags int, to Sockaddr) error {
	panic("todo: Sendto")
	return nil
}

func Recvmsg(fd int, p, oob []byte, flags int) (n, oobn, recvflags int, from Sockaddr, err error) {
	panic("todo: Recvmsg")
	return 0, 0, 0, nil, nil
}

func SendmsgN(fd int, p, oob []byte, to Sockaddr, flags int) (n int, err error) {
	panic("todo: SendmsgN")
	return 0, nil
}

func GetsockoptInt(fd, level, opt int) (value int, err error) {
	panic("todo: GetsockoptInt")
	return 0, nil
}

func SetsockoptInt(fd, level, opt int, value int) error {
	panic("todo: SetsockoptInt")
	return nil
}

func SetReadDeadline(fd int, t int64) error {
	panic("todo: SetReadDeadline")
	return nil
}

func SetWriteDeadline(fd int, t int64) error {
	panic("todo: SetWriteDeadline")
	return nil
}

func Shutdown(fd int, how int) error {
	panic("todo: Shutdown")
	return nil
}
