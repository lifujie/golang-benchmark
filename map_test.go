package main

import (
	`testing`
)

func BenchmarkMapInsert(b *testing.B) {
	items := map[string]int{}
	b.ResetTimer()
  for t := 0; t < b.N; t++ {
		items[`test`] = 1
	}
}

func BenchmarkMapGet(b *testing.B) {
	items := map[string]int{`test`: 1}
	b.ResetTimer()
  for t := 0; t < b.N; t++ {
		_ = items[`test`]
	}
}

func BenchmarkMapGetExists(b *testing.B) {
	items := map[string]int{`test`: 1}
	b.ResetTimer()
  for t := 0; t < b.N; t++ {
		_, _ = items[`test`]
	}
}
