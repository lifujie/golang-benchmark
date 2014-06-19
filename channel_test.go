package main

import (
	`testing`
)

func BenchmarkChannel(b *testing.B) {
	items := make(chan int, 1)
	items <- 0
	b.ResetTimer()
  for t := 0; t < b.N; t++ {
		items <- <- items
	}
}
