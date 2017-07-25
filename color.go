package color

import (
	"fmt"
)

var baseColors = map[string]string{
	"black":   "\033[30m",
	"red":     "\033[31m",
	"green":   "\033[32m",
	"yellow":  "\033[33m",
	"blue":    "\033[34m",
	"magenta": "\033[35m",
	"cyan":    "\033[36m",
	"white":   "\033[37m",
}

// append to the end of string
const end = "\033[0m"

// C function return a colored string or magenta if color not found
//
// fmt.Printf(color.C("blue","I'm blue!"))
//
// colors: black, red, green, yellow, blue, magenta(default), cyan, white
func C(color string, msg ...interface{}) string {
	for c, v := range baseColors {
		if c == color {
			return fmt.Sprint(v + fmt.Sprint(msg...) + end)
		}
	}
	// return magenta if no color found
	return fmt.Sprint(baseColors["magenta"] + fmt.Sprint(msg...) + end)
}
