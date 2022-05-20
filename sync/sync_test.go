package sync_test

import (
	"sync"
	"testing"
)

func BenchmarkSyncMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		m := new(sync.Map)
		m.Store(0, 0)
		m.Load(0)
	}
}

func BenchmarkMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		m := make(map[int]int)
		m[0] = 0
		_ = m[0]
	}
}
