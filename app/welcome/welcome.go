package welcome

import (
	"github.com/fatih/color"
	"github.com/getwe/figlet4go"
	"flag"
	"fmt"
)

func SayHello () {
	flag_str := flag.String("str", "Fresh Ego Kid", "input string")

	str := * flag_str
	ascii := figlet4go.NewAsciiRender()
 
	// Most simple usage
	render_str, _ := ascii.Render(str)
	fmt.Println(render_str)
 
	// Change the font color
	colors := [...] color.Attribute {
	 color.FgMagenta,
	 color.FgYellow,
	 color.FgBlue,
	 color.FgCyan,
	 color.FgRed,
	 color.FgWhite,
	}
 
	options := figlet4go.NewRenderOptions()
	options.FontColor = make([]color.Attribute, len(str))
 
	for i := range options.FontColor {
	 options.FontColor[i] = colors[i % len(colors)]
	}
 
	render_str, _ = ascii.RenderOpts(str, options)
	fmt.Println(render_str)
}