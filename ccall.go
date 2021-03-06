// Copyright 2021 The Go Darwin Authors
// SPDX-License-Identifier: BSD-3-Clause

//go:build darwin && amd64
// +build darwin,amd64

package sys

import (
	"bytes"
	"unsafe"

	"github.com/go-darwin/sys/unsafeheader"
)

//go:linkname ccall syscall.syscall
//go:noescape
func ccall(fn, a1, a2, a3 uintptr) (r1, r2 uintptr, err Errno)

// Ccall calls a function in libc on behalf of the syscall package.
//
// syscall takes a pointer to a struct like:
//
//	struct {
//	 fn    uintptr
//	 a1    uintptr
//	 a2    uintptr
//	 a3    uintptr
//	 r1    uintptr
//	 r2    uintptr
//	 err   uintptr
//	}
//
// Ccall must be called on the g0 stack with the
// C calling convention (use libcCall).
//
// Ccall expects a 32-bit result and tests for 32-bit -1
// to decide there was an error.
//
//go:nosplit
func Ccall(fn, a1, a2, a3 uintptr) (r1, r2 uintptr, err Errno) {
	return ccall(fn, a1, a2, a3)
}

//go:linkname ccall6 syscall.syscall6
//go:noescape
func ccall6(fn, a1, a2, a3, a4, a5, a6 uintptr) (r1, r2 uintptr, err Errno)

// Ccall6 calls a function in libc on behalf of the syscall package.
//
// Ccall6 takes a pointer to a struct like:
//
//	struct {
//	 fn    uintptr
//	 a1    uintptr
//	 a2    uintptr
//	 a3    uintptr
//	 a4    uintptr
//	 a5    uintptr
//	 a6    uintptr
//	 r1    uintptr
//	 r2    uintptr
//	 err   uintptr
//	}
//
// Ccall6 must be called on the g0 stack with the
// C calling convention (use libcCall).
//
// Ccall6 expects a 32-bit result and tests for 32-bit -1
// to decide there was an error.
//
//go:nosplit
func Ccall6(fn, a1, a2, a3, a4, a5, a6 uintptr) (r1, r2 uintptr, err Errno) {
	return ccall6(fn, a1, a2, a3, a4, a5, a6)
}

//go:linkname ccall6X syscall.syscall6X
//go:noescape
func ccall6X(fn, a1, a2, a3, a4, a5, a6 uintptr) (r1, r2 uintptr, err Errno)

// Ccall6X calls a function in libc on behalf of the syscall package.
//
// Ccall6X takes a pointer to a struct like:
//
//	struct {
//	 fn    uintptr
//	 a1    uintptr
//	 a2    uintptr
//	 a3    uintptr
//	 a4    uintptr
//	 a5    uintptr
//	 a6    uintptr
//	 r1    uintptr
//	 r2    uintptr
//	 err   uintptr
//	}
//
// Ccall6X must be called on the g0 stack with the
// C calling convention (use libcCall).
//
// Ccall6X is like syscall6 but expects a 64-bit result
// and tests for 64-bit -1 to decide there was an error.
//
//go:nosplit
func Ccall6X(fn, a1, a2, a3, a4, a5, a6 uintptr) (r1, r2 uintptr, err Errno) {
	return ccall6X(fn, a1, a2, a3, a4, a5, a6)
}

// Ccall9 calls a function in libc on behalf of the syscall package.
//
// Ccall9 takes a pointer to a struct like:
//
//	struct {
//	 fn    uintptr
//	 a1    uintptr
//	 a2    uintptr
//	 a3    uintptr
//	 a4    uintptr
//	 a5    uintptr
//	 a6    uintptr
//	 a7    uintptr
//	 a8    uintptr
//	 a9    uintptr
//	 r1    uintptr
//	 r2    uintptr
//	 err   uintptr
//	}
//
// Ccall9 must be called on the g0 stack with the
// C calling convention (use libcCall).
//
// Ccall9 expects a 32-bit result and tests for 32-bit -1
// to decide there was an error.
//
//go:linkname Ccall9 syscall.Syscall9
//go:nosplit
func Ccall9(fn, a1, a2, a3, a4, a5, a6, a7, a8, a9 uintptr) (r1, r2 uintptr, err Errno)

//go:linkname ccallPtr syscall.syscallPtr
//go:noescape
func ccallPtr(fn, a1, a2, a3 uintptr) (r1, r2 uintptr, err Errno)

// CcallPtr is like syscallX except that the libc function reports an
// error by returning NULL and setting errno.
//
//go:nosplit
func CcallPtr(fn, a1, a2, a3 uintptr) (r1, r2 uintptr, err Errno) {
	return ccallPtr(fn, a1, a2, a3)
}

//go:linkname rawCcall syscall.rawSyscall
//go:noescape
func rawCcall(fn, a1, a2, a3 uintptr) (r1, r2 uintptr, err Errno)

// RawCcall calls a function in libc on behalf of the syscall package.
//
//go:nosplit
func RawCcall(fn, a1, a2, a3 uintptr) (r1, r2 uintptr, err Errno) {
	return rawCcall(fn, a1, a2, a3)
}

//go:linkname rawSyscall6 syscall.rawSyscall6
//go:noescape
func rawSyscall6(fn, a1, a2, a3, a4, a5, a6 uintptr) (r1, r2 uintptr, err Errno)

// RawCcall6 calls a function in libc on behalf of the syscall package.
//
//go:nosplit
func RawCcall6(fn, a1, a2, a3, a4, a5, a6 uintptr) (r1, r2 uintptr, err Errno) {
	return rawSyscall6(fn, a1, a2, a3, a4, a5, a6)
}

// RawCcall9 calls a function in libc on behalf of the syscall package.
//
//go:noescape
//go:nosplit
func RawCcall9(fn, a1, a2, a3, a4, a5, a6, a7, a8, a9 uintptr) (r1, r2 uintptr, err Errno)

// ByteSliceFromString returns a NUL-terminated slice of bytes
// containing the text of s.
func ByteSliceFromString(s string) []byte {
	a := make([]byte, len(s)+1)
	copy(a, s)

	return a
}

// BytePtrFromString returns a pointer to a NUL-terminated array of
// bytes containing the text of s.
func BytePtrFromString(s string) *byte {
	a := ByteSliceFromString(s)

	return &a[0]
}

// ByteSliceToString returns a string form of the text represented by the slice s, with a terminating NUL and any
// bytes after the NUL removed.
func ByteSliceToString(s []byte) string {
	if i := bytes.IndexByte(s, 0); i != -1 {
		s = s[:i]
	}

	return string(s)
}

// BytePtrToString takes a pointer to a sequence of text and returns the corresponding string.
// If the pointer is nil, it returns the empty string. It assumes that the text sequence is terminated
// at a zero byte; if the zero byte is not present, the program may crash.
func BytePtrToString(p *byte) string {
	if p == nil || *p == 0 {
		return ""
	}

	// Find NUL terminator.
	n := 0
	for ptr := unsafe.Pointer(p); *(*byte)(ptr) != 0; n++ {
		ptr = unsafe.Pointer(uintptr(ptr) + 1)
	}

	var b []byte
	h := (*unsafeheader.Slice)(unsafe.Pointer(&b))
	h.Data = unsafe.Pointer(p)
	h.Len = n
	h.Cap = n

	return string(b)
}
