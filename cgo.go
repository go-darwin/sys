// Copyright 2021 The Go Darwin Authors
// SPDX-License-Identifier: BSD-3-Clause

//go:build darwin
// +build darwin

package sys

import (
	_ "runtime" // for go:linkname
	"unsafe"
)

// CgoCall calls cgo fn function.
//go:noescape
//go:nosplit
//go:linkname CgoCall runtime.cgocall
func CgoCall(fn unsafe.Pointer, arg uintptr) int32

// CString emulates C.String function without cgo.
//go:nosplit
func CString(s string) *C_char {
	p := (*string)(unsafe.Pointer(&make([]byte, unsafe.Sizeof(string("")))[0]))
	*p = s
	return (*C_char)(unsafe.Pointer(&p))
}

// CBytes emulates C.Bytes function without cgo.
//go:nosplit
func CBytes(b []byte) uintptr {
	p := (*[]byte)(unsafe.Pointer(&make([]byte, unsafe.Sizeof([]byte("")))[0]))
	*p = b

	return uintptr(unsafe.Pointer(&p))
}
