// SPDX-FileCopyrightText: 2021 The Go Darwin Authors
// SPDX-License-Identifier: BSD-3-Clause

//go:build darwin && amd64 && gc
// +build darwin,amd64,gc

package sys

import (
	_ "runtime" // for go:linkname
	"unsafe"
)

//go:linkname runtime_cgocall runtime.cgocall
//go:nosplit
func runtime_cgocall(fn unsafe.Pointer, arg uintptr) int32

// CgoCall calls cgo fn function.
func CgoCall(fn unsafe.Pointer, arg uintptr) int32 {
	return runtime_cgocall(fn, arg)
}
