//go:build !riscv64 || gccgo || appengine
// +build !riscv64 gccgo appengine

package atomic128

func compareAndSwapUint128riscv64(*[2]uint64, [2]uint64, [2]uint64) bool { panic("not implemented") }
func loadUint128riscv64(*[2]uint64) [2]uint64                            { panic("not implemented") }
func storeUint128riscv64(*[2]uint64, [2]uint64)                          { panic("not implemented") }
func swapUint128riscv64(*[2]uint64, [2]uint64) [2]uint64                 { panic("not implemented") }
func addUint128riscv64(ptr *[2]uint64, incr [2]uint64) [2]uint64         { panic("not implemented") }
func andUint128riscv64(ptr *[2]uint64, incr [2]uint64) [2]uint64         { panic("not implemented") }
func orUint128riscv64(ptr *[2]uint64, incr [2]uint64) [2]uint64          { panic("not implemented") }
func xorUint128riscv64(ptr *[2]uint64, incr [2]uint64) [2]uint64         { panic("not implemented") }
