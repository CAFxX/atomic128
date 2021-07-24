// Copyright (c) 2017, Tom Thorogood
// Copyright (c) 2021, Carlo Alberto Ferraris
// All rights reserved.
// Use of this source code is governed by a
// Modified BSD License that can be found in
// the LICENSE file.

// +build arm64,!gccgo,!appengine

#include "textflag.h"

TEXT ·compareAndSwapUint128arm64(SB),NOSPLIT,$0
	MOVD addr+0(FP), R0
	MOVD old+8(FP), R2
	MOVD old+16(FP), R3
	MOVD new+24(FP), R4
	MOVD new+32(FP), R5
	MOVD R2, R6
	MOVD R3, R7
	CASPD (R4, R5), (R0), (R2, R3)
	EOR R2, R6, R6
	EOR R3, R7, R7
	ORR R6, R7, R6
	CMP $0, R6
	CSET EQ, R6
	MOVB R6, swapped+40(FP)
	RET

