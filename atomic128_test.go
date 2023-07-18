package atomic128

import (
	"math/rand"
	"testing"
)

func TestLoadStore(t *testing.T) {
	runTests(t, func(t *testing.T) {
		n := &Uint128{}

		v := LoadUint128(n)
		if got, expected := v, [2]uint64{0, 0}; got != expected {
			t.Fatalf("got %v, expected %v", got, expected)
		}

		StoreUint128(n, [2]uint64{1, ^uint64(0)})
		v = LoadUint128(n)
		if got, expected := v, [2]uint64{1, ^uint64(0)}; got != expected {
			t.Fatalf("got %v, expected %v", got, expected)
		}
	})
}

func TestAdd(t *testing.T) {
	runTests(t, func(t *testing.T) {
		n := &Uint128{}
		v := AddUint128(n, [2]uint64{2, 40})
		if got, expected := v, [2]uint64{2, 40}; got != expected {
			t.Fatalf("got %v, expected %v", got, expected)
		}
		v = LoadUint128(n)
		if got, expected := v, [2]uint64{2, 40}; got != expected {
			t.Fatalf("got %v, expected %v", got, expected)
		}
		v = AddUint128(n, [2]uint64{40, 2})
		if got, expected := v, [2]uint64{42, 42}; got != expected {
			t.Fatalf("got %v, expected %v", got, expected)
		}
		v = LoadUint128(n)
		if got, expected := v, [2]uint64{42, 42}; got != expected {
			t.Fatalf("got %v, expected %v", got, expected)
		}
		v = AddUint128(n, [2]uint64{^uint64(0), 0})
		if got, expected := v, [2]uint64{41, 43}; got != expected {
			t.Fatalf("got %v, expected %v", got, expected)
		}
		v = LoadUint128(n)
		if got, expected := v, [2]uint64{41, 43}; got != expected {
			t.Fatalf("got %v, expected %v", got, expected)
		}
		v = AddUint128(n, [2]uint64{0, ^uint64(0)})
		if got, expected := v, [2]uint64{41, 42}; got != expected {
			t.Fatalf("got %v, expected %v", got, expected)
		}
		v = LoadUint128(n)
		if got, expected := v, [2]uint64{41, 42}; got != expected {
			t.Fatalf("got %v, expected %v", got, expected)
		}
	})
}

func TestCompareAndSwap(t *testing.T) {
	runTests(t, func(t *testing.T) {
		n := &Uint128{}
		StoreUint128(n, [2]uint64{12345, 67890})
		ok := CompareAndSwapUint128(n, [2]uint64{12345, 67890}, [2]uint64{67890, 12345})
		if !ok {
			t.Fatalf("unexpected CAS failure")
		}
		v := LoadUint128(n)
		if got, expected := v, [2]uint64{67890, 12345}; got != expected {
			t.Fatalf("got %v, expected %v", got, expected)
		}
		ok = CompareAndSwapUint128(n, [2]uint64{12345, 67890}, [2]uint64{42, 42})
		if ok {
			t.Fatalf("unexpected CAS success")
		}
		v = LoadUint128(n)
		if got, expected := v, [2]uint64{67890, 12345}; got != expected {
			t.Fatalf("got %v, expected %v", got, expected)
		}
	})
}

func TestSwap(t *testing.T) {
	runTests(t, func(t *testing.T) {
		n := &Uint128{}
		StoreUint128(n, [2]uint64{12345, 67890})
		v := SwapUint128(n, [2]uint64{67890, 12345})
		if got, expected := v, [2]uint64{12345, 67890}; got != expected {
			t.Fatalf("got %v, expected %v", got, expected)
		}
		v = LoadUint128(n)
		if got, expected := v, [2]uint64{67890, 12345}; got != expected {
			t.Fatalf("got %v, expected %v", got, expected)
		}
		v = SwapUint128(n, [2]uint64{42, 42})
		if got, expected := v, [2]uint64{67890, 12345}; got != expected {
			t.Fatalf("got %v, expected %v", got, expected)
		}
		v = LoadUint128(n)
		if got, expected := v, [2]uint64{42, 42}; got != expected {
			t.Fatalf("got %v, expected %v", got, expected)
		}
	})
}

func TestAnd(t *testing.T) {
	runTests(t, func(t *testing.T) {
		n := &Uint128{}
		StoreUint128(n, [2]uint64{0x01234567, 0x89abcdef})
		v := AndUint128(n, [2]uint64{0xffff0000, 0x0000ffff})
		if got, expected := v, [2]uint64{0x01230000, 0x0000cdef}; got != expected {
			t.Fatalf("got %v, expected %v", got, expected)
		}
		v = LoadUint128(n)
		if got, expected := v, [2]uint64{0x01230000, 0x0000cdef}; got != expected {
			t.Fatalf("got %v, expected %v", got, expected)
		}
		v = AndUint128(n, [2]uint64{0x0000ffff, 0xffff0000})
		if got, expected := v, [2]uint64{0, 0}; got != expected {
			t.Fatalf("got %v, expected %v", got, expected)
		}
		v = LoadUint128(n)
		if got, expected := v, [2]uint64{0, 0}; got != expected {
			t.Fatalf("got %v, expected %v", got, expected)
		}
	})
}

func TestOr(t *testing.T) {
	runTests(t, func(t *testing.T) {
		n := &Uint128{}
		StoreUint128(n, [2]uint64{0x01234567, 0x89abcdef})
		v := OrUint128(n, [2]uint64{0xffff0000, 0x0000ffff})
		if got, expected := v, [2]uint64{0xffff4567, 0x89abffff}; got != expected {
			t.Fatalf("got %v, expected %v", got, expected)
		}
		v = LoadUint128(n)
		if got, expected := v, [2]uint64{0xffff4567, 0x89abffff}; got != expected {
			t.Fatalf("got %v, expected %v", got, expected)
		}
		v = OrUint128(n, [2]uint64{0x0000ffff, 0xffff0000})
		if got, expected := v, [2]uint64{0xffffffff, 0xffffffff}; got != expected {
			t.Fatalf("got %v, expected %v", got, expected)
		}
		v = LoadUint128(n)
		if got, expected := v, [2]uint64{0xffffffff, 0xffffffff}; got != expected {
			t.Fatalf("got %v, expected %v", got, expected)
		}
	})
}

func TestXor(t *testing.T) {
	runTests(t, func(t *testing.T) {
		n := &Uint128{}
		StoreUint128(n, [2]uint64{0x01234567, 0x89abcdef})
		v := XorUint128(n, [2]uint64{0xffff0000, 0x0000ffff})
		if got, expected := v, [2]uint64{0x01234567 ^ 0xffff0000, 0x89abcdef ^ 0x0000ffff}; got != expected {
			t.Fatalf("got %v, expected %v", got, expected)
		}
		v = LoadUint128(n)
		if got, expected := v, [2]uint64{0x01234567 ^ 0xffff0000, 0x89abcdef ^ 0x0000ffff}; got != expected {
			t.Fatalf("got %v, expected %v", got, expected)
		}
		v = XorUint128(n, [2]uint64{0x0000ffff, 0xffff0000})
		if got, expected := v, [2]uint64{0x01234567 ^ 0xffffffff, 0x89abcdef ^ 0xffffffff}; got != expected {
			t.Fatalf("got %v, expected %v", got, expected)
		}
		v = LoadUint128(n)
		if got, expected := v, [2]uint64{0x01234567 ^ 0xffffffff, 0x89abcdef ^ 0xffffffff}; got != expected {
			t.Fatalf("got %v, expected %v", got, expected)
		}
	})
}

func BenchmarkLoad(b *testing.B) {
	n := &Uint128{}
	runBenchmarks(b, func(pb *testing.PB) {
		for pb.Next() {
			_ = LoadUint128(n)
		}
	})
}

func BenchmarkStore(b *testing.B) {
	n := &Uint128{}
	runBenchmarks(b, func(pb *testing.PB) {
		i, j := rand.Uint64(), rand.Uint64()
		for pb.Next() {
			StoreUint128(n, [2]uint64{i, j})
		}
	})
}

func BenchmarkSwap(b *testing.B) {
	n := &Uint128{}
	runBenchmarks(b, func(pb *testing.PB) {
		i, j := rand.Uint64(), rand.Uint64()
		for pb.Next() {
			_ = SwapUint128(n, [2]uint64{i, j})
		}
	})
}

func BenchmarkAdd(b *testing.B) {
	n := &Uint128{}
	runBenchmarks(b, func(pb *testing.PB) {
		i, j := rand.Uint64(), rand.Uint64()
		for pb.Next() {
			_ = AddUint128(n, [2]uint64{i, j})
		}
	})
}

func BenchmarkAnd(b *testing.B) {
	n := &Uint128{}
	runBenchmarks(b, func(pb *testing.PB) {
		i, j := rand.Uint64(), rand.Uint64()
		for pb.Next() {
			_ = AndUint128(n, [2]uint64{i, j})
		}
	})
}

func BenchmarkOr(b *testing.B) {
	n := &Uint128{}
	runBenchmarks(b, func(pb *testing.PB) {
		i, j := rand.Uint64(), rand.Uint64()
		for pb.Next() {
			_ = OrUint128(n, [2]uint64{i, j})
		}
	})
}

func BenchmarkXor(b *testing.B) {
	n := &Uint128{}
	runBenchmarks(b, func(pb *testing.PB) {
		i, j := rand.Uint64(), rand.Uint64()
		for pb.Next() {
			_ = XorUint128(n, [2]uint64{i, j})
		}
	})
}

func BenchmarkCAS(b *testing.B) {
	n := &Uint128{}
	_i, _j := rand.Uint64(), rand.Uint64()
	runBenchmarks(b, func(pb *testing.PB) {
		i, j := _i, _j
		for pb.Next() {
			_ = CompareAndSwapUint128(n, [2]uint64{i, j}, [2]uint64{j, i})
			i, j = j, i
		}
	})
}
