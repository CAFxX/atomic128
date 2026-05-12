//go:build riscv64 && !gccgo && !appengine
// +build riscv64,!gccgo,!appengine

#include "textflag.h"

// AMOCAS.Q A2, A4, (A0) => WORD $0x2ee5462f
// A0 = ptr, A2/A3 = expected/result, A4/A5 = new value

TEXT ·compareAndSwapUint128riscv64(SB),NOSPLIT,$0
	MOV	addr+0(FP), A0
	MOV	old+8(FP), A2
	MOV	old+16(FP), A3
	MOV	new+24(FP), A4
	MOV	new+32(FP), A5

	// A6 = A2, A7 = A3 to compare later
	MOV A2, A6
	MOV A3, A7

	// amocas.q.aqrl a2, a4, (a0)
	WORD $0x2ee5462f

	// if a2 != a6, it failed
	BNE A2, A6, fail
	BNE A3, A7, fail
	MOV $1, A1
	MOVB A1, swapped+40(FP)
	RET
fail:
	MOV $0, A1
	MOVB A1, swapped+40(FP)
	RET

TEXT ·loadUint128riscv64(SB),NOSPLIT,$0
	MOV	addr+0(FP), A0
	// Load using amocas.q with dummy values to get current value
	MOV $0, A2
	MOV $0, A3
	MOV $0, A4
	MOV $0, A5
	// amocas.q.aqrl a2, a4, (a0)
	WORD $0x2ee5462f

	MOV A2, val+8(FP)
	MOV A3, val+16(FP)
	RET

TEXT ·storeUint128riscv64(SB),NOSPLIT,$0
	MOV	addr+0(FP), A0
	MOV	new+8(FP), A4
	MOV	new+16(FP), A5

	// Load current value first
	MOV 0(A0), A2
	MOV 8(A0), A3
loop:
	MOV A2, A6
	MOV A3, A7
	// amocas.q.aqrl a2, a4, (a0)
	WORD $0x2ee5462f
	BNE A2, A6, loop
	BNE A3, A7, loop
	RET

TEXT ·swapUint128riscv64(SB),NOSPLIT,$0
	MOV	addr+0(FP), A0
	MOV	new+8(FP), A4
	MOV	new+16(FP), A5

	MOV 0(A0), A2
	MOV 8(A0), A3
loop:
	MOV A2, A6
	MOV A3, A7
	// amocas.q.aqrl a2, a4, (a0)
	WORD $0x2ee5462f
	BNE A2, A6, loop
	BNE A3, A7, loop

	MOV A6, old+24(FP)
	MOV A7, old+32(FP)
	RET

TEXT ·addUint128riscv64(SB),NOSPLIT,$0
	MOV	addr+0(FP), A0
	MOV	incr+8(FP), T0
	MOV	incr+16(FP), T1

	MOV 0(A0), A2
	MOV 8(A0), A3
loop:
	MOV A2, A4
	MOV A3, A5
	ADD T0, A4
	SLTU A4, T0, T2 // Check carry out: if A4 < T0, carry = 1
	ADD T2, A5
	ADD T1, A5

	MOV A2, A6
	MOV A3, A7
	// amocas.q.aqrl a2, a4, (a0)
	WORD $0x2ee5462f
	BNE A2, A6, loop
	BNE A3, A7, loop

	MOV A4, val+24(FP)
	MOV A5, val+32(FP)
	RET

TEXT ·andUint128riscv64(SB),NOSPLIT,$0
	MOV	addr+0(FP), A0
	MOV	incr+8(FP), T0
	MOV	incr+16(FP), T1

	MOV 0(A0), A2
	MOV 8(A0), A3
loop:
	MOV A2, A4
	MOV A3, A5
	AND T0, A4
	AND T1, A5

	MOV A2, A6
	MOV A3, A7
	// amocas.q.aqrl a2, a4, (a0)
	WORD $0x2ee5462f
	BNE A2, A6, loop
	BNE A3, A7, loop

	MOV A4, val+24(FP)
	MOV A5, val+32(FP)
	RET

TEXT ·orUint128riscv64(SB),NOSPLIT,$0
	MOV	addr+0(FP), A0
	MOV	incr+8(FP), T0
	MOV	incr+16(FP), T1

	MOV 0(A0), A2
	MOV 8(A0), A3
loop:
	MOV A2, A4
	MOV A3, A5
	OR T0, A4
	OR T1, A5

	MOV A2, A6
	MOV A3, A7
	// amocas.q.aqrl a2, a4, (a0)
	WORD $0x2ee5462f
	BNE A2, A6, loop
	BNE A3, A7, loop

	MOV A4, val+24(FP)
	MOV A5, val+32(FP)
	RET

TEXT ·xorUint128riscv64(SB),NOSPLIT,$0
	MOV	addr+0(FP), A0
	MOV	incr+8(FP), T0
	MOV	incr+16(FP), T1

	MOV 0(A0), A2
	MOV 8(A0), A3
loop:
	MOV A2, A4
	MOV A3, A5
	XOR T0, A4
	XOR T1, A5

	MOV A2, A6
	MOV A3, A7
	// amocas.q.aqrl a2, a4, (a0)
	WORD $0x2ee5462f
	BNE A2, A6, loop
	BNE A3, A7, loop

	MOV A4, val+24(FP)
	MOV A5, val+32(FP)
	RET
