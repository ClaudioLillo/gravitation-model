package main

import (
	"image"
	"image/color"

	"github.com/claudiolillo/gravitation-model/internal/constants"
	"github.com/claudiolillo/gravitation-model/internal/system"
	"github.com/llgcode/draw2d/draw2dimg"
)



func main() {
	// Initialize the graphic context on an RGBA image
	dest := image.NewRGBA(image.Rect(0, 0, 1000, 1000))
	gc := draw2dimg.NewGraphicContext(dest)

	sys := system.New()

	// Here you can set particles providing initial position, mass, speed and a color to trace
	p1 := &system.Particle{X: 250, Y: 250, Vx: 0.5, Vy: -0.4, Color: color.RGBA{255,0,255,255}, Mass: 2 * constants.MT, Key: "1"}
	p2 := &system.Particle{X: 750, Y: 750, Vx: -0.5, Vy: 0.4, Color: color.RGBA{250, 255, 100, 255}, Mass: 2 * constants.MT, Key: "2"}

	// Particles should be added to the system
	sys.AddParticle(p1)
	sys.AddParticle(p2)

	// The system is build
	sys.Build()

	for i := 0; i < 10500; i++ {
		// line width (float)
		lw := 3.0
		for _, value := range sys.Particles {
			gc.SetFillColor(color.RGBA{0x44, 0xff, 0x44, 0xff})
			gc.SetStrokeColor(value.Color)
			gc.SetLineWidth(1)
			gc.MoveTo(value.X, value.Y)
			gc.LineTo(value.X + lw, value.Y)
			gc.LineTo(value.X + lw, value.Y + lw)
			gc.LineTo(value.X, value.Y + lw)
			gc.FillStroke()
		}
		sys.Describe()
		sys.Next()
	}
	
	// Save to file
	draw2dimg.SaveToPngFile("model.png", dest)
}
