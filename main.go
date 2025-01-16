package main

import (
	"benchwithshell/internal"
	"fmt"
)

func main() {
	r := internal.ParseResults()
	v := r.Runtime[0]
	var index int
	for i, _ := range r.Runtime {
		if i+1 != len(r.Runtime) && v > r.Runtime[i+1] {
			v = r.Runtime[i+1]
			index = i
		}
	}
	fmt.Printf("fastest result:\n%s : %.2f ns/op\n", r.Name[index], r.Runtime[index])
}

/*
bench calc

	timeStart := time.Now()
	result := benchmark.Add()
	log.Printf("Add runtime :%v,result:%d", time.Since(timeStart).Microseconds(), result)

	timeStart2 := time.Now()
	result = 0
	benchmark.AddConcurrentPointer(&result)
	log.Printf("AddConcurrentPointer runtime:%v,result:%d", time.Since(timeStart2).Microseconds(), result)

	timeStart3 := time.Now()
	res := benchmark.AddConcurrentNoPointer()
	log.Printf("AddConcurrentNoPointer runtime:%v, result:%d", time.Since(timeStart3).Microseconds(), res)

*/
