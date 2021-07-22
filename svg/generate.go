package svg

import (
	"fmt"
	"io"
	"strconv"
	"strings"

	svg "github.com/ajstarks/svgo"
)

const (
	DefaultFontFamily  = "Mr+Dafoe"
	DefaultFontSize    = 400 // px
	DefaultText        = "FooBar"
	DefaultFill        = "#08d2a1"
	DefaultStroke      = "#088a6a"
	DefaultStrokeWidth = 32
)

type Config struct {
	BackgroundColor string
	Fill            string
	FontFamily      string
	FontSize        int
	FontStyle       string
	Rotate          int
	Stroke          string
	StrokeWidth     int
	Text            string
}

func NewConfigFromMap(params map[string]string) *Config {
	c := NewConfig()
	if backgroundColor := params["background-color"]; backgroundColor != "" {
		c.BackgroundColor = backgroundColor
	}

	if fill := params["fill"]; fill != "" {
		c.Fill = fill
	}

	if fontFamily := strings.Title(params["font-family"]); fontFamily != "" {
		c.FontFamily = fontFamily
	}

	if fontSize, _ := strconv.Atoi(params["font-size"]); fontSize != 0 {
		c.FontSize = fontSize
	}

	if fontStyle, _ := params["font-style"]; fontStyle != "" {
		c.FontStyle = fontStyle
	}

	if rotate, _ := strconv.Atoi(params["rotate"]); rotate != 0 {
		c.Rotate = rotate
	}

	if stroke := params["stroke"]; stroke != "" {
		c.Stroke = stroke
	}

	if strokeWidth, _ := strconv.Atoi(params["stroke-width"]); strokeWidth != 0 {
		c.StrokeWidth = strokeWidth
	}

	if text := params["text"]; text != "" {
		c.Text = text
	}

	return c
}

func NewConfig() *Config {
	return &Config{
		Fill:        DefaultFill,
		FontFamily:  DefaultFontFamily,
		FontSize:    DefaultFontSize,
		Stroke:      DefaultStroke,
		StrokeWidth: DefaultStrokeWidth,
		Text:        DefaultText,
	}
}

func Generate(c *Config, w io.Writer) {
	var props []string
	if c.BackgroundColor != "" {
		props = append(props, fmt.Sprintf(`style="background-color:%s"`, c.BackgroundColor))
	}

	width := c.FontSize * len(c.Text)
	heigh := float64(c.FontSize) * 1.6

	props = append(props, fmt.Sprintf(`viewBox="0, 0, %v, %v"`, width, heigh))
	props = append(props, `preserveAspectRatio="xMidYMid meet"`)

	textProps := []string{
		`paint-order="stroke fill markers"`,
		`dominant-baseline="middle" text-anchor="middle"`,
		fmt.Sprintf(`font-size="%vpx"`, c.FontSize),
		fmt.Sprintf("font-family=%q", strings.Replace(c.FontFamily, "+", " ", -1)),
		fmt.Sprintf(`fill=%q`, c.Fill),
		fmt.Sprintf(`stroke=%q`, c.Stroke),
		fmt.Sprintf(`stroke-width="%v"`, c.StrokeWidth),
	}

	if c.FontStyle != "" {
		textProps = append(textProps, fmt.Sprintf(`font-style=%q`, c.FontStyle))
	}

	textTransformValues := []string{
		fmt.Sprintf(`translate(%v, %v)`, width/2, heigh/2), // Moves the title to be centered and fit within the viewbox.
	}

	if c.Rotate != 0 {
		textTransformValues = append(textTransformValues, fmt.Sprintf(`rotate(%v)`, c.Rotate))
	}

	textProps = append(textProps, fmt.Sprintf(`transform="%s"`, strings.Join(textTransformValues, ", ")))

	s := svg.New(w)
	s.Startraw(props...)
	s.Style("text/css", fmt.Sprintf(`@import url("https://fonts.googleapis.com/css?family=%s&text=%s");`, c.FontFamily, c.Text))
	s.Text(0, 0, c.Text, textProps...)
	s.End()
}
