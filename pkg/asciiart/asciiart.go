package asciiart

import (
	"fmt"
	"math/rand"

	"github.com/common-nighthawk/go-figure"
)

// ANSI color codes
var colors = []string{
	"\033[31m", // Red
	"\033[32m", // Green
	"\033[33m", // Yellow
	"\033[34m", // Blue
	"\033[35m", // Magenta
	"\033[36m", // Cyan
	"\033[37m", // White
	"\033[91m", // Bright Red
	"\033[92m", // Bright Green
	"\033[93m", // Bright Yellow
	"\033[94m", // Bright Blue
	"\033[95m", // Bright Magenta
	"\033[96m", // Bright Cyan
}

// Reset color
const reset = "\033[0m"

// GenerateColorfulName generates and returns colorful ASCII art for the given name
func GenerateColorfulName(name string) string {
	// Generate ASCII art using go-figure
	fig := figure.NewFigure(name, "", true)
	asciiArt := fig.String()

	// Colorize the ASCII art
	result := ""
	for _, char := range asciiArt {
		if char == ' ' || char == '\n' {
			result += string(char)
			continue
		}
		// Random color for each character
		color := colors[rand.Intn(len(colors))]
		result += fmt.Sprintf("%s%c%s", color, char, reset)
	}
	return result
}
