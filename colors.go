package colors

import "fmt"

const Reset = "\033[0m"

const DefaulT = "\033[39m" // default foreground
const DEFAULT = "\033[49m" // default background
// foreground
const Black = "\033[30m"
const Red = "\033[31m"
const Green = "\033[32m"
const Yellow = "\033[33m"
const Blue = "\033[34m"
const Magenta = "\033[35m"
const Cyan = "\033[36m"
const White = "\033[37m"
const Gray = "\033[90m"
const Red_light = "\033[91m"
const Green_light = "\033[92m"
const Yellow_light = "\033[93m"
const Blue_light = "\033[94m"
const Magenta_light = "\033[95m"
const Cyan_light = "\033[96m"
const Gray_light = "\033[97m"

// background
const BLACK = "\033[40m"
const RED = "\033[41m"
const GREEN = "\033[42m"
const YELLOW = "\033[43m"
const BLUE = "\033[44m"
const MAGENTA = "\033[45m"
const CYAN = "\033[46m"
const WHITE = "\033[47m"
const GRAY = "\033[100m"
const RED_LIGHT = "\033[101m"
const GREEN_LIGHT = "\033[102m"
const YELLOW_LIGHT = "\033[103m"
const BLUE_LIGHT = "\033[104m"
const MAGENTA_LIGHT = "\033[105m"
const CYAN_LIGHT = "\033[106m"
const GRAY_LIGHT = "\033[107m"

// effects
const Bold = "\033[1m"       // bold
const Faint = "\033[2m"      // faint
const Standout = "\033[3m"   // standout
const Underlined = "\033[4m" // underlined
const Blink = "\033[5m"      // blink
const Reverse = "\033[7m"    // reverse

func Println(content any, features ...string) {
	for _, feature := range features {
		print(feature)
	}
	fmt.Print(content)
	println(Reset)
}
func Print(content any, features ...string) {
	for _, feature := range features {
		print(feature)
	}
	fmt.Print(content)
	print(Reset)
}

func Color(features ...string) {
	for _, feature := range features {
		print(feature)
	}
}

func Reseting() { print(Reset) }
