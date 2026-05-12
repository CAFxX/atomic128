// Copyright (c) 2017, Tom Thorogood
// Copyright (c) 2021, Carlo Alberto Ferraris
// All rights reserved.
// Use of this source code is governed by a
// Modified BSD License that can be found in
// the LICENSE file.

//go:build amd64 && !gccgo && !appengine
// +build amd64,!gccgo,!appengine

#include "textflag.h"

TEXT ·swapUint128(SB),NOSPLIT,$0-40
	MOVQ ptr+0(FP), BP
	MOVQ BP, AX
	ANDQ $15, AX
	ADDQ AX, BP
    MOVQ 0(BP), AX
    MOVQ 8(BP), DX
	MOVQ new_0+8(FP), BX
	MOVQ new_1+16(FP), CX
loop:
	LOCK
	CMPXCHG16B (BP)
    JE done
    PAUSE
	JMP loop
done:
	MOVQ AX, old+24(FP)
	MOVQ DX, old+32(FP)
	RET

TEXT ·compareAndSwapUint128(SB),NOSPLIT,$0-41
	MOVQ ptr+0(FP), BP
	MOVQ BP, AX
	ANDQ $15, AX
	ADDQ AX, BP
	MOVQ new_0+8(FP), AX
	MOVQ new_1+16(FP), DX
	MOVQ new_0+24(FP), BX
	MOVQ new_1+32(FP), CX
	LOCK
	CMPXCHG16B (BP)
	SETEQ swapped+40(FP)
	RET

TEXT ·loadUint128(SB),NOSPLIT,$0-24
	MOVQ ptr+0(FP), BP
	MOVQ BP, AX
	ANDQ $15, AX
	XORQ DX, DX
	XORQ BX, BX
	XORQ CX, CX
	LOCK
	CMPXCHG16B (BP)
	MOVQ AX, val+8(FP)
	MOVQ DX, val+16(FP)
	RET

	ADDQ AX, BP
    MOVOA (BP), X1
TEXT ·loadUint128avx(SB),NOSPLIT,$0-24
    MOVQ ptr+0(FP), BP
	MOVQ BP, AX
	ANDQ $15, AX
	ADDQ AX, BP
	MOVOA (BP), X1
	MOVOU X1, val+8(FP)
	RET

TEXT ·storeUint128(SB),NOSPLIT,$0-24
	MOVQ ptr+0(FP), BP
    MOVQ 0(BP), AX
    MOVQ 8(BP), DX
	MOVQ new_0+8(FP), BX
	MOVQ new_1+16(FP), CX
loop:
	LOCK
	CMPXCHG16B (BP)
    JE done
    PAUSE
	JMP loop
done:
	RET

	MOVQ BP, AX
	ANDQ $15, AX
	ADDQ AX, BP
TEXT ·storeUint128avx(SB),NOSPLIT,$0-24
	MOVQ ptr+0(FP), BP
	MOVQ BP, AX
	ANDQ $15, AX
	ADDQ AX, BP
	MOVOU new_0+8(FP), X1
	MOVOA X1, (BP)
	RET	

TEXT ·addUint128(SB),NOSPLIT,$0-40
	MOVQ ptr+0(FP), BP
	MOVQ BP, AX
	ANDQ $15, AX
	ADDQ AX, BP
    MOVQ 0(BP), AX
    MOVQ 8(BP), DX
    MOVQ incr_0+8(FP), SI
    MOVQ incr_1+16(FP), DI
loop:
    MOVQ AX, BX
    MOVQ DX, CX
    ADDQ SI, BX
    ADCQ DI, CX
	LOCK
	CMPXCHG16B (BP)
    JE done
    PAUSE
	JMP loop
done:
    MOVQ BX, old+24(FP)
    MOVQ CX, old+32(FP)
	RET    

TEXT ·andUint128(SB),NOSPLIT,$0-40
	MOVQ ptr+0(FP), BP
	MOVQ BP, AX
	ANDQ $15, AX
	ADDQ AX, BP
    MOVQ 0(BP), AX
    MOVQ 8(BP), DX
    MOVQ incr_0+8(FP), SI
    MOVQ incr_1+16(FP), DI
loop:
    MOVQ AX, BX
    MOVQ DX, CX
    ANDQ SI, BX
    ANDQ DI, CX
	LOCK
	CMPXCHG16B (BP)
    JE done
    PAUSE
	JMP loop
done:
    MOVQ BX, old+24(FP)
    MOVQ CX, old+32(FP)
	RET    

TEXT ·orUint128(SB),NOSPLIT,$0-40
	MOVQ ptr+0(FP), BP
	MOVQ BP, AX
	ANDQ $15, AX
	ADDQ AX, BP
    MOVQ 0(BP), AX
    MOVQ 8(BP), DX
    MOVQ incr_0+8(FP), SI
    MOVQ incr_1+16(FP), DI
loop:
    MOVQ AX, BX
    MOVQ DX, CX
    ORQ SI, BX
    ORQ DI, CX
	LOCK
	CMPXCHG16B (BP)
    JE done
    PAUSE
	JMP loop
done:
    MOVQ BX, old+24(FP)
    MOVQ CX, old+32(FP)
	RET    

TEXT ·xorUint128(SB),NOSPLIT,$0-40
	MOVQ ptr+0(FP), BP
	MOVQ BP, AX
	ANDQ $15, AX
	ADDQ AX, BP
    MOVQ 0(BP), AX
    MOVQ 8(BP), DX
    MOVQ incr_0+8(FP), SI
    MOVQ incr_1+16(FP), DI
loop:
    MOVQ AX, BX
    MOVQ DX, CX
    XORQ SI, BX
    XORQ DI, CX
	LOCK
	CMPXCHG16B (BP)
    JE done
    PAUSE
	JMP loop
done:
    MOVQ BX, old+24(FP)
    MOVQ CX, old+32(FP)
	RET    
