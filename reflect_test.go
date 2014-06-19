package main

import (
	`testing`
	`reflect`
)

func BenchmarkReflectValue(b *testing.B) {
	p := Point{x: 1, y:2}
	b.ResetTimer()
  for t := 0; t < b.N; t++ {
		reflect.ValueOf(&p)
	}
}

func BenchmarkReflectFunctionByName(b *testing.B) {
	p  := Point{x: 1, y:2}
	rp := reflect.ValueOf(&p)
	b.ResetTimer()
  for t := 0; t < b.N; t++ {
		rp.MethodByName(`Nop`)
	}
}

func BenchmarkReflectCallFunctionEmptyParams(b *testing.B) {
	p  := Point{x: 1, y:2}
	rp := reflect.ValueOf(&p)
	rf := rp.MethodByName(`Nop`)
	params := []reflect.Value{}
	b.ResetTimer()
  for t := 0; t < b.N; t++ {
		rf.Call(params)
	}
}

func BenchmarkReflectCallFunctionWithParams(b *testing.B) {
	p  := Point{x: 1, y:2}
	rp := reflect.ValueOf(&p)
	rf := rp.MethodByName(`NopArgs`)
	params := []reflect.Value{reflect.ValueOf(1), reflect.ValueOf(2), reflect.ValueOf(3)}
	b.ResetTimer()
  for t := 0; t < b.N; t++ {
		rf.Call(params)
	}
}
