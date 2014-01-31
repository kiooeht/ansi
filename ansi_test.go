package ansi

import (
	"fmt"
	"testing"
	"strconv"
)

var (
	tmp_colors []string
	attrs      []string = []string{"b", "B", "u", "h", "i"}
)

func init() {
	tmp_colors = make([]string, len(colors))
	for c, i := range colors {
		if i >= len(colors) {
			i = -1
		}
		fmt.Println(i, c)
		tmp_colors[i+1] = c
	}
	fmt.Println(tmp_colors)
}

func TestColors(t *testing.T) {
	for _, fg := range tmp_colors {
		for _, bg := range tmp_colors {
			fmt.Printf("%stest%s", ColorCode(fmt.Sprintf("%s:%s", fg, bg)), Reset)
		}
		fmt.Println()
	}
}

func TestFGAttrs(t *testing.T) {
	for _, fg := range tmp_colors {
		fmt.Printf("%stest%s", ColorCode(fg), Reset)
		for _, attr := range attrs {
			fmt.Printf("%stest%s", ColorCode(fmt.Sprintf("%s+%s", fg, attr)), Reset)
		}
		fmt.Println()
	}
}

func TestBGAttrs(t *testing.T) {
	for _, bg := range tmp_colors {
		fmt.Printf("%stest%s", ColorCode(":"+bg), Reset)
		fmt.Printf("%stest%s", ColorCode(":"+bg+"+h"), Reset)
		fmt.Println()
	}
}

func ExampleColorCode() {
	str := fmt.Sprintf("%stest%stest%s", ColorCode("magenta+B"), ColorCode("default-B+i:cyan"), Reset)
	fmt.Println(strconv.Quote(str))
	// Output:
	// "\x1b[5;35mtest\x1b[7;25;39;46mtest\x1b[0m"
}


func BenchmarkColorCode1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ColorCode("red:green")
	}
}

func BenchmarkColor(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Color("asdf", "red:green")
	}
}

func BenchmarkColorFunc(b *testing.B) {
	f := ColorFunc("red:green")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		f("asdf")
	}
}
