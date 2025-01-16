package main

import "testing"

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add()
	}
}

func BenchmarkAddConcurrentNoPointer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AddConcurrentNoPointer()
	}
}

func BenchmarkAddConcurrentPointer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var res int
		AddConcurrentPointer(&res)
	}
}
