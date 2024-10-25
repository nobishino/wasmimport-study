//go:build wasip1 && wasm

package main

import (
	"fmt"
	"runtime"
	"unsafe"
)

// GOOS=wasip1 GOARCH=wasm go build -o main
// WASMTIME_LOG=wasmtime_wasi=trace wasmtime main
func main() {
	var buf = []byte("Hello, Wasm\n")
	_, err := write(1, buf)
	if err != nil {
		panic(err)
	}
}

// syscall.Writeのwasip1実装を(ほとんど)コピーしてきたもの
// 参考元は https://github.com/golang/go/blob/master/src/syscall/fs_wasip1.go#L910-L915
func write(fd int, b []byte) (int, error) {
	var nwritten size
	errno := fd_write(int32(fd), makeIOVec(b), 1, unsafe.Pointer(&nwritten))
	runtime.KeepAlive(b)
	return int(nwritten), errnoErr(errno)
}

type size = uint32
type Errno uint32

func (e Errno) Error() string {
	return fmt.Sprintf("errno %d", e)
}

type uintptr32 = uint32

func makeIOVec(b []byte) unsafe.Pointer {
	return unsafe.Pointer(&iovec{
		buf:    uintptr32(uintptr(bytesPointer(b))),
		bufLen: size(len(b)),
	})
}
func bytesPointer(b []byte) unsafe.Pointer {
	return unsafe.Pointer(unsafe.SliceData(b))
}

type iovec struct {
	buf    uintptr32
	bufLen size
}

// 本質的でないので簡略化した
func errnoErr(e Errno) error {
	switch e {
	case 0:
		return nil
	}
	return e
}

//go:wasmimport wasi_snapshot_preview1 fd_write
//go:noescape
func fd_write(fd int32, iovs unsafe.Pointer, iovsLen size, nwritten unsafe.Pointer) Errno
