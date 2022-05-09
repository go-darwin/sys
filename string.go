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

//go:linkname concatstrings runtime.concatstrings
func concatstrings(buf *TmpBuf, a []string) string

// Concatstrings implements a Go string concatenation x+y+z+...
//
// The operands are passed in the slice a.
// If buf != nil, the compiler has determined that the result does not
// escape the calling function, so the string data can be stored in buf
// if small enough.
func Concatstrings(buf *TmpBuf, a []string) string {
	return concatstrings(buf, a)
}

//go:linkname concatstring2 runtime.concatstring2
func concatstring2(buf *TmpBuf, a0, a1 string) string

// Concatstring2 concats two strings.
func Concatstring2(buf *TmpBuf, a0, a1 string) string {
	return concatstring2(buf, a0, a1)
}

//go:linkname concatstring3 runtime.concatstring3
func concatstring3(buf *TmpBuf, a0, a1, a2 string) string

// Concatstring3 concats three strings.
func Concatstring3(buf *TmpBuf, a0, a1, a2 string) string {
	return concatstring3(buf, a0, a1, a2)
}

//go:linkname concatstring4 runtime.concatstring4
func concatstring4(buf *TmpBuf, a0, a1, a2, a3 string) string

// Concatstring4 concats four strings.
func Concatstring4(buf *TmpBuf, a0, a1, a2, a3 string) string {
	return concatstring4(buf, a0, a1, a2, a3)
}

//go:linkname concatstring5 runtime.concatstring5
func concatstring5(buf *TmpBuf, a0, a1, a2, a3, a4 string) string

// Concatstring5 concats five strings.
func Concatstring5(buf *TmpBuf, a0, a1, a2, a3, a4 string) string {
	return concatstring5(buf, a0, a1, a2, a3, a4)
}

//go:linkname slicebytetostring runtime.slicebytetostring
func slicebytetostring(buf *TmpBuf, ptr *byte, n int) (str string)

// SliceByteToString converts a byte slice to a string.
//
// It is inserted by the compiler into generated code.
// ptr is a pointer to the first element of the slice;
// n is the length of the slice.
//
// Buf is a fixed-size buffer for the result,
// it is not nil if the result does not escape.
func SliceByteToString(buf *TmpBuf, ptr *byte, n int) (str string) {
	return slicebytetostring(buf, ptr, n)
}

//go:linkname stringDataOnStack runtime.stringDataOnStack
func stringDataOnStack(s string) bool

// StringDataOnStack reports whether the string's data is
// stored on the current goroutine's stack.
func StringDataOnStack(s string) bool {
	return stringDataOnStack(s)
}

//go:linkname rawstringtmp runtime.rawstringtmp
func rawstringtmp(buf *TmpBuf, l int) (s string, b []byte)

// RawStringTmp returns a "string" referring to the actual []byte bytes.
func RawStringTmp(buf *TmpBuf, l int) (s string, b []byte) {
	return rawstringtmp(buf, l)
}

//go:linkname slicebytetostringtmp runtime.slicebytetostringtmp
func slicebytetostringtmp(ptr *byte, n int) (str string)

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
//   - Used for m[T1{... Tn{..., string(k), ...} ...}] and m[string(k)]
//     where k is []byte, T1 to Tn is a nesting of struct and array literals.
//   - Used for "<"+string(b)+">" concatenation where b is []byte.
//   - Used for string(b)=="foo" comparison where b is []byte.
func SliceByteToStringTmp(ptr *byte, n int) (str string) {
	return slicebytetostringtmp(ptr, n)
}

//go:linkname stringtoslicebyte runtime.stringtoslicebyte
func stringtoslicebyte(buf *TmpBuf, s string) []byte

// StringToSliceByte converts a string to a byte slice.
func StringToSliceByte(buf *TmpBuf, s string) []byte {
	return stringtoslicebyte(buf, s)
}

//go:linkname stringtoslicerune runtime.stringtoslicerune
func stringtoslicerune(buf *[TmpStringBufSize]rune, s string) []rune

// StringToSliceRune converts a string to a rune slice.
func StringToSliceRune(buf *[TmpStringBufSize]rune, s string) []rune {
	return stringtoslicerune(buf, s)
}

//go:linkname slicerunetostring runtime.slicerunetostring
func slicerunetostring(buf *TmpBuf, a []rune) string

// SliceRuneToString converts a rune slice to a string.
func SliceRuneToString(buf *TmpBuf, a []rune) string {
	return slicerunetostring(buf, a)
}

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

//go:linkname stringStructOf runtime.stringStructOf
func stringStructOf(sp *string) *StringStruct

// StringStructOf converts a sp to StringStruct.
func StringStructOf(sp *string) *StringStruct {
	return stringStructOf(sp)
}

//go:linkname intstring runtime.intstring
func intstring(buf *[4]byte, v int64) (s string)

// IntString converts a int64 v to string.
func IntString(buf *[4]byte, v int64) (s string) {
	return intstring(buf, v)
}

//go:linkname rawstring runtime.rawstring
func rawstring(size int) (s string, b []byte)

// RawString allocates storage for a new string. The returned
// string and byte slice both refer to the same storage.
// The storage is not zeroed. Callers should use
// b to set the string contents and then drop b.
func RawString(size int) (s string, b []byte) {
	return rawstring(size)
}

//go:linkname rawbyteslice runtime.rawbyteslice
func rawbyteslice(size int) (b []byte)

// RawByteSlice allocates a new byte slice. The byte slice is not zeroed.
func RawByteSlice(size int) (b []byte) {
	return rawbyteslice(size)
}

//go:linkname rawruneslice runtime.rawruneslice
func rawruneslice(size int) (b []rune)

// RawRuneSlice allocates a new rune slice. The rune slice is not zeroed.
func RawRuneSlice(size int) (b []rune) {
	return rawruneslice(size)
}

//go:linkname gobytes runtime.gobytes
func gobytes(p *byte, n int) (b []byte)

// GoBytes converts a n length C pointer to Go byte slice.
// This function used by C.GoBytes.
func GoBytes(p *byte, n int) (b []byte) {
	return gobytes(p, n)
}

//go:linkname gostringn runtime.gostringn
func gostringn(p *byte, l int) string

// GostringN converts a l length C string to Go string.
// This function used by C.GostringN.
func GoStringN(p *byte, l int) string {
	return gostringn(p, l)
}

//go:nosplit
//go:linkname findnull runtime.findnull
func findnull(s *byte) int

// FindNull finds NULL in *byte type s.
//
//go:nosplit
func FindNull(s *byte) int {
	return findnull(s)
}

//go:linkname findnullw runtime.findnullw
func findnullw(s *uint16) int

// FindNullW finds NULL in *uint16 type s.
func FindNullW(s *uint16) int {
	return findnullw(s)
}

//go:nosplit
//go:linkname gostringnocopy runtime.gostringnocopy
func gostringnocopy(str *byte) string

// GoString converts a C string to a Go string.
// This function used by C.GoString.
//
//go:nosplit
func GoString(str *byte) string {
	return gostringnocopy(str)
}

func gostring(s *int8) string {
	n, arr := 0, (*[1 << 20]byte)(unsafe.Pointer(s))
	for arr[n] != 0 {
		n++
	}
	return string(arr[:n])
}

//go:linkname gostringw runtime.gostringw
func gostringw(strw *uint16) string

// GoStringW converts a uint16 pointer to a string.
func GoStringW(strw *uint16) string {
	return gostringw(strw)
}
