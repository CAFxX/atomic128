//go:build (amd64.v2 || amd64.v3) && !gccgo && !appengine

package atomic128

// Uint128 is an opaque container for an atomic uint128.
// Uint128 must not be copied.
// The zero value is a valid value representing [2]uint64{0, 0}.
type Uint128 struct {
	d [3]uint64
}
