//go:build arm64 && !gccgo && !appengine && !arm64_casp
// +build arm64,!gccgo,!appengine,!arm64_casp

#include "textflag.h"

TEXT ·compareAndSwapUint128(SB),NOSPLIT,$0
    MOVD ptr+0(FP), R0
    MOVD old_0+8(FP), R1
    MOVD old_1+16(FP), R2
    MOVD new_0+24(FP), R3
    MOVD new_1+32(FP), R4
loop:
    LDAXP (R0), (R5, R6)
    CMP R1, R5
    BNE fail
    CMP R2, R6
    BNE fail
    STLXP (R3, R4), (R0), R7
    CBNZ R7, loop
    MOVD $1, R7
    MOVB R7, ret+40(FP)
    RET
fail:
    MOVD $0, R7
    MOVB R7, ret+40(FP)
    RET

TEXT ·loadUint128(SB),NOSPLIT,$0
    MOVD ptr+0(FP), R0
loop:
    LDAXP (R0), (R1, R2)
    STLXP (R1, R2), (R0), R3
    CBNZ R3, loop
    MOVD R1, ret_0+8(FP)
    MOVD R2, ret_1+16(FP)
    RET

TEXT ·storeUint128(SB),NOSPLIT,$0
    MOVD ptr+0(FP), R0
    MOVD new_0+8(FP), R1
    MOVD new_1+16(FP), R2
loop:
    LDAXP (R0), (R3, R4)
    STLXP (R1, R2), (R0), R5
    CBNZ R5, loop
    RET

TEXT ·swapUint128(SB),NOSPLIT,$0
    MOVD ptr+0(FP), R0
    MOVD new_0+8(FP), R1
    MOVD new_1+16(FP), R2
loop:
    LDAXP (R0), (R3, R4)
    STLXP (R1, R2), (R0), R5
    CBNZ R5, loop
    MOVD R3, ret_0+24(FP)
    MOVD R4, ret_1+32(FP)
    RET

TEXT ·addUint128(SB),NOSPLIT,$0
    MOVD ptr+0(FP), R0
    MOVD incr_0+8(FP), R1
    MOVD incr_1+16(FP), R2
loop:
    LDAXP (R0), (R3, R4)
    ADDS R1, R3, R5
    ADC R2, R4, R6
    STLXP (R5, R6), (R0), R7
    CBNZ R7, loop
    MOVD R5, ret_0+24(FP)
    MOVD R6, ret_1+32(FP)
    RET

TEXT ·andUint128(SB),NOSPLIT,$0
    MOVD ptr+0(FP), R0
    MOVD op_0+8(FP), R1
    MOVD op_1+16(FP), R2
loop:
    LDAXP (R0), (R3, R4)
    AND R1, R3, R5
    AND R2, R4, R6
    STLXP (R5, R6), (R0), R7
    CBNZ R7, loop
    MOVD R5, ret_0+24(FP)
    MOVD R6, ret_1+32(FP)
    RET

TEXT ·orUint128(SB),NOSPLIT,$0
    MOVD ptr+0(FP), R0
    MOVD op_0+8(FP), R1
    MOVD op_1+16(FP), R2
loop:
    LDAXP (R0), (R3, R4)
    ORR R1, R3, R5
    ORR R2, R4, R6
    STLXP (R5, R6), (R0), R7
    CBNZ R7, loop
    MOVD R5, ret_0+24(FP)
    MOVD R6, ret_1+32(FP)
    RET

TEXT ·xorUint128(SB),NOSPLIT,$0
    MOVD ptr+0(FP), R0
    MOVD op_0+8(FP), R1
    MOVD op_1+16(FP), R2
loop:
    LDAXP (R0), (R3, R4)
    EOR R1, R3, R5
    EOR R2, R4, R6
    STLXP (R5, R6), (R0), R7
    CBNZ R7, loop
    MOVD R5, ret_0+24(FP)
    MOVD R6, ret_1+32(FP)
    RET
