// Copyright (c) 2017, Tom Thorogood
// Copyright (c) 2021, Carlo Alberto Ferraris
// All rights reserved.
// Use of this source code is governed by a
// Modified BSD License that can be found in
// the LICENSE file.

//go:build amd64 && !gccgo && !appengine
// +build amd64,!gccgo,!appengine

#include "textflag.h"

TEXT ·swapUint128amd64(SB),NOSPLIT,$0
	MOVQ addr+0(FP), R8
	ADDQ $15, R8
	ANDQ $-16, R8
    MOVQ 0(R8), AX
    MOVQ 8(R8), DX
	MOVQ new+8(FP), BX
	MOVQ new+16(FP), CX

	LOCK
	CMPXCHG16B (R8)
    JE done
loop:
	LOCK
	CMPXCHG16B (R8)
    JE done
    PAUSE
	JMP loop
done:
	MOVQ AX, old+24(FP)
	MOVQ DX, old+32(FP)
	RET

TEXT ·compareAndSwapUint128amd64(SB),NOSPLIT,$0
	MOVQ addr+0(FP), R8
	ADDQ $15, R8
	ANDQ $-16, R8
	MOVQ old+8(FP), AX
	MOVQ old+16(FP), DX
	MOVQ new+24(FP), BX
	MOVQ new+32(FP), CX
	LOCK
	CMPXCHG16B (R8)
	SETEQ swapped+40(FP)
	RET

TEXT ·loadUint128amd64(SB),NOSPLIT,$0
	MOVQ addr+0(FP), R8
	ADDQ $15, R8
	ANDQ $-16, R8
	XORQ AX, AX
	XORQ DX, DX
	XORQ BX, BX
	XORQ CX, CX
	LOCK
	CMPXCHG16B (R8)
	MOVQ AX, val+8(FP)
	MOVQ DX, val+16(FP)
	RET

TEXT ·loadUint128amd64avx(SB),NOSPLIT,$0
	MOVQ addr+0(FP), R8
	ADDQ $15, R8
	ANDQ $-16, R8
    MOVOA (R8), X1
	MOVOU X1, val+8(FP)
	RET

TEXT ·storeUint128amd64(SB),NOSPLIT,$0
	MOVQ addr+0(FP), R8
	ADDQ $15, R8
	ANDQ $-16, R8
    MOVQ 0(R8), AX
    MOVQ 8(R8), DX
	MOVQ new+8(FP), BX
	MOVQ new+16(FP), CX

	LOCK
	CMPXCHG16B (R8)
    JE done
loop:
	LOCK
	CMPXCHG16B (R8)
    JE done
    PAUSE
	JMP loop
done:
	RET

TEXT ·storeUint128amd64avx(SB),NOSPLIT,$0
	MOVQ addr+0(FP), R8
	ADDQ $15, R8
	ANDQ $-16, R8
	MOVOU new+8(FP), X1
	MOVOA X1, (R8)
	RET	

TEXT ·addUint128amd64(SB),NOSPLIT,$0
	MOVQ addr+0(FP), R8
	ADDQ $15, R8
	ANDQ $-16, R8
    MOVQ 0(R8), AX
    MOVQ 8(R8), DX
    MOVQ incr+8(FP), SI
    MOVQ incr+16(FP), DI

    MOVQ AX, BX
    MOVQ DX, CX
    ADDQ SI, BX
    ADCQ DI, CX
	LOCK
	CMPXCHG16B (R8)
    JE done
loop:
    MOVQ AX, BX
    MOVQ DX, CX
    ADDQ SI, BX
    ADCQ DI, CX
	LOCK
	CMPXCHG16B (R8)
    JE done
    PAUSE
	JMP loop
done:
    MOVQ BX, val+24(FP)
    MOVQ CX, val+32(FP)
	RET    

TEXT ·andUint128amd64(SB),NOSPLIT,$0
	MOVQ addr+0(FP), R8
	ADDQ $15, R8
	ANDQ $-16, R8
    MOVQ 0(R8), AX
    MOVQ 8(R8), DX
    MOVQ incr+8(FP), SI
    MOVQ incr+16(FP), DI

    MOVQ AX, BX
    MOVQ DX, CX
    ANDQ SI, BX
    ANDQ DI, CX
	LOCK
	CMPXCHG16B (R8)
    JE done
loop:
    MOVQ AX, BX
    MOVQ DX, CX
    ANDQ SI, BX
    ANDQ DI, CX
	LOCK
	CMPXCHG16B (R8)
    JE done
    PAUSE
	JMP loop
done:
    MOVQ BX, val+24(FP)
    MOVQ CX, val+32(FP)
	RET    

TEXT ·orUint128amd64(SB),NOSPLIT,$0
	MOVQ addr+0(FP), R8
	ADDQ $15, R8
	ANDQ $-16, R8
    MOVQ 0(R8), AX
    MOVQ 8(R8), DX
    MOVQ incr+8(FP), SI
    MOVQ incr+16(FP), DI

    MOVQ AX, BX
    MOVQ DX, CX
    ORQ SI, BX
    ORQ DI, CX
	LOCK
	CMPXCHG16B (R8)
    JE done
loop:
    MOVQ AX, BX
    MOVQ DX, CX
    ORQ SI, BX
    ORQ DI, CX
	LOCK
	CMPXCHG16B (R8)
    JE done
    PAUSE
	JMP loop
done:
    MOVQ BX, val+24(FP)
    MOVQ CX, val+32(FP)
	RET    

TEXT ·xorUint128amd64(SB),NOSPLIT,$0
	MOVQ addr+0(FP), R8
	ADDQ $15, R8
	ANDQ $-16, R8
    MOVQ 0(R8), AX
    MOVQ 8(R8), DX
    MOVQ incr+8(FP), SI
    MOVQ incr+16(FP), DI

    MOVQ AX, BX
    MOVQ DX, CX
    XORQ SI, BX
    XORQ DI, CX
	LOCK
	CMPXCHG16B (R8)
    JE done
loop:
    MOVQ AX, BX
    MOVQ DX, CX
    XORQ SI, BX
    XORQ DI, CX
	LOCK
	CMPXCHG16B (R8)
    JE done
    PAUSE
	JMP loop
done:
    MOVQ BX, val+24(FP)
    MOVQ CX, val+32(FP)
	RET    

// RTM Variants
TEXT ·storeUint128amd64rtm(SB),NOSPLIT,$0-24
	MOVQ ptr+0(FP), R8
	ADDQ $15, R8
	ANDQ $-16, R8
	MOVQ new_0+8(FP), BX
	MOVQ new_1+16(FP), CX

	XBEGIN fallback
	MOVQ BX, 0(R8)
	MOVQ CX, 8(R8)
	XEND
	RET
fallback:
	MOVQ 0(R8), AX
	MOVQ 8(R8), DX
	LOCK
	CMPXCHG16B (R8)
	JE done
loop:
	LOCK
	CMPXCHG16B (R8)
	JE done
	PAUSE
	JMP loop
done:
	RET

TEXT ·swapUint128amd64rtm(SB),NOSPLIT,$0-40
	MOVQ ptr+0(FP), R8
	ADDQ $15, R8
	ANDQ $-16, R8
	MOVQ new_0+8(FP), BX
	MOVQ new_1+16(FP), CX

	XBEGIN fallback
	MOVQ 0(R8), AX
	MOVQ 8(R8), DX
	MOVQ BX, 0(R8)
	MOVQ CX, 8(R8)
	XEND
	JMP done_ret
fallback:
	MOVQ 0(R8), AX
	MOVQ 8(R8), DX
	LOCK
	CMPXCHG16B (R8)
	JE done_ret
loop:
	LOCK
	CMPXCHG16B (R8)
	JE done_ret
	PAUSE
	JMP loop
done_ret:
	MOVQ AX, ret_0+24(FP)
	MOVQ DX, ret_1+32(FP)
	RET

TEXT ·addUint128amd64rtm(SB),NOSPLIT,$0-40
	MOVQ ptr+0(FP), R8
	ADDQ $15, R8
	ANDQ $-16, R8
	MOVQ incr_0+8(FP), SI
	MOVQ incr_1+16(FP), DI

	XBEGIN fallback
	MOVQ 0(R8), BX
	MOVQ 8(R8), CX
	ADDQ SI, BX
	ADCQ DI, CX
	MOVQ BX, 0(R8)
	MOVQ CX, 8(R8)
	XEND
	JMP done_ret
fallback:
	MOVQ 0(R8), AX
	MOVQ 8(R8), DX
	MOVQ AX, BX
	MOVQ DX, CX
	ADDQ SI, BX
	ADCQ DI, CX
	LOCK
	CMPXCHG16B (R8)
	JE done_ret
loop:
	MOVQ AX, BX
	MOVQ DX, CX
	ADDQ SI, BX
	ADCQ DI, CX
	LOCK
	CMPXCHG16B (R8)
	JE done_ret
	PAUSE
	JMP loop
done_ret:
	MOVQ BX, ret_0+24(FP)
	MOVQ CX, ret_1+32(FP)
	RET

TEXT ·andUint128amd64rtm(SB),NOSPLIT,$0-40
	MOVQ ptr+0(FP), R8
	ADDQ $15, R8
	ANDQ $-16, R8
	MOVQ incr_0+8(FP), SI
	MOVQ incr_1+16(FP), DI

	XBEGIN fallback
	MOVQ 0(R8), BX
	MOVQ 8(R8), CX
	ANDQ SI, BX
	ANDQ DI, CX
	MOVQ BX, 0(R8)
	MOVQ CX, 8(R8)
	XEND
	JMP done_ret
fallback:
	MOVQ 0(R8), AX
	MOVQ 8(R8), DX
	MOVQ AX, BX
	MOVQ DX, CX
	ANDQ SI, BX
	ANDQ DI, CX
	LOCK
	CMPXCHG16B (R8)
	JE done_ret
loop:
	MOVQ AX, BX
	MOVQ DX, CX
	ANDQ SI, BX
	ANDQ DI, CX
	LOCK
	CMPXCHG16B (R8)
	JE done_ret
	PAUSE
	JMP loop
done_ret:
	MOVQ BX, ret_0+24(FP)
	MOVQ CX, ret_1+32(FP)
	RET

TEXT ·orUint128amd64rtm(SB),NOSPLIT,$0-40
	MOVQ ptr+0(FP), R8
	ADDQ $15, R8
	ANDQ $-16, R8
	MOVQ incr_0+8(FP), SI
	MOVQ incr_1+16(FP), DI

	XBEGIN fallback
	MOVQ 0(R8), BX
	MOVQ 8(R8), CX
	ORQ SI, BX
	ORQ DI, CX
	MOVQ BX, 0(R8)
	MOVQ CX, 8(R8)
	XEND
	JMP done_ret
fallback:
	MOVQ 0(R8), AX
	MOVQ 8(R8), DX
	MOVQ AX, BX
	MOVQ DX, CX
	ORQ SI, BX
	ORQ DI, CX
	LOCK
	CMPXCHG16B (R8)
	JE done_ret
loop:
	MOVQ AX, BX
	MOVQ DX, CX
	ORQ SI, BX
	ORQ DI, CX
	LOCK
	CMPXCHG16B (R8)
	JE done_ret
	PAUSE
	JMP loop
done_ret:
	MOVQ BX, ret_0+24(FP)
	MOVQ CX, ret_1+32(FP)
	RET

TEXT ·xorUint128amd64rtm(SB),NOSPLIT,$0-40
	MOVQ ptr+0(FP), R8
	ADDQ $15, R8
	ANDQ $-16, R8
	MOVQ incr_0+8(FP), SI
	MOVQ incr_1+16(FP), DI

	XBEGIN fallback
	MOVQ 0(R8), BX
	MOVQ 8(R8), CX
	XORQ SI, BX
	XORQ DI, CX
	MOVQ BX, 0(R8)
	MOVQ CX, 8(R8)
	XEND
	JMP done_ret
fallback:
	MOVQ 0(R8), AX
	MOVQ 8(R8), DX
	MOVQ AX, BX
	MOVQ DX, CX
	XORQ SI, BX
	XORQ DI, CX
	LOCK
	CMPXCHG16B (R8)
	JE done_ret
loop:
	MOVQ AX, BX
	MOVQ DX, CX
	XORQ SI, BX
	XORQ DI, CX
	LOCK
	CMPXCHG16B (R8)
	JE done_ret
	PAUSE
	JMP loop
done_ret:
	MOVQ BX, ret_0+24(FP)
	MOVQ CX, ret_1+32(FP)
	RET
