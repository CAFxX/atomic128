//go:build (!amd64 && !arm64) || gccgo || appengine
// +build !amd64,!arm64 gccgo appengine

package atomic128

func compareAndSwapUint128(*[2]uint64, [2]uint64, [2]uint64) bool { panic("not implemented") }
func loadUint128(*[2]uint64) [2]uint64                            { panic("not implemented") }
func storeUint128(*[2]uint64, [2]uint64)                          { panic("not implemented") }
func swapUint128(*[2]uint64, [2]uint64) [2]uint64                 { panic("not implemented") }
func addUint128(ptr *[2]uint64, incr [2]uint64) [2]uint64         { panic("not implemented") }
func andUint128(ptr *[2]uint64, incr [2]uint64) [2]uint64         { panic("not implemented") }
func orUint128(ptr *[2]uint64, incr [2]uint64) [2]uint64          { panic("not implemented") }
func xorUint128(ptr *[2]uint64, incr [2]uint64) [2]uint64         { panic("not implemented") }
