package mmap

import (
	"testing"
)

func TestRW(t *testing.T) {
	m := InitMmap(".mmap", 20)
	m.Write([]byte("hello word"))
	//m.Write([]byte{0x01, 0x02, 0x03, 0x04, 0x05})
	m.Destroy()
}

func TestR(t *testing.T) {
	m := InitMmap(".mmap", 20)
	defer m.Destroy()
	bytes, err := m.Read()
	if err != nil {
		t.Error(err)
	}
	t.Log(string(bytes))

	//t.Logf("%+v", bytes)

}
