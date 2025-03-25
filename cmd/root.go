package cmd

import (
	"fmt"

	"github.com/ppriyankuu/echome/pkg/asciiart"
	"github.com/ppriyankuu/echome/pkg/colors"
	"github.com/ppriyankuu/echome/pkg/socials"
)

// Execute runs the CLI tool.
func Execute() {
	fmt.Println(colors.Colorize("Hi! My name is..", colors.Red))

	fmt.Println(colors.Colorize(asciiart.GenerateColorfulName("Priyanku Gogoi."), colors.Cyan))

	fmt.Println(colors.Colorize("I am a Software Engineer!", colors.Green))

	fmt.Println(colors.Colorize("\nCurious? Find me on my socials.", colors.Yellow))

	socials.DisplaySocials()
}
