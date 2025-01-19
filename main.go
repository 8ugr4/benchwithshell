package main

import (
	"benchwithshell/internal"
	"fmt"
)

func main() {
	parseResults()
}

func parseResults() {
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
