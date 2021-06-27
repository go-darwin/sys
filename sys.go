// SPDX-FileCopyrightText: 2021 The Go Darwin Authors
// SPDX-License-Identifier: BSD-3-Clause

//go:build darwin && amd64 && gc
// +build darwin,amd64,gc

package sys

import (
	_ "syscall" // for go:linkname
	"unsafe"

	"golang.org/x/sys/unix"
)

//go:linkname libcCall runtime.libcCall
//go:nosplit
func libcCall(fn, arg unsafe.Pointer) int32

// LibcCall call fn with arg as its argument. Return what fn returns.
// fn is the raw pc value of the entry point of the desired function.
// Switches to the system stack, if not already there.
// Preserves the calling point as the location where a profiler traceback will begin.
func LibcCall(fn, arg unsafe.Pointer) int32 {
	return libcCall(fn, arg)
}

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
