/*
Small library for creating ANSI color and styling codes

Mostly based off of mgutz's library:
	https://github.com/mgutz/ansi

Style format:
	"forgroundColor+attributes-attributes:backgroundColor+attributes-attributes"

Colors:
	default
	black
	red
	green
	yellow
	blue
	magenta
	cyan
	white

Attributes:
	b = bold (foreground)
	B = blink (foreground)
	u = underline (foreground)
	h = high intensity (foreground, background)
	i = inverse
*/
package ansi

