// Package atomic128 implements atomic operations on 128 bit values.
// When possible (e.g. on amd64 processors that support CMPXCHG16B), it automatically uses
// native CPU features to implement the operations; otherwise it falls back to an approach
// based on mutexes.
package atomic128

// The Uint128 type is defined in architecture-specific files.

func load(d *[3]uint64) [2]uint64 {
	return [2]uint64{d[0], d[1]}
}

func store(d *[3]uint64, v [2]uint64) {
	d[0], d[1] = v[0], v[1]
}

func load2(d *[2]uint64) [2]uint64 {
	return [2]uint64{d[0], d[1]}
}

func store2(d *[2]uint64, v [2]uint64) {
	d[0], d[1] = v[0], v[1]
}
