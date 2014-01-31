package ansi

import (
	"fmt"
)

// TODO: "default" color

const (
	black = iota
	red
	green
	yellow
	blue
	magenta
	cyan
	white

	normalFG = 30
	highFG   = 90
	normalBG = 40
	highBG   = 100

	esc = "\033["
	bold = "1;"
	no_bold = "21;"
	underline = "4;"
	no_underline = "24;"
	blink = "5;"
	no_blink = "25;"
	inverse = "7;"
	no_inverse = "27;"
	Reset = "\033[0m"
)

var (
	plain = false
	colors = map[string]int{
		"default": 9,
		"black":   black,
		"red":     red,
		"green":   green,
		"yellow":  yellow,
		"blue":    blue,
		"magenta": magenta,
		"cyan":    cyan,
		"white":   white,
	}
)


// Gets ANSI escape code for a color style
func ColorCode(style string) string {
	if plain || style == "" {
		return ""
	}
	if style == "reset" {
		return Reset
	}

	var vars [6]string
	baseVar := 0
	workingVar := 0
	pos := 0
	var i int
	var c rune
	for i, c = range style {
		if c == '+' || c == '-' || c == ':' {
			vars[workingVar] += style[pos:i]
			pos = i + 1 // skip +, -, : char
			switch c {
			case '+':
				workingVar = baseVar + 1
			case '-':
				workingVar = baseVar + 2
			case ':':
				baseVar = 3
				workingVar = baseVar
			}
		}
	}
	// include last part
	vars[workingVar] = style[pos:i+1]

	fgColor, fgStyle, fgNoStyle := vars[0], vars[1], vars[2]
	bgColor, bgStyle, _         := vars[3], vars[4], vars[5]

	fg := -1
	if tmp, ok := colors[fgColor]; ok {
		fg = tmp
	}
	bg := -1
	if tmp, ok := colors[bgColor]; ok {
		bg = tmp
	}

	code := esc
	base := normalFG
	for _, c := range fgStyle {
		switch c {
		case 'b':
			code += bold
		case 'B':
			code += blink
		case 'u':
			code += underline
		case 'i':
			code += inverse
		case 'h':
			base = highFG
		}
	}
	for _, c := range fgNoStyle {
		switch c {
		case 'b':
			code += no_bold
		case 'B':
			code += no_blink
		case 'u':
			code += no_underline
		case 'i':
			code += no_inverse
		}
	}
	if fg > -1 {
		code += fmt.Sprintf("%d;", base+fg)
	}

	base = normalBG
	for _, c := range bgStyle {
		switch c {
		case 'h':
			base = highBG
		}
	}
	if bg > -1 {
		code += fmt.Sprintf("%d;", base+bg)
	}

	// remove last ';'
	return code[:len(code)-1] + "m"
}

// Surround string with ANSI color and reset codes
func Color(s, style string) string {
	if plain || style == "" {
		return s
	}
	return ColorCode(style) + s + Reset
}

// Creates a closure for surrounding a string with ANSI color and reset codes
func ColorFunc(style string) func(string) string {
	if style == "" {
		return func(s string) string {
			return s
		}
	}
	code := ColorCode(style)
	return func(s string) string {
		if plain || s == "" {
			return s
		}
		return code + s + Reset
	}
}

// Disables ANSI color codes. Returns the previous setting
func DisableColors(disable bool) (ret bool) {
	ret = plain
	plain = disable
	return
}
