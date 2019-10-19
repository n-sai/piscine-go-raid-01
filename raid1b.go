package student

import "github.com/01-edu/z01"

func Raid1b(width, height int) {
	if width < 0 || height < 0 {
		return
	}
	for i := 0; i < height; i++ {
		if i == 0 {
			for j := 0; j < width; j++ {
				if j == 0 {
					z01.PrintRune(47)
				} else if j == (width - 1) {
					z01.PrintRune(92)
				} else {
					z01.PrintRune('*')
				}
			}
			z01.PrintRune(10)
		} else if i == height-1 {
			for j := 0; j < width; j++ {
				if j == 0 {
					z01.PrintRune(92)
				} else if j == width-1 {
					z01.PrintRune(47)
				} else {
					z01.PrintRune('*')
				}
			}
			z01.PrintRune(10)
		} else {
			for j := 0; j < width; j++ {
				if j == 0 || j == width-1 {
					z01.PrintRune('*')
				} else {
					z01.PrintRune(' ')
				}
			}
			z01.PrintRune(10)
		}
	}
}
