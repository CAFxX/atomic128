#include "textflag.h"

// Load Uint128
TEXT ·loadUint128arm64(SB),NOSPLIT,$0
	MOVD addr+0(FP), R0
	ADD $15, R0
	AND $~15, R0

loop:
	LDAXP (R0), (R2, R3)
	// STLXP is used here just to clear the exclusive monitor and prove we could have written.
	// We write back the same value we loaded, which guarantees an atomic read of the 128-bit value.
	STLXP (R2, R3), (R0), R4
	CBNZ R4, loop

	MOVD R2, val+8(FP)
	MOVD R3, val+16(FP)
	RET

// Store Uint128
TEXT ·storeUint128arm64(SB),NOSPLIT,$0
	MOVD addr+0(FP), R0
	ADD $15, R0
	AND $~15, R0
	MOVD new+8(FP), R2
	MOVD new+16(FP), R3

loop:
	LDAXP (R0), (R4, R5)
	STLXP (R2, R3), (R0), R6
	CBNZ R6, loop

	RET

// Swap Uint128
TEXT ·swapUint128arm64(SB),NOSPLIT,$0
	MOVD addr+0(FP), R0
	ADD $15, R0
	AND $~15, R0
	MOVD new+8(FP), R2
	MOVD new+16(FP), R3

loop:
	LDAXP (R0), (R4, R5)
	STLXP (R2, R3), (R0), R6
	CBNZ R6, loop

	MOVD R4, old+24(FP)
	MOVD R5, old+32(FP)
	RET

// CompareAndSwap Uint128 (Baseline)
TEXT ·compareAndSwapUint128arm64(SB),NOSPLIT,$0
	MOVD addr+0(FP), R0
	ADD $15, R0
	AND $~15, R0
	MOVD old+8(FP), R2
	MOVD old+16(FP), R3
	MOVD new+24(FP), R4
	MOVD new+32(FP), R5

loop:
	LDAXP (R0), (R6, R7)
	CMP R2, R6
	BNE fail
	CMP R3, R7
	BNE fail
	STLXP (R4, R5), (R0), R8
	CBNZ R8, loop

	MOVD $1, R8
	MOVB R8, swapped+40(FP)
	RET
fail:
	// Need to clear exclusive monitor, a dummy store or clrex is often used, but not strictly required
	// if we just bail. However, it's good practice. Go's arm64 assembler supports CLREX.
	CLREX
	MOVD $0, R8
	MOVB R8, swapped+40(FP)
	RET

// CompareAndSwap Uint128 (CASPD variant)
TEXT ·compareAndSwapUint128arm64caspd(SB),NOSPLIT,$0
	MOVD addr+0(FP), R0
	ADD $15, R0
	AND $~15, R0
	MOVD old+8(FP), R2
	MOVD old+16(FP), R3
	MOVD new+24(FP), R4
	MOVD new+32(FP), R5

	// CASPD compares (R2, R3) against memory at (R0).
	// If equal, it stores (R4, R5) to memory.
	// The original memory value is loaded into (R2, R3).
	CASPD (R2, R3), (R0), (R4, R5)

	// Now check if the original memory value (now in R2, R3)
	// equals the old value we passed in.
	MOVD old+8(FP), R6
	MOVD old+16(FP), R7

	CMP R2, R6
	BNE fail
	CMP R3, R7
	BNE fail

	MOVD $1, R8
	MOVB R8, swapped+40(FP)
	RET
fail:
	MOVD $0, R8
	MOVB R8, swapped+40(FP)
	RET

// Add Uint128
TEXT ·addUint128arm64(SB),NOSPLIT,$0
	MOVD addr+0(FP), R0
	ADD $15, R0
	AND $~15, R0
	MOVD incr+8(FP), R2
	MOVD incr+16(FP), R3

loop:
	LDAXP (R0), (R4, R5)

	// Perform addition
	MOVD R4, R6
	MOVD R5, R7
	ADDS R2, R6
	ADC R3, R7

	STLXP (R6, R7), (R0), R8
	CBNZ R8, loop

	MOVD R6, val+24(FP)
	MOVD R7, val+32(FP)
	RET

// And Uint128
TEXT ·andUint128arm64(SB),NOSPLIT,$0
	MOVD addr+0(FP), R0
	ADD $15, R0
	AND $~15, R0
	MOVD incr+8(FP), R2
	MOVD incr+16(FP), R3

loop:
	LDAXP (R0), (R4, R5)

	// Perform AND
	MOVD R4, R6
	MOVD R5, R7
	AND R2, R6
	AND R3, R7

	STLXP (R6, R7), (R0), R8
	CBNZ R8, loop

	MOVD R6, val+24(FP)
	MOVD R7, val+32(FP)
	RET

// Or Uint128
TEXT ·orUint128arm64(SB),NOSPLIT,$0
	MOVD addr+0(FP), R0
	ADD $15, R0
	AND $~15, R0
	MOVD incr+8(FP), R2
	MOVD incr+16(FP), R3

loop:
	LDAXP (R0), (R4, R5)

	// Perform OR
	MOVD R4, R6
	MOVD R5, R7
	ORR R2, R6
	ORR R3, R7

	STLXP (R6, R7), (R0), R8
	CBNZ R8, loop

	MOVD R6, val+24(FP)
	MOVD R7, val+32(FP)
	RET

// Xor Uint128
TEXT ·xorUint128arm64(SB),NOSPLIT,$0
	MOVD addr+0(FP), R0
	ADD $15, R0
	AND $~15, R0
	MOVD incr+8(FP), R2
	MOVD incr+16(FP), R3

loop:
	LDAXP (R0), (R4, R5)

	// Perform XOR
	MOVD R4, R6
	MOVD R5, R7
	EOR R2, R6
	EOR R3, R7

	STLXP (R6, R7), (R0), R8
	CBNZ R8, loop

	MOVD R6, val+24(FP)
	MOVD R7, val+32(FP)
	RET
