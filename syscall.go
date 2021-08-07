// SPDX-FileCopyrightText: 2021 The Go Darwin Authors
// SPDX-License-Identifier: BSD-3-Clause

//go:build darwin && gc
// +build darwin,gc

package sys

import (
	"bytes"
	"unsafe"

	"golang.org/x/sys/unix"

	"go-darwin.dev/sys/unsafeheader"
)

//go:linkname syscall_syscall syscall.syscall
func syscall_syscall(fn, a1, a2, a3 uintptr) (r1, r2 uintptr, err unix.Errno)

// Syscall calls a function in libc on behalf of the syscall package.
//
// syscall takes a pointer to a struct like:
//  struct {
//   fn    uintptr
//   a1    uintptr
//   a2    uintptr
//   a3    uintptr
//   r1    uintptr
//   r2    uintptr
//   err   uintptr
//  }
//
// syscall must be called on the g0 stack with the
// C calling convention (use libcCall).
//
// syscall expects a 32-bit result and tests for 32-bit -1
// to decide there was an error.
func Syscall(fn, a1, a2, a3 uintptr) (r1, r2 uintptr, err unix.Errno) {
	return syscall_syscall(fn, a1, a2, a3)
}

//go:linkname syscall_syscall6 syscall.syscall6
func syscall_syscall6(fn, a1, a2, a3, a4, a5, a6 uintptr) (r1, r2 uintptr, err unix.Errno)

// Syscall6 calls a function in libc on behalf of the syscall package.
//
// syscall takes a pointer to a struct like:
//  struct {
//   fn    uintptr
//   a1    uintptr
//   a2    uintptr
//   a3    uintptr
//   r1    uintptr
//   r2    uintptr
//   err   uintptr
//  }
//
// syscall must be called on the g0 stack with the
// C calling convention (use libcCall).
//
// syscall expects a 32-bit result and tests for 32-bit -1
// to decide there was an error.
func Syscall6(fn, a1, a2, a3, a4, a5, a6 uintptr) (r1, r2 uintptr, err unix.Errno) {
	return syscall_syscall6(fn, a1, a2, a3, a4, a5, a6)
}

//go:linkname syscall_syscall6X syscall.syscall6X
func syscall_syscall6X(fn, a1, a2, a3, a4, a5, a6 uintptr) (r1, r2 uintptr, err unix.Errno)

// Syscall6X calls a function in libc on behalf of the syscall package.
//
// syscall6X takes a pointer to a struct like:
//  struct {
//   fn    uintptr
//   a1    uintptr
//   a2    uintptr
//   a3    uintptr
//   a4    uintptr
//   a5    uintptr
//   a6    uintptr
//   r1    uintptr
//   r2    uintptr
//   err   uintptr
//  }
//
// syscall6X must be called on the g0 stack with the
// C calling convention (use libcCall).
//
// syscall6X is like syscall6 but expects a 64-bit result
// and tests for 64-bit -1 to decide there was an error.
func Syscall6X(fn, a1, a2, a3, a4, a5, a6 uintptr) (r1, r2 uintptr, err unix.Errno) {
	return syscall_syscall6X(fn, a1, a2, a3, a4, a5, a6)
}

//go:linkname syscall_syscallPtr syscall.syscallPtr
func syscall_syscallPtr(fn, a1, a2, a3 uintptr) (r1, r2 uintptr, err unix.Errno)

// SyscallPtr is like syscallX except that the libc function reports an
// error by returning NULL and setting errno.
func SyscallPtr(fn, a1, a2, a3 uintptr) (r1, r2 uintptr, err unix.Errno) {
	return syscall_syscallPtr(fn, a1, a2, a3)
}

//go:linkname syscall_rawSyscall syscall.rawSyscall
func syscall_rawSyscall(fn, a1, a2, a3 uintptr) (r1, r2 uintptr, err unix.Errno)

// RawSyscall calls a function in libc on behalf of the syscall package.
func RawSyscall(fn, a1, a2, a3 uintptr) (r1, r2 uintptr, err unix.Errno) {
	return syscall_rawSyscall(fn, a1, a2, a3)
}

//go:linkname syscall_rawSyscall6 syscall.rawSyscall6
func syscall_rawSyscall6(fn, a1, a2, a3, a4, a5, a6 uintptr) (r1, r2 uintptr, err unix.Errno)

// RawSyscall6 calls a function in libc on behalf of the syscall package.
func RawSyscall6(fn, a1, a2, a3, a4, a5, a6 uintptr) (r1, r2 uintptr, err unix.Errno) {
	return syscall_rawSyscall6(fn, a1, a2, a3, a4, a5, a6)
}

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
