// Copyright 2021 The Go Darwin Authors
// SPDX-License-Identifier: BSD-3-Clause

//go:build amd64 && gc
// +build amd64,gc

package sys

import (
	_ "runtime" // for go:linkname
	"unsafe"
)

//go:linkname cgocall runtime.cgocall
//go:noescape
//go:nosplit
func cgocall(fn unsafe.Pointer, arg uintptr) int32

// CgoCall calls cgo fn function.
//
//go:nosplit
func CgoCall(fn unsafe.Pointer, arg uintptr) int32 {
	return cgocall(fn, arg)
}

// CString emulates C.String function without cgo.
//
//go:nosplit
func CString(s string) *C_char {
	n := len(s)
	ret := make([]byte, n+1)
	copy(ret, s)
	ret[n] = '\x00'

	return (*C_char)(unsafe.Pointer(&ret[0]))
}

// CBytes emulates C.Bytes function without cgo.
//
//go:nosplit
func CBytes(b []byte) uintptr {
	p := (*[]byte)(unsafe.Pointer(&make([]byte, unsafe.Sizeof([]byte("")))[0]))
	*p = b

	return uintptr(unsafe.Pointer(&p))
}
