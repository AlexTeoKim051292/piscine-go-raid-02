package main

import (
	"github.com/01-edu/z01"
)

const n int = 9

func main() {
	info := "Error"
	mssg := ""
	//arguments := os.Args
	arguments := []string{".96.4...1", "1...6...4", "5.481.39.", "..795..43", ".3..8....", "4.5.23.18", ".1.63..59", ".59.7.83.", "..359...7"}
	//arg := []string(arguments[1:])
	lenA := 0
	for index := range arguments {
		lenA = index + 1
	}
	field := make([][]rune, n)
	for i := range field {
		field[i] = make([]rune, n)
	}

	if lenA == n {
		z01.PrintRune('a')

		for index := range field {
			field[index] = []rune(arguments[index])
		}

		steps := 0
		for row := 0; row < n; row++ {
			for _, num := range field[row] {
				if IsEmpty(num) {
					steps++
				}
			}
		}
		stepsRec := make([][][]rune, steps)
		for i := range stepsRec {
			stepsRec[i] = make([][]rune, n)
		}
		posRec := make([][]int, steps)
		for i := range posRec {
			posRec[i] = make([]int, 2)
		}

		step := -1
		startValue := '1'
		for row := 0; row < n; row++ {
			for column, value := range field[row] {
				if IsEmpty(value) && mssg == "" {
					z01.PrintRune('b')
					v := startValue
					for ; v < rune('1'+n); v++ {
						field[row][column] = rune(v)
						//value = rune(v)
						PrintField(field)
						if CheckPos(field, row, column) {
							z01.PrintRune('c')
							step++
							stepsRec[step] = field
							posRec[step] = []int{row, column}
							break
						} else if value == '9' {
							z01.PrintRune('d')
							if step > 0 {
								value = '.'
								step--
								field = stepsRec[step]
								row = posRec[step][0]
								column = posRec[step][1]
								startValue = field[row][column] + 1
								break
							} else {
								z01.PrintRune('f')
								mssg = info
							}
						}
					}

				} else if mssg != "" {
					z01.PrintRune('e')
					column = n
					row = n
				}
			}
		}

	} else {
		z01.PrintRune('g')
		mssg = info
	}

	if mssg == "" {
		z01.PrintRune('h')
		PrintField(field)
	} else {
		for _, l := range []rune(mssg) {
			z01.PrintRune(l)
		}
		z01.PrintRune('\n')
	}

}

func PrintField(field [][]rune) {
	for i := range field {
		for _, l := range field[i] {
			z01.PrintRune(l)
		}
		z01.PrintRune('\n')
	}
}

func IsEmpty(r rune) bool {
	if r == '.' {
		return true
	}
	return false
}

func CheckPos(field [][]rune, row, column int) bool {

	value := field[row][column]

	count := 0
	//Check in column
	for i := 0; i < n; i++ {
		if field[i][column] == value {
			count++
		}
	}

	//Check in row
	for j := 0; j < n; j++ {
		if field[row][j] == value {
			count++
		}
	}

	//square 1, 2,3,4,5,6,7,8,9
	for i := (row / 3) * 3; i < (row/3+1)*3; i++ {
		for j := (column / 3) * 3; j < (column/3+1)*3; j++ {
			if field[i][j] == value {
				count++
			}
		}
	}

	if count == 3 {
		return true
	}
	return false
}
