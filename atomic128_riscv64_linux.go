//go:build riscv64 && linux
// +build riscv64,linux

package atomic128

import (
	"bytes"
	"os"
)

func init() {
	buf, err := os.ReadFile("/proc/cpuinfo")
	if err == nil {
		useNativeRiscv64 = bytes.Contains(buf, []byte("_zacas")) || bytes.Contains(buf, []byte(" zacas")) || bytes.Contains(buf, []byte("zacas "))
	}
}
