//go:build arm64 && !gccgo && !appengine && arm64_casp
// +build arm64,!gccgo,!appengine,arm64_casp

#include "textflag.h"

TEXT ·compareAndSwapUint128(SB),NOSPLIT,$0
    MOVD ptr+0(FP), R4
    MOVD old_0+8(FP), R0
    MOVD old_1+16(FP), R1
    MOVD new_0+24(FP), R2
    MOVD new_1+32(FP), R3
    MOVD R0, R5
    MOVD R1, R6
    DMB $0xb
    CASPD (R0, R1), (R4), (R2, R3)
    DMB $0xb
    CMP R0, R5
    BNE fail
    CMP R1, R6
    BNE fail
    MOVD $1, R7
    MOVB R7, ret+40(FP)
    RET
fail:
    MOVD $0, R7
    MOVB R7, ret+40(FP)
    RET

TEXT ·swapUint128(SB),NOSPLIT,$0
    MOVD ptr+0(FP), R4
    MOVD new_0+8(FP), R2
    MOVD new_1+16(FP), R3
    LDXP (R4), (R0, R1)
loop:
    MOVD R0, R5
    MOVD R1, R6
    DMB $0xb
    CASPD (R0, R1), (R4), (R2, R3)
    DMB $0xb
    CMP R0, R5
    BNE fail
    CMP R1, R6
    BNE fail
    MOVD R0, ret_0+24(FP)
    MOVD R1, ret_1+32(FP)
    RET
fail:
    B loop

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
    MOVD ptr+0(FP), R4
    MOVD new_0+8(FP), R2
    MOVD new_1+16(FP), R3
    LDXP (R4), (R0, R1)
loop:
    MOVD R0, R5
    MOVD R1, R6
    DMB $0xb
    CASPD (R0, R1), (R4), (R2, R3)
    DMB $0xb
    CMP R0, R5
    BNE fail
    CMP R1, R6
    BNE fail
    RET
fail:
    B loop

TEXT ·addUint128(SB),NOSPLIT,$0
    MOVD ptr+0(FP), R4
    MOVD incr_0+8(FP), R8
    MOVD incr_1+16(FP), R9
    LDXP (R4), (R0, R1)
loop:
    ADDS R8, R0, R2
    ADC R9, R1, R3
    MOVD R0, R5
    MOVD R1, R6
    DMB $0xb
    CASPD (R0, R1), (R4), (R2, R3)
    DMB $0xb
    CMP R0, R5
    BNE fail
    CMP R1, R6
    BNE fail
    MOVD R2, ret_0+24(FP)
    MOVD R3, ret_1+32(FP)
    RET
fail:
    B loop

TEXT ·andUint128(SB),NOSPLIT,$0
    MOVD ptr+0(FP), R4
    MOVD op_0+8(FP), R8
    MOVD op_1+16(FP), R9
    LDXP (R4), (R0, R1)
loop:
    AND R8, R0, R2
    AND R9, R1, R3
    MOVD R0, R5
    MOVD R1, R6
    DMB $0xb
    CASPD (R0, R1), (R4), (R2, R3)
    DMB $0xb
    CMP R0, R5
    BNE fail
    CMP R1, R6
    BNE fail
    MOVD R2, ret_0+24(FP)
    MOVD R3, ret_1+32(FP)
    RET
fail:
    B loop

TEXT ·orUint128(SB),NOSPLIT,$0
    MOVD ptr+0(FP), R4
    MOVD op_0+8(FP), R8
    MOVD op_1+16(FP), R9
    LDXP (R4), (R0, R1)
loop:
    ORR R8, R0, R2
    ORR R9, R1, R3
    MOVD R0, R5
    MOVD R1, R6
    DMB $0xb
    CASPD (R0, R1), (R4), (R2, R3)
    DMB $0xb
    CMP R0, R5
    BNE fail
    CMP R1, R6
    BNE fail
    MOVD R2, ret_0+24(FP)
    MOVD R3, ret_1+32(FP)
    RET
fail:
    B loop

TEXT ·xorUint128(SB),NOSPLIT,$0
    MOVD ptr+0(FP), R4
    MOVD op_0+8(FP), R8
    MOVD op_1+16(FP), R9
    LDXP (R4), (R0, R1)
loop:
    EOR R8, R0, R2
    EOR R9, R1, R3
    MOVD R0, R5
    MOVD R1, R6
    DMB $0xb
    CASPD (R0, R1), (R4), (R2, R3)
    DMB $0xb
    CMP R0, R5
    BNE fail
    CMP R1, R6
    BNE fail
    MOVD R2, ret_0+24(FP)
    MOVD R3, ret_1+32(FP)
    RET
fail:
    B loop
