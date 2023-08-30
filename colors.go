package colors

import "fmt"

const Reset = "\033[0m"

var Frontground_colors = map[string]string{
	"black":   "\033[30m",
	"red":     "\033[31m",
	"green":   "\033[32m",
	"yellow":  "\033[33m",
	"blue":    "\033[34m",
	"magenta": "\033[35m",
	"cyan":    "\033[36m",
	"white":   "\033[37m",
	"gray":    "\033[90m",
	// light colors
	"red_light":     "\033[91m",
	"green_light":   "\033[92m",
	"yellow_light":  "\033[93m",
	"blue_light":    "\033[94m",
	"magenta_light": "\033[95m",
	"cyan_light":    "\033[96m",
	"gray_light":    "\033[97m",
}
var FGC = Frontground_colors
var Background_colors = map[string]string{
	"BLACK":   "\033[40m",
	"RED":     "\033[41m",
	"GREEN":   "\033[42m",
	"YELLOW":  "\033[43m",
	"BLUE":    "\033[44m",
	"MAGENTA": "\033[45m",
	"CYAN":    "\033[46m",
	"WHITE":   "\033[47m",
	"GRAY":    "\033[100m",
	// light colors
	"RED_LIGHT":     "\033[101m",
	"GREEN_LIGHT":   "\033[102m",
	"YELLOW_LIGHT":  "\033[103m",
	"BLUE_LIGHT":    "\033[104m",
	"MAGENTA_LIGHT": "\033[105m",
	"CYAN_LIGHT":    "\033[106m",
	"GRAY_LIGHT":    "\033[107m",
}
var BGC = Background_colors
var Effects = map[string]string{
	"bold":       "\033[1m", // bold
	"faint":      "\033[2m", // faint
	"standout":   "\033[3m", // standout
	"underlined": "\033[4m", // underlined
	"blinkk":     "\033[5m", // blink
	"reverse":    "\033[7m", // reverse
}

var All = map[string]string{
	"off":     "\033[0m",  // off
	"defaulT": "\033[39m", // default foreground
	"DEFAULT": "\033[49m", // default background

	"bold":        "\033[1m",  // bold
	"faint":       "\033[2m",  // faint
	"standoutt":   "\033[3m",  // standout
	"underlinedl": "\033[4m",  // underlined
	"blinkk":      "\033[5m",  // blink
	"reversev":    "\033[7m",  // reverse
	"hidden":      "\033[8m",  // hidden
	"nost":        "\033[23m", // no standout
	"noul":        "\033[24m", // no underlined
	"nobk":        "\033[25m", // no blink
	"norv":        "\033[27m", // no reverse
	// colors
	"black":   "\033[30m",
	"BLACK":   "\033[40m",
	"red":     "\033[31m",
	"RED":     "\033[41m",
	"green":   "\033[32m",
	"GREEN":   "\033[42m",
	"yellow":  "\033[33m",
	"YELLOW":  "\033[43m",
	"blue":    "\033[34m",
	"BLUE":    "\033[44m",
	"magenta": "\033[35m",
	"MAGENTA": "\033[45m",
	"cyan":    "\033[36m",
	"CYAN":    "\033[46m",
	"white":   "\033[37m",
	"WHITE":   "\033[47m",

	// light colors
	"red_light":     "\033[91m",
	"green_light":   "\033[92m",
	"yellow_light":  "\033[93m",
	"blue_light":    "\033[94m",
	"magenta_light": "\033[95m",
	"cyan_light":    "\033[96m",
	"gray_light":    "\033[97m",
	"RED_LIGHT":     "\033[101m",
	"GREEN_LIGHT":   "\033[102m",
	"YELLOW_LIGHT":  "\033[103m",
	"BLUE_LIGHT":    "\033[104m",
	"MAGENTA_LIGHT": "\033[105m",
	"CYAN_LIGHT":    "\033[106m",
	"GRAY_LIGHT":    "\033[107m",
}

func Println(content any, features ...string) {
	for _, feature := range features {
		if code, ok := All[feature]; ok {
			fmt.Print(code)
		}
	}
	fmt.Print(content)
	fmt.Println(Reset)
}
func Color_scope(color string) {
	fmt.Print(color)
	defer fmt.Print(Reset)
}

func Zeroize() { fmt.Print(Reset) }

type Colorable string

func (text Colorable) Print(features ...string) {
	Println(text, features...)
}
