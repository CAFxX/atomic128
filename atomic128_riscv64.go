//go:build riscv64 && !gccgo && !appengine
// +build riscv64,!gccgo,!appengine

package atomic128

func compareAndSwapUint128riscv64(*[2]uint64, [2]uint64, [2]uint64) bool
func loadUint128riscv64(*[2]uint64) [2]uint64
func storeUint128riscv64(*[2]uint64, [2]uint64)
func swapUint128riscv64(*[2]uint64, [2]uint64) [2]uint64
func addUint128riscv64(ptr *[2]uint64, incr [2]uint64) [2]uint64
func andUint128riscv64(ptr *[2]uint64, incr [2]uint64) [2]uint64
func orUint128riscv64(ptr *[2]uint64, incr [2]uint64) [2]uint64
func xorUint128riscv64(ptr *[2]uint64, incr [2]uint64) [2]uint64
