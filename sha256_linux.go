//go:build linux

package sha256

import (
	"errors"
	"golang.org/x/sys/unix"
	"hash"
	"io"
	"os"
)

const Size = 32
const BlockSize = 64

type Sha256Linux struct {
	data []byte
}

func (s *Sha256Linux) Reset() {
	s.data = nil
}
func (s Sha256Linux) BlockSize() int {
	return BlockSize
}
func (s Sha256Linux) Size() int {
	return Size
}
func (s *Sha256Linux) Write(b []byte) (int, error) {
	s.data = append(s.data, b...)
	return len(b), nil
}

func (s *Sha256Linux) Sum(b []byte) []byte {
	ret, err := Sum256E(s.data)
	if err != nil {
		return []byte{}
	}
	return append(b, ret...)
}

func New() hash.Hash {
	return &Sha256Linux{}
}

func createSocket(s string) (int, error) {
	fd, err := unix.Socket(unix.AF_ALG, unix.SOCK_SEQPACKET, 0)
	if err != nil {
		return int(-1), err
	}
	sa := &unix.SockaddrALG{Type: "hash", Name: s}

	err = unix.Bind(fd, sa)
	if err != nil {
		return int(-1), err
	}
	return fd, nil
}

func acceptAsFile(fd int, hashType string) (*os.File, error) {
	nfd, _, errno := unix.Syscall(unix.SYS_ACCEPT, uintptr(fd), 0, 0)
	if errno != 0 {
		var err error = errno
		return nil, err
	}

	f := os.NewFile(nfd, hashType)
	if f == nil {
		return f, errors.New("NewFile error")
	}

	return f, nil
}

func Sum256E(data []byte) ([]byte, error) {
	fd, err := createSocket("sha256")
	if err != nil {
		return []byte{}, err
	}
	defer unix.Close(fd)

	f, err := acceptAsFile(fd, "sha256")
	if err != nil {
		return []byte{}, err
	}
	defer f.Close()

	_, err = f.Write(data)
	if err != nil {
		return []byte{}, err
	}

	f.Seek(0, 0)
	buf := make([]byte, Size)
	_, err = f.Read(buf)
	if err != nil && err != io.EOF {
		return []byte{}, err
	}
	return buf, nil
}

func Sum256(data []byte) [Size]byte {
	b, _ := Sum256E(data)

	ret := [Size]byte{}
	for i, v := range b {
		ret[i] = v
	}
	return ret
}
