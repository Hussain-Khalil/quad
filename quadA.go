package main

import "github.com/01-edu/z01"

func QuadA(x, y int) {
	lastline := y - 1
	lastcolumn := x - 1
	if x < 1 || y < 1 {
		return
	}
	for line := 0; line < y; line++ {
		for column := 0; column < x; column++ {
			if (line == 0 || line == lastline) && (column == 0 || column == lastcolumn) {
				z01.PrintRune('o')
			} else if line == 0 || line == lastline {
				z01.PrintRune('-')
			} else if column == 0 || column == lastcolumn {
				z01.PrintRune('|')
			} else {
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune('\n')
	}
}
