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

func runTests(t *testing.T, fn func(*testing.T)) {
	if hasNative() {
		t.Run("native", fn)
	}
	t.Run("fallback", func(t *testing.T) {
		fallback(t)
		fn(t)
	})
}

func runBenchmarks(b *testing.B, fn func(*testing.PB)) {
	if hasNative() {
		b.Run("native", func(b *testing.B) {
			b.RunParallel(fn)
		})
	}
	b.Run("fallback", func(b *testing.B) {
		fallback(b)
		b.RunParallel(fn)
	})
}

func hasNative() bool {
	return compareAndSwapUint128 != nil || loadUint128 != nil || storeUint128 != nil || swapUint128 != nil || addUint128 != nil
}

func fallback(tb testing.TB) {
	cas := compareAndSwapUint128
	load := loadUint128
	store := storeUint128
	swap := swapUint128
	add := addUint128
	compareAndSwapUint128 = nil
	loadUint128 = nil
	storeUint128 = nil
	swapUint128 = nil
	addUint128 = nil
	tb.Cleanup(func() {
		compareAndSwapUint128 = cas
		loadUint128 = load
		storeUint128 = store
		swapUint128 = swap
		addUint128 = add
	})
}