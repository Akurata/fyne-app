package theme

import (
	"fyne-app/assets"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

var _ fyne.Theme = (*CustomTheme)(nil)

type CustomTheme struct {
	fyne.Theme
}

func (t *CustomTheme) Font(s fyne.TextStyle) fyne.Resource {
	switch {
	case s.Bold && s.Italic:
		return assets.AssetIBMPlexSansBoldItalicTtf
	case s.Bold:
		return assets.AssetIBMPlexSansBoldTtf
	case s.Italic:
		return assets.AssetIBMPlexSansItalicTtf
	}

	return assets.AssetIBMPlexSansRegularTtf
}

func (t *CustomTheme) Icon(name fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(name)
}

func (t *CustomTheme) Size(name fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(name)
}
