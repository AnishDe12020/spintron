package spinner

import (
	"errors"

	"github.com/fatih/color"
)

// errInvalidColor is returned when attempting to set an invalid color
var errInvalidColor = errors.New("invalid color")

// validColors holds an array of the only colors allowed
var validColors = map[string]bool{
	// default colors for backwards compatibility
	"black":   true,
	"red":     true,
	"green":   true,
	"yellow":  true,
	"blue":    true,
	"magenta": true,
	"cyan":    true,
	"white":   true,

	// attributes
	"reset":        true,
	"bold":         true,
	"faint":        true,
	"italic":       true,
	"underline":    true,
	"blinkslow":    true,
	"blinkrapid":   true,
	"reversevideo": true,
	"concealed":    true,
	"crossedout":   true,

	// foreground text
	"fgBlack":   true,
	"fgRed":     true,
	"fgGreen":   true,
	"fgYellow":  true,
	"fgBlue":    true,
	"fgMagenta": true,
	"fgCyan":    true,
	"fgWhite":   true,

	// foreground Hi-Intensity text
	"fgHiBlack":   true,
	"fgHiRed":     true,
	"fgHiGreen":   true,
	"fgHiYellow":  true,
	"fgHiBlue":    true,
	"fgHiMagenta": true,
	"fgHiCyan":    true,
	"fgHiWhite":   true,

	// background text
	"bgBlack":   true,
	"bgRed":     true,
	"bgGreen":   true,
	"bgYellow":  true,
	"bgBlue":    true,
	"bgMagenta": true,
	"bgCyan":    true,
	"bgWhite":   true,

	// background Hi-Intensity text
	"bgHiBlack":   true,
	"bgHiRed":     true,
	"bgHiGreen":   true,
	"bgHiYellow":  true,
	"bgHiBlue":    true,
	"bgHiMagenta": true,
	"bgHiCyan":    true,
	"bgHiWhite":   true,
}

// returns a valid color's foreground text color attribute
var colorAttributeMap = map[string]color.Attribute{
	// default colors for backwards compatibility
	"black":   color.FgBlack,
	"red":     color.FgRed,
	"green":   color.FgGreen,
	"yellow":  color.FgYellow,
	"blue":    color.FgBlue,
	"magenta": color.FgMagenta,
	"cyan":    color.FgCyan,
	"white":   color.FgWhite,

	// attributes
	"reset":        color.Reset,
	"bold":         color.Bold,
	"faint":        color.Faint,
	"italic":       color.Italic,
	"underline":    color.Underline,
	"blinkslow":    color.BlinkSlow,
	"blinkrapid":   color.BlinkRapid,
	"reversevideo": color.ReverseVideo,
	"concealed":    color.Concealed,
	"crossedout":   color.CrossedOut,

	// foreground text colors
	"fgBlack":   color.FgBlack,
	"fgRed":     color.FgRed,
	"fgGreen":   color.FgGreen,
	"fgYellow":  color.FgYellow,
	"fgBlue":    color.FgBlue,
	"fgMagenta": color.FgMagenta,
	"fgCyan":    color.FgCyan,
	"fgWhite":   color.FgWhite,

	// foreground Hi-Intensity text colors
	"fgHiBlack":   color.FgHiBlack,
	"fgHiRed":     color.FgHiRed,
	"fgHiGreen":   color.FgHiGreen,
	"fgHiYellow":  color.FgHiYellow,
	"fgHiBlue":    color.FgHiBlue,
	"fgHiMagenta": color.FgHiMagenta,
	"fgHiCyan":    color.FgHiCyan,
	"fgHiWhite":   color.FgHiWhite,

	// background text colors
	"bgBlack":   color.BgBlack,
	"bgRed":     color.BgRed,
	"bgGreen":   color.BgGreen,
	"bgYellow":  color.BgYellow,
	"bgBlue":    color.BgBlue,
	"bgMagenta": color.BgMagenta,
	"bgCyan":    color.BgCyan,
	"bgWhite":   color.BgWhite,

	// background Hi-Intensity text colors
	"bgHiBlack":   color.BgHiBlack,
	"bgHiRed":     color.BgHiRed,
	"bgHiGreen":   color.BgHiGreen,
	"bgHiYellow":  color.BgHiYellow,
	"bgHiBlue":    color.BgHiBlue,
	"bgHiMagenta": color.BgHiMagenta,
	"bgHiCyan":    color.BgHiCyan,
	"bgHiWhite":   color.BgHiWhite,
}

// validColor will make sure the given color is actually allowed.
func validColor(c string) bool {
	return validColors[c]
}
