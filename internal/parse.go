package internal

import (
	"io"
	"os"
	"strconv"
	"strings"
)

const (
	testFile = "test.out"
)

type Results struct {
	Name           []string
	IterationCount []int64
	Runtime        []float64
}

func ParseResults() *Results {
	r := Results{}
	file, err := os.Open(testFile)
	if err != nil {
		panic(err)
	}
	body, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(body), "\n")
	for _, line := range lines {
		allWords := strings.Split(line, "\t")
		for i, word := range allWords {
			if strings.HasPrefix(word, "Benchmark") {
				r.Name = append(r.Name, allWords[i])
				convInt, err := strconv.ParseInt(strings.TrimSpace(allWords[i+1]), 10, 64)
				if err != nil {
					panic(err)
				}
				r.IterationCount = append(r.IterationCount, convInt)
				trimmed, ok := strings.CutSuffix(strings.TrimSpace(allWords[i+2]), " ns/op")
				if !ok {
					panic(ok)
				}
				convFloat, err := strconv.ParseFloat(trimmed, 64)
				if err != nil {
					panic(err)
				}
				r.Runtime = append(r.Runtime, convFloat)
			}
		}
	}
	return &r
}
