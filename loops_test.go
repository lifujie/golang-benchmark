package main

import (
	`testing`
)

func BenchmarkLoopSeek(b *testing.B) {
	items := make([]int, 1000)
	for i := 0; i < len(items); i++ {
		items[i] = i
	}
	b.ResetTimer()
  for t := 0; t < b.N; t++ {
		for i := 0; i < len(items); i++ {
			if -1 == items[i] {
				break
			}
		}
	}
}

func BenchmarkLoopSeekStaticLength(b *testing.B) {
	items  := make([]int, 1000)
	for i := 0; i < len(items); i++ {
		items[i] = i
	}
	b.ResetTimer()
	length := len(items)
  for t := 0; t < b.N; t++ {
		for i := 0; i < length; i++ {
			if -1 == items[i] {
				break
			}
		}
	}
}

func BenchmarkLoopSeekRange(b *testing.B) {
	items := make([]int, 1000)
	for i := 0; i < len(items); i++ {
		items[i] = i
	}
	b.ResetTimer()
  for t := 0; t < b.N; t++ {
		for _, i := range items {
			if -1 == i {
				break
			}
		}
	}
}
