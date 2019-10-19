package student

import "github.com/01-edu/z01"

func Raid1c(width, height int) {
	if width < 0 || height < 0 {
		return
	}
	for i := 0; i < height; i++ {
		if i == 0 {
			for j := 0; j < width; j++ {
				if j == 0 || j == width-1 {
					z01.PrintRune('A')
				} else {
					z01.PrintRune('B')
				}
			}
			z01.PrintRune(10)
		}
		if i > 0 && i < height-1 {
			for j := 0; j < width; j++ {
				if j == 0 || j == width-1 {
					z01.PrintRune('B')
				} else {
					z01.PrintRune(' ')
				}
			}
			z01.PrintRune(10)
		}
		if i == height-1 {
			for j := 0; j < width; j++ {
				if j == 0 || j == width-1 {
					z01.PrintRune('C')
				} else {
					z01.PrintRune('B')
				}
			}
			z01.PrintRune(10)
		}
	}
}