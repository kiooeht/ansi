# ansi

Small library for creating ANSI color and styling codes in Go.

## Example

		import "github.com/kiooeht/ansi"

		// Colorize a string
		foo := ansi.Color("foo", "red+b:white")

		// Create a closure for a code
		green_on_black := ansi.ColorFunc("green:black")
		msg := green_on_black("Matrix style")

		// Keep code for manual use
		code := ansi.ColorCode("green:black")

		msg := code + "Matrix style" + ansi.Reset

## More Examples

		ColorCode("red+B")     // Red blinking
		ColorCode("cyan+b-B")  // Cyan bold, turn off blinking
		ColorCode("yellow+u")  // Yellow underline
		ColorCode("magenta+i") // Magenta inverse

## Style format

		"forgroundColor+attributes-attributes:backgroundColor+attributes-attributes"

## Colors

* default
* black
* red
* green
* yellow
* blue
* magenta
* cyan
* white

## Attributes

* b = bold (foreground)
* B = blink (foreground)
* u = underline (foreground)
* h = high intensity (foreground, background)
* i = inverse

## MIT License

Copyright (c) 2014 Anthony Moore

See the file LICENSE for copying permission.
