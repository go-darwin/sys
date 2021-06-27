// SPDX-FileCopyrightText: 2021 The Go Darwin Authors
// SPDX-License-Identifier: BSD-3-Clause

//go:build darwin && amd64 && gc
// +build darwin,amd64,gc

package sys

import (
	"unsafe"
)

// Noescape hides a pointer from escape analysis.
//
// Noescape is the identity function but escape analysis doesn't think the
// output depends on the input.
//
// Noescape is inlined and currently compiles down to zero instructions.
//go:nosplit
//go:nocheckptr
func Noescape(p unsafe.Pointer) unsafe.Pointer {
	x := uintptr(p)
	return unsafe.Pointer(x ^ 0)
}
