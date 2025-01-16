package benchmark

import "sync"

func AddConcurrentPointer(res *int) *int {
	var wg sync.WaitGroup
	var mu sync.Mutex
	wg.Add(1)

	go func() {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			mu.Lock()
			*res += i
			mu.Unlock()
		}
	}()
	wg.Wait()
	return res
}

func AddConcurrentNoPointer() int {
	result := 0
	var wg sync.WaitGroup
	var mu sync.Mutex

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			mu.Lock()
			result += i
			mu.Unlock()
		}
	}()
	wg.Wait()
	return result
}

func Add() int {
	res := 0
	for i := 0; i < 100; i++ {
		res += i
	}
	return res
}
