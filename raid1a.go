package student

import "github.com/01-edu/z01"

func Raid1a(width, height int) {
	for i := 0; i < height; i++ {
		if i == 0 || i == height-1 {
			for j := 0; j < width; j++ {
				if width < 0 || height < 0 {
					return
				} else if j == 0 || j == width-1 {
					z01.PrintRune('o')
				} else {
					z01.PrintRune('-')
				}
			}
			z01.PrintRune(10)
		} else {
			for j := 0; j < width; j++ {
				if j == 0 || j == width-1 {
					z01.PrintRune('|')
				} else {
					z01.PrintRune(' ')
				}
			}
			z01.PrintRune(10)
		}
	}
}
