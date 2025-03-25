package colors

import (
	"fmt"
)

// Color codes for terminal output
const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Purple = "\033[35m"
	Cyan   = "\033[36m"
)

// Colorize wraps the given text with the specified color code.
func Colorize(text, color string) string {
	return fmt.Sprintf("%s%s%s", color, text, Reset)
}
