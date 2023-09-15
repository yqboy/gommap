package mmap

import (
	"bytes"
	"os"
	"syscall"
)

type (
	MMAP interface {
		Read() ([]byte, error)
		Write(data []byte) bool
		Destroy()
	}
	mmap struct {
		fd   *os.File
		size int
	}
)

func InitMmap(name string, size int) *mmap {
	fd, _ := os.OpenFile(name, os.O_RDWR|os.O_CREATE, 0644)
	_ = fd.Truncate(int64(size))
	return &mmap{fd, size}
}

func (m *mmap) Read() ([]byte, error) {
	b, err := syscall.Mmap(int(m.fd.Fd()), 0, m.size, syscall.PROT_READ, syscall.MAP_SHARED)
	if err != nil {
		return nil, err
	}
	defer syscall.Munmap(b)
	data := make([]byte, m.size)
	copy(data, b)
	return bytes.TrimRight(data, "\x00"), nil
}

func (m *mmap) Write(data []byte) bool {
	b, err := syscall.Mmap(int(m.fd.Fd()), 0, m.size, syscall.PROT_WRITE, syscall.MAP_SHARED)
	if err != nil {
		return false
	}
	defer syscall.Munmap(b)
	copy(b, data)
	return true
}

func (m *mmap) Destroy() {
	_ = m.fd.Close()
}

//func isExist(str string) bool {
//	_, err := os.Stat(str)
//	return err == nil || os.IsExist(err)
//}
