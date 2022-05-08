// Copyright 2021 The Go Darwin Authors
// SPDX-License-Identifier: BSD-3-Clause

//go:build darwin
// +build darwin

package sys

import (
	_ "unsafe" // for go:linkname
)

// FuncPC returns the entry PC of the function f.
// It assumes that f is a func value. Otherwise the behavior is undefined.
//
// CAREFUL: In programs with plugins, funcPC can return different values
// for the same function (because there are actually multiple copies of
// the same function in the address space). To be safe, don't use the
// results of this function in any == expression. It is only safe to
// use the result as an address at which to start executing code.
//
//go:nosplit
//go:linkname FuncPC runtime.funcPC
func FuncPC(f interface{}) uintptr
