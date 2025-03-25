package socials

import (
	"fmt"

	"github.com/ppriyankuu/echome/pkg/colors"
)

// Social represents a single social media link.
type Social struct {
	Name string
	URL  string
}

// GetSocials returns a list of your social media profiles.
func GetSocials() []Social {
	return []Social{
		{"Portfolio Website", "https://ppriyankuu.vercel.app"},
		{"GitHub", "https://github.com/ppriyankuu"},
		{"LinkedIn", "https://linkedin.com/in/priyanku-gogoi"},
		{"Instagram", "https://instagram.com/ppriyankuu_"},
	}
}

var colorSequence = []string{
	colors.Green,
	colors.Red,
	colors.Yellow,
	colors.Blue,
	colors.Purple,
	colors.Cyan,
}

// ColorizeName colors each letter differently
func ColorizeName(name string) string {
	var result string
	for i, char := range name {
		color := colorSequence[i%len(colorSequence)]
		result += colors.Colorize(string(char), color)
	}
	return result
}

// DisplaySocials prints all social media links in a formatted way.
func DisplaySocials() {
	for _, social := range GetSocials() {
		fmt.Printf("- %s: %s\n", ColorizeName(social.Name), social.URL)
	}
}
