package atomic128

import (
	"math/rand"
	"testing"
)

func TestLoadStore(t *testing.T) {
	runTests(t, func(t *testing.T) {
		n := &Uint128{}

		v := n.Load()
		if got, expected := v, [2]uint64{0, 0}; got != expected {
			t.Fatalf("got %v, expected %v", got, expected)
		}

		n.Store([2]uint64{1, ^uint64(0)})
		v = n.Load()
		if got, expected := v, [2]uint64{1, ^uint64(0)}; got != expected {
			t.Fatalf("got %v, expected %v", got, expected)
		}
	})
}

func TestAdd(t *testing.T) {
	runTests(t, func(t *testing.T) {
		n := &Uint128{}
		v := n.Add([2]uint64{2, 40})
		if got, expected := v, [2]uint64{2, 40}; got != expected {
			t.Fatalf("got %v, expected %v", got, expected)
		}
		v = n.Load()
		if got, expected := v, [2]uint64{2, 40}; got != expected {
			t.Fatalf("got %v, expected %v", got, expected)
		}
		v = n.Add([2]uint64{40, 2})
		if got, expected := v, [2]uint64{42, 42}; got != expected {
			t.Fatalf("got %v, expected %v", got, expected)
		}
		v = n.Load()
		if got, expected := v, [2]uint64{42, 42}; got != expected {
			t.Fatalf("got %v, expected %v", got, expected)
		}
		v = n.Add([2]uint64{^uint64(0), 0})
		if got, expected := v, [2]uint64{41, 43}; got != expected {
			t.Fatalf("got %v, expected %v", got, expected)
		}
		v = n.Load()
		if got, expected := v, [2]uint64{41, 43}; got != expected {
			t.Fatalf("got %v, expected %v", got, expected)
		}
		v = n.Add([2]uint64{0, ^uint64(0)})
		if got, expected := v, [2]uint64{41, 42}; got != expected {
			t.Fatalf("got %v, expected %v", got, expected)
		}
		v = n.Load()
		if got, expected := v, [2]uint64{41, 42}; got != expected {
			t.Fatalf("got %v, expected %v", got, expected)
		}
	})
}

func TestCompareAndSwap(t *testing.T) {
	runTests(t, func(t *testing.T) {
		n := &Uint128{}
		n.Store([2]uint64{12345, 67890})
		ok := n.CompareAndSwap([2]uint64{12345, 67890}, [2]uint64{67890, 12345})
		if !ok {
			t.Fatalf("unexpected CAS failure")
		}
		v := n.Load()
		if got, expected := v, [2]uint64{67890, 12345}; got != expected {
			t.Fatalf("got %v, expected %v", got, expected)
		}
		ok = n.CompareAndSwap([2]uint64{12345, 67890}, [2]uint64{42, 42})
		if ok {
			t.Fatalf("unexpected CAS success")
		}
		v = n.Load()
		if got, expected := v, [2]uint64{67890, 12345}; got != expected {
			t.Fatalf("got %v, expected %v", got, expected)
		}
	})
}

func TestSwap(t *testing.T) {
	runTests(t, func(t *testing.T) {
		n := &Uint128{}
		n.Store([2]uint64{12345, 67890})
		v := n.Swap([2]uint64{67890, 12345})
		if got, expected := v, [2]uint64{12345, 67890}; got != expected {
			t.Fatalf("got %v, expected %v", got, expected)
		}
		v = n.Load()
		if got, expected := v, [2]uint64{67890, 12345}; got != expected {
			t.Fatalf("got %v, expected %v", got, expected)
		}
		v = n.Swap([2]uint64{42, 42})
		if got, expected := v, [2]uint64{67890, 12345}; got != expected {
			t.Fatalf("got %v, expected %v", got, expected)
		}
		v = n.Load()
		if got, expected := v, [2]uint64{42, 42}; got != expected {
			t.Fatalf("got %v, expected %v", got, expected)
		}
	})
}

func TestAnd(t *testing.T) {
	runTests(t, func(t *testing.T) {
		n := &Uint128{}
		n.Store([2]uint64{0x01234567, 0x89abcdef})
		v := n.And([2]uint64{0xffff0000, 0x0000ffff})
		if got, expected := v, [2]uint64{0x01230000, 0x0000cdef}; got != expected {
			t.Fatalf("got %v, expected %v", got, expected)
		}
		v = n.Load()
		if got, expected := v, [2]uint64{0x01230000, 0x0000cdef}; got != expected {
			t.Fatalf("got %v, expected %v", got, expected)
		}
		v = n.And([2]uint64{0x0000ffff, 0xffff0000})
		if got, expected := v, [2]uint64{0, 0}; got != expected {
			t.Fatalf("got %v, expected %v", got, expected)
		}
		v = n.Load()
		if got, expected := v, [2]uint64{0, 0}; got != expected {
			t.Fatalf("got %v, expected %v", got, expected)
		}
	})
}

func TestOr(t *testing.T) {
	runTests(t, func(t *testing.T) {
		n := &Uint128{}
		n.Store([2]uint64{0x01234567, 0x89abcdef})
		v := n.Or([2]uint64{0xffff0000, 0x0000ffff})
		if got, expected := v, [2]uint64{0xffff4567, 0x89abffff}; got != expected {
			t.Fatalf("got %v, expected %v", got, expected)
		}
		v = n.Load()
		if got, expected := v, [2]uint64{0xffff4567, 0x89abffff}; got != expected {
			t.Fatalf("got %v, expected %v", got, expected)
		}
		v = n.Or([2]uint64{0x0000ffff, 0xffff0000})
		if got, expected := v, [2]uint64{0xffffffff, 0xffffffff}; got != expected {
			t.Fatalf("got %v, expected %v", got, expected)
		}
		v = n.Load()
		if got, expected := v, [2]uint64{0xffffffff, 0xffffffff}; got != expected {
			t.Fatalf("got %v, expected %v", got, expected)
		}
	})
}

func TestXor(t *testing.T) {
	runTests(t, func(t *testing.T) {
		n := &Uint128{}
		n.Store([2]uint64{0x01234567, 0x89abcdef})
		v := n.Xor([2]uint64{0xffff0000, 0x0000ffff})
		if got, expected := v, [2]uint64{0x01234567 ^ 0xffff0000, 0x89abcdef ^ 0x0000ffff}; got != expected {
			t.Fatalf("got %v, expected %v", got, expected)
		}
		v = n.Load()
		if got, expected := v, [2]uint64{0x01234567 ^ 0xffff0000, 0x89abcdef ^ 0x0000ffff}; got != expected {
			t.Fatalf("got %v, expected %v", got, expected)
		}
		v = n.Xor([2]uint64{0x0000ffff, 0xffff0000})
		if got, expected := v, [2]uint64{0x01234567 ^ 0xffffffff, 0x89abcdef ^ 0xffffffff}; got != expected {
			t.Fatalf("got %v, expected %v", got, expected)
		}
		v = n.Load()
		if got, expected := v, [2]uint64{0x01234567 ^ 0xffffffff, 0x89abcdef ^ 0xffffffff}; got != expected {
			t.Fatalf("got %v, expected %v", got, expected)
		}
	})
}

func BenchmarkLoad(b *testing.B) {
	n := &Uint128{}
	runBenchmarks(b, func(pb *testing.PB) {
		for pb.Next() {
			_ = n.Load()
		}
	})
}

func BenchmarkStore(b *testing.B) {
	n := &Uint128{}
	runBenchmarks(b, func(pb *testing.PB) {
		i, j := rand.Uint64(), rand.Uint64()
		for pb.Next() {
			n.Store([2]uint64{i, j})
		}
	})
}

func BenchmarkSwap(b *testing.B) {
	n := &Uint128{}
	runBenchmarks(b, func(pb *testing.PB) {
		i, j := rand.Uint64(), rand.Uint64()
		for pb.Next() {
			_ = n.Swap([2]uint64{i, j})
		}
	})
}

func BenchmarkAdd(b *testing.B) {
	n := &Uint128{}
	runBenchmarks(b, func(pb *testing.PB) {
		i, j := rand.Uint64(), rand.Uint64()
		for pb.Next() {
			_ = n.Add([2]uint64{i, j})
		}
	})
}

func BenchmarkAnd(b *testing.B) {
	n := &Uint128{}
	runBenchmarks(b, func(pb *testing.PB) {
		i, j := rand.Uint64(), rand.Uint64()
		for pb.Next() {
			_ = n.And([2]uint64{i, j})
		}
	})
}

func BenchmarkOr(b *testing.B) {
	n := &Uint128{}
	runBenchmarks(b, func(pb *testing.PB) {
		i, j := rand.Uint64(), rand.Uint64()
		for pb.Next() {
			_ = n.Or([2]uint64{i, j})
		}
	})
}

func BenchmarkXor(b *testing.B) {
	n := &Uint128{}
	runBenchmarks(b, func(pb *testing.PB) {
		i, j := rand.Uint64(), rand.Uint64()
		for pb.Next() {
			_ = n.Xor([2]uint64{i, j})
		}
	})
}

func BenchmarkCAS(b *testing.B) {
	n := &Uint128{}
	_i, _j := rand.Uint64(), rand.Uint64()
	runBenchmarks(b, func(pb *testing.PB) {
		i, j := _i, _j
		for pb.Next() {
			_ = n.CompareAndSwap([2]uint64{i, j}, [2]uint64{j, i})
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

