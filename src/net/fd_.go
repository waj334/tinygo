package net

import (
	"os"
	"syscall"
	"time"
)

type netFD struct {
	//pfd poll.FD

	// immutable until Close
	family      int
	sotype      int
	isConnected bool // handshake completed or use of association with peer
	net         string
	laddr       Addr
	raddr       Addr
}

func (fd *netFD) Close() error {
	panic("Not implemented")
}

func (fd *netFD) closeRead() error {
	panic("Not implemented")
}

func (fd *netFD) closeWrite() error {
	panic("Not implemented")
}

func (fd *netFD) Read(p []byte) (n int, err error) {
	panic("Not implemented")
}

func (fd *netFD) readFrom(buf []byte) (int, syscall.Sockaddr, error) {
	panic("Not implemented")
}

func (fd *netFD) readMsg(p []byte, oob []byte) (n, oobn, flags int, sa syscall.Sockaddr, err error) {
	panic("Not implemented")
}

func (fd *netFD) Write(p []byte) (nn int, err error) {
	panic("Not implemented")
}

func (fd *netFD) writeTo(buf []byte, sa syscall.Sockaddr) (int, error) {
	panic("Not implemented")
}

func (fd *netFD) writeMsg(p []byte, oob []byte, sa syscall.Sockaddr) (n int, oobn int, err error) {
	panic("Not implemented")
}

func (fd *netFD) accept() (netfd *netFD, err error) {
	panic("Not implemented")
}

func (fd *netFD) dup() (f *os.File, err error) {
	panic("Not implemented")
}

func (fd *netFD) SetDeadline(t time.Time) error {
	panic("Not implemented")
}

func (fd *netFD) SetReadDeadline(t time.Time) error {
	panic("Not implemented")
}

func (fd *netFD) SetWriteDeadline(t time.Time) error {
	panic("Not implemented")
}

func (fd *netFD) RawControl(f func(uintptr)) error { //From poll
	panic("Not implemented")
}