package student

import (
	"github.com/01-edu/z01"
)

func Raid1d(x, y int) {
	if x < 0 || y < 0 {
		return
	}

	for i := 0; i < y; i++ {
		if i == 0 || i == y-1 {
			for j := 0; j < x; j++ {
				if j == 0 {
					z01.PrintRune('A')
				}
				if j < x-2 {
					z01.PrintRune('B')
				}
				if j == x-1 && x > 1 {
					z01.PrintRune('C')
				}
			}
			z01.PrintRune(10)
		} else {
			for j := 0; j < x; j++ {
				if j == 0 || j == x-1 {
					z01.PrintRune('B')
				} else {
					z01.PrintRune(' ')
				}

			}
			z01.PrintRune(10)
		}
	}
}
