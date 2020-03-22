package main

import (
	"fmt"
	"testing"
)

func TestCountNeighbours(t *testing.T) {
	fmt.Println(countNeighbours(Matrix{
		[]bool{false, false, false},
		[]bool{true, true, true},
		[]bool{false, true, true},
	}, 2, 2))
	//fmt.Println(countNeighbours(Matrix{
	//	[]bool{false, false, false},
	//	[]bool{true, true, true},
	//	[]bool{false, true, true},
	//}, 1, 0))
	//fmt.Println(countNeighbours(Matrix{
	//	[]bool{false, false, false},
	//	[]bool{true, true, true},
	//	[]bool{false, true, true},
	//}, 1, 1))
}
