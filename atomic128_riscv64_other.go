//go:build riscv64 && !linux
// +build riscv64,!linux

package atomic128

func init() {
	useNativeRiscv64 = false
}
