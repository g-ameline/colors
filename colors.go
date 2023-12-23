package colors

import "fmt"

const reset = "\033[0m"

const defaulT = "\033[39m" // default foreground
const DEFAULT = "\033[49m" // default background
// foreground
const black = "\033[30m"
const red = "\033[31m"
const green = "\033[32m"
const yellow = "\033[33m"
const blue = "\033[34m"
const magenta = "\033[35m"
const cyan = "\033[36m"
const white = "\033[37m"
const gray = "\033[90m"
const red_light = "\033[91m"
const green_light = "\033[92m"
const yellow_light = "\033[93m"
const blue_light = "\033[94m"
const magenta_light = "\033[95m"
const cyan_light = "\033[96m"
const gray_light = "\033[97m"

// background
const bg_black = "\033[40m"
const bg_red = "\033[41m"
const bg_green = "\033[42m"
const bg_yellow = "\033[43m"
const bg_blue = "\033[44m"
const bg_magenta = "\033[45m"
const bg_cyan = "\033[46m"
const bg_white = "\033[47m"
const bg_gray = "\033[100m"
const bg_red_light = "\033[101m"
const bg_green_light = "\033[102m"
const bg_yellow_light = "\033[103m"
const bg_blue_light = "\033[104m"
const bg_magenta_light = "\033[105m"
const bg_cyan_light = "\033[106m"
const bg_gray_light = "\033[107m"

// effects
const bold = "\033[1m"       // bold
const faint = "\033[2m"      // faint
const standout = "\033[3m"   // standout
const underlined = "\033[4m" // underlined
const blink = "\033[5m"      // blink
const reverse = "\033[7m"    // reverse

func Red(things ...any) {
	fmt.Print(red)
	fmt.Println(things...)
	fmt.Print(reset)
}
func Yellow(things ...any) {
	fmt.Print(yellow)
	fmt.Println(things...)
	fmt.Print(reset)
}
func Green(things ...any) {
	fmt.Print(green)
	fmt.Println(things...)
	fmt.Print(reset)
}
func Blue(things ...any) {
	fmt.Print(blue)
	fmt.Println(things...)
	fmt.Print(reset)
}
func Magenta(things ...any) {
	fmt.Print(magenta)
	fmt.Println(things...)
	fmt.Print(reset)
}
func Cyan(things ...any) {
	fmt.Print(cyan)
	fmt.Println(things...)
	fmt.Print(reset)
}
func White(things ...any) {
	fmt.Print(white)
	fmt.Println(things...)
	fmt.Print(reset)
}
func Gray(things ...any) {
	fmt.Print(gray)
	fmt.Println(things...)
	fmt.Print(reset)
}

func Print_map(damap map[string]any) {
	for k, v := range damap {
		fmt.Print(k, magenta)
		fmt.Print(" : ")
		fmt.Printf("%T = ", v)
		fmt.Println(v, cyan)
	}
}
func Print_struct(dastruct any) {
	fmt.Print(yellow)
	fmt.Printf("%+v\n", dastruct)
}

func Indent(level int, stuff ...any) {
	var indentation string = " "
	for i := 0; i < level; i++ {
		indentation += indentation
	}
	fmt.Print(indentation)
	fmt.Println(stuff...)
}
