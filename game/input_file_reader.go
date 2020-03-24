package game

import (
	"io/ioutil"
	"strconv"
	"strings"
)

const (
	alive  byte = 'o'
	dead   byte = 'b'
	EOF    byte = '!'
	newRow byte = '$'
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func ReadInputFile(path string) LifeInput {
	dat, err := ioutil.ReadFile(path)
	check(err)

	data := strings.SplitAfterN(string(dat), "\n", 2)
	firstLine := strings.Split(strings.ReplaceAll(data[0], " ", ""), ",")
	cells := data[1]

	x, err := strconv.Atoi(strings.TrimPrefix(firstLine[0], "x="))
	check(err)
	y, err := strconv.Atoi(strings.TrimPrefix(firstLine[1], "y="))
	check(err)

	m := InitMatrix(x, y)
	r := 0
	c := 0
	aux := 0
	for i := 0; i < len(cells); i++ {
		char := cells[i]
		if isNumber(char) { //number
			aux = aux*10 + int((char - '0'))
		} else { //char
			switch char {
			case alive:
				if aux == 0 {
					m[r][c] = 1
					c++
				} else {
					col := 0
					for ; col < aux; col++ {
						m[r][c+col] = 1
					}
					c += col
				}
			case dead:
				if aux == 0 {
					c++
				} else {
					c += aux
				}
			case newRow:
				if aux == 0 {
					r++
				} else {
					r += aux
				}
				c = 0
			case EOF: //do nothing
			case '\n': //do nothing
			default:
				panic("Invalid File format")
			}
			aux = 0
		}

	}
	return LifeInput{
		Cells:  m,
		Width:  x,
		Height: y,
	}
}

func isNumber(b byte) bool {
	return ((b >= '0') && (b <= '9'))
}
