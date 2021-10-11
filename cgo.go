// SPDX-FileCopyrightText: 2021 The Go Darwin Authors
// SPDX-License-Identifier: BSD-3-Clause

//go:build darwin && amd64 && gc
// +build darwin,amd64,gc

package sys

import (
	_ "runtime" // for go:linkname
	"unsafe"
)

//go:nosplit
//go:noescape
//go:linkname runtime_cgocall runtime.cgocall
func runtime_cgocall(fn unsafe.Pointer, arg uintptr) int32

// CgoCall calls cgo fn function.
func CgoCall(fn unsafe.Pointer, arg uintptr) int32 {
	return runtime_cgocall(fn, arg)
}

// CString
func CString(s string) *C_char {
	p := (*string)(unsafe.Pointer(&make([]byte, unsafe.Sizeof(string("")))[0]))
	*p = s
	// p := cmalloc(uint64(len(s) + 1))
	// pp := (*[1 << 30]byte)(p)
	// copy(pp[:], s)
	// pp[len(s)] = 0
	return (*C_char)(unsafe.Pointer(&p))
}

// CBytes
func CBytes(b []byte) unsafe.Pointer {
	p := (*[]byte)(unsafe.Pointer(&make([]byte, unsafe.Sizeof([]byte("")))[0]))
	*p = b
	// p := cmalloc(uint64(len(b)))
	// pp := (*[1 << 30]byte)(p)
	// copy(pp[:], b)
	return unsafe.Pointer(&p)
}
