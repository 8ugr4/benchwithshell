package sorting

import "math/rand/v2"

var inputList []int

func createInput() {
	inputList = rand.Perm(50)
}
