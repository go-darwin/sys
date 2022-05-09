// Copyright 2021 The Go Darwin Authors
// SPDX-License-Identifier: BSD-3-Clause

//go:build darwin
// +build darwin

package sys

import (
	_ "syscall" // for go:linkname
	"unsafe"
)

//go:noescape
//go:nosplit
//go:linkname libcCall runtime.libcCall
func libcCall(fn, arg unsafe.Pointer) int32

// LibcCall call fn with arg as its argument. Return what fn returns.
// fn is the raw pc value of the entry point of the desired function.
// Switches to the system stack, if not already there.
// Preserves the calling point as the location where a profiler traceback will begin.
//
//go:nosplit
func LibcCall(fn, arg unsafe.Pointer) int32 {
	return libcCall(fn, arg)
}
