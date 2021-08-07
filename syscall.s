// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build darwin && amd64 && gc
// +build darwin,amd64,gc

#include "textflag.h"
#include "funcdata.h"

// func RawSyscall9(fn, a1, a2, a3, a4, a5, a6, a7, a8, a9 uintptr) (r1, r2 uintptr, err Errno)
TEXT	·RawSyscall9(SB),NOSPLIT,$0-104
	MOVQ	fn+0(FP), AX	// syscall entry
	MOVQ	a1+8(FP), DI
	MOVQ	a2+16(FP), SI
	MOVQ	a3+24(FP), DX
	MOVQ	a4+32(FP), R10
	MOVQ	a5+40(FP), R8
	MOVQ	a6+48(FP), R9
	MOVQ	a7+56(FP), R11
	MOVQ	a8+64(FP), R12
	MOVQ	a9+72(FP), R13
	SUBQ	$32, SP
	MOVQ	R11, 8(SP)
	MOVQ	R12, 16(SP)
	MOVQ	R13, 24(SP)
	ADDQ	$0x2000000, AX
	SYSCALL
	JCC	ok9
	ADDQ	$32, SP
	MOVQ	$-1, r1+80(FP)
	MOVQ	$0, r2+88(FP)
	MOVQ	AX, err+96(FP)
	RET
ok9:
	ADDQ	$32, SP
	MOVQ	AX, r1+80(FP)
	MOVQ	DX, r2+88(FP)
	MOVQ	$0, err+96(FP)
	RET
