// Copyright 2021 The Go Darwin Authors
// SPDX-License-Identifier: BSD-3-Clause

//go:build darwin
// +build darwin

package sys

import "unsafe"

// TmpStringBufSize constant is known to the compiler.
// There is no fundamental theory behind this number.
const TmpStringBufSize = 32

// TmpBuf is a temporary byte array.
type TmpBuf [TmpStringBufSize]byte

// Concatstrings implements a Go string concatenation x+y+z+...
//
// The operands are passed in the slice a.
// If buf != nil, the compiler has determined that the result does not
// escape the calling function, so the string data can be stored in buf
// if small enough.
//go:linkname Concatstrings runtime.concatstrings
func Concatstrings(buf *TmpBuf, a []string) string

// Concatstring2 concats two strings.
//go:linkname Concatstring2 runtime.concatstring2
func Concatstring2(buf *TmpBuf, a0, a1 string) string

// Concatstring3 concats three strings.
//go:linkname Concatstring3 runtime.concatstring3
func Concatstring3(buf *TmpBuf, a0, a1, a2 string) string

// Concatstring4 concats four strings.
//go:linkname Concatstring4 runtime.concatstring4
func Concatstring4(buf *TmpBuf, a0, a1, a2, a3 string) string

// Concatstring5 concats five strings.
//go:linkname Concatstring5 runtime.concatstring5
func Concatstring5(buf *TmpBuf, a0, a1, a2, a3, a4 string) string

// SliceByteToString converts a byte slice to a string.
//
// It is inserted by the compiler into generated code.
// ptr is a pointer to the first element of the slice;
// n is the length of the slice.
//
// Buf is a fixed-size buffer for the result,
// it is not nil if the result does not escape.
//go:linkname SliceByteToString runtime.slicebytetostring
func SliceByteToString(buf *TmpBuf, ptr *byte, n int) (str string)

// StringDataOnStack reports whether the string's data is
// stored on the current goroutine's stack.
//go:linkname StringDataOnStack runtime.stringDataOnStack
func StringDataOnStack(s string) bool

// RawStringTmp returns a "string" referring to the actual []byte bytes.
//go:linkname RawStringTmp runtime.rawstringtmp
func RawStringTmp(buf *TmpBuf, l int) (s string, b []byte)

// SliceByteToStringTmp returns a "string" referring to the actual []byte bytes.
//
// Callers need to ensure that the returned string will not be used after
// the calling goroutine modifies the original slice or synchronizes with
// another goroutine.
//
// The function is only called when instrumenting
// and otherwise intrinsified by the compiler.
//
// Some internal compiler optimizations use this function.
//  - Used for m[T1{... Tn{..., string(k), ...} ...}] and m[string(k)]
//   where k is []byte, T1 to Tn is a nesting of struct and array literals.
//  - Used for "<"+string(b)+">" concatenation where b is []byte.
//  - Used for string(b)=="foo" comparison where b is []byte.
//go:linkname SliceByteToStringTmp runtime.slicebytetostringtmp
func SliceByteToStringTmp(ptr *byte, n int) (str string)

// StringToSliceByte converts a string to a byte slice.
//go:linkname StringToSliceByte runtime.stringtoslicebyte
func StringToSliceByte(buf *TmpBuf, s string) []byte

// StringToSliceRune converts a string to a rune slice.
//go:linkname StringToSliceRune runtime.stringtoslicerune
func StringToSliceRune(buf *[TmpStringBufSize]rune, s string) []rune

// SliceRuneToString converts a rune slice to a string.
//go:linkname SliceRuneToString runtime.slicerunetostring
func SliceRuneToString(buf *TmpBuf, a []rune) string

// StringStruct actual string type struct.
type StringStruct struct {
	Str    unsafe.Pointer
	Length int
}

// StringStructDWARF variant with *byte pointer type for DWARF debugging.
type StringStructDWARF struct {
	Str    *byte
	Length int
}

// StringStructOf converts a sp to StringStruct.
//go:linkname StringStructOf runtime.StringStructOf
func StringStructOf(sp *string) *StringStruct

// IntString converts a int64 v to string.
//go:linkname IntString runtime.intstring
func IntString(buf *[4]byte, v int64) (s string)

// RawString allocates storage for a new string. The returned
// string and byte slice both refer to the same storage.
// The storage is not zeroed. Callers should use
// b to set the string contents and then drop b.
//go:linkname RawString runtime.rawstring
func RawString(size int) (s string, b []byte)

// RawByteSlice allocates a new byte slice. The byte slice is not zeroed.
//go:linkname RawByteSlice runtime.rawbyteslice
func RawByteSlice(size int) (b []byte)

// RawRuneSlice allocates a new rune slice. The rune slice is not zeroed.
//go:linkname RawRuneSlice runtime.rawruneslice
func RawRuneSlice(size int) (b []rune)

// GoBytes converts a n length C pointer to Go byte slice.
// This function used by C.GoBytes.
//go:linkname GoBytes runtime.gobytes
func GoBytes(p *byte, n int) (b []byte)

// GostringN converts a l length C string to Go string.
// This function used by C.GostringN.
//go:linkname GoStringN runtime.gostringn
func GoStringN(p *byte, l int) string

// FindNull finds NULL in *byte type s.
//go:nosplit
//go:linkname FindNull runtime.findnull
func FindNull(s *byte) int

// FindNullW finds NULL in *uint16 type s.
//go:linkname FindNullW runtime.findnullw
func FindNullW(s *uint16) int

// GoString converts a C string to a Go string.
// This function used by C.GoString.
//go:nosplit
//go:linkname GoString runtime.gostringnocopy
func GoString(str *byte) string

// GoStringW converts a uint16 pointer to a string.
//go:linkname GoStringW runtime.gostringw
func GoStringW(strw *uint16) string
