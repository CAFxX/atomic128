//go:build !amd64 && !arm64

package atomic128

import "testing"

func hasNative() bool {
	return false
}

func fallback(tb testing.TB) {
}
