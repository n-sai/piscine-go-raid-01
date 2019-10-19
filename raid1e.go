package student

import "github.com/01-edu/z01"

func Raid1e(x, y int) {
	if x == 0 || y == 0 {
		return
	}
	for z := 0; z < y; z++ {
		if z == 0 {
			for c := 0; c < x; c++ {
				if c == 0 {
					z01.PrintRune('A')
				} else if c == x-1 {
					z01.PrintRune('C')
				} else {
					z01.PrintRune('B')
				}

			}
			z01.PrintRune(10)
		} else if z == y-1 {
			for c := 0; c < x; c++ {
				if c == 0 {
					z01.PrintRune('C')
				} else if c == x-1 {
					z01.PrintRune('A')
				} else {
					z01.PrintRune('B')
				}

			}
			z01.PrintRune(10)

		} else {
			for c := 0; c < x; c++ {
				if c == 0 {
					z01.PrintRune('B')
				} else if c == x-1 {
					z01.PrintRune('B')
				} else {
					z01.PrintRune(' ')
				}

			}
			z01.PrintRune('\n')
		}
	}
}
