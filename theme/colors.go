package theme

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

const (
	primary = uint32(0x0f62fe)
	gray100 = uint32(0x161616)
)

var Primary color.Color = hexToColor(primary)

// https://github.com/carbon-design-system/carbon/blob/v10/packages/colors/src/colors.js
var colorMap = map[fyne.ThemeColorName]uint32{
	theme.ColorNameBackground: gray100,
	theme.ColorNameHover:      gray100,
	theme.ColorNameButton:     primary,
	theme.ColorNamePrimary:    primary,

	// Imported colors from carbon
	fyne.ThemeColorName("ui-background"): 0x000000,
}

func (t *CustomTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	if c, ok := colorMap[name]; ok {
		return hexToColor(c)
	}

	return theme.DefaultTheme().Color(name, variant)
}

// Take a regular hex color value in a 32bit
// number and extract the RGB values
//
// Returns a the corresponding RGB color
func hexToColor(hexColor uint32, alpha ...float32) color.Color {
	// Extract RGB values in reverse from uint32 value
	vals := make([]uint8, 3)

	// Capture 1 byte at a time in a 32 bit mask
	mask := uint32(0xff)
	for i := 0; i < len(vals); i++ {
		// Read intersect values at bit mask,
		// then shift right to get the actual value.
		// Shift the mask a byte left to get the next group.
		vals[i] = uint8((hexColor & mask) >> (i * 8))
		mask = mask << 8
	}

	alphaVal := float32(255)
	if len(alpha) > 0 {
		alphaVal = alphaVal * alpha[0]
	}

	return color.RGBA{
		R: vals[2],
		G: vals[1],
		B: vals[0],
		A: uint8(alphaVal), // Floor value on cast
	}
}
