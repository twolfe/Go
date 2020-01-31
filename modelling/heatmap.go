package main

import (
  //"math"
  "log"
  //"image/color"
	//"math/rand"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	//"gonum.org/v1/plot/vg/draw"

)

func main() {
  p, err := plot.New()
  if err != nil {
    log.Panic(err)
  }

plotter.DefaultLineStyle.Width = vg.Points(1)
plotter.DefaultGlyphStyle.Radius = vg.Points(3)

p.Y.Tick.Marker = plot.ConstantTicks([]plot.Tick{
    {0, "0"}, {0.25, ""}, {0.5, "0.5"}, {0.75, ""}, {1, "1"},
})
p.X.Tick.Marker = plot.ConstantTicks([]plot.Tick{
    {0, "0"}, {0.25, ""}, {0.5, "0.5"}, {0.75, ""}, {1, "1"},
})

pts := plotter.XYs{{0, 0}, {0, 1}, {0.5, 1}, {0.5, 0.6}, {0, 0.6}}
line, err := plotter.NewLine(pts)
if err != nil {
    log.Panic(err)
}
scatter, err := plotter.NewScatter(pts)
if err != nil {
    log.Panic(err)
}
p.Add(line, scatter)

pts = plotter.XYs{{1, 0}, {0.75, 0}, {0.75, 0.75}}
line, err = plotter.NewLine(pts)
if err != nil {
    log.Panic(err)
}
scatter, err = plotter.NewScatter(pts)
if err != nil {
    log.Panic(err)
}
p.Add(line, scatter)

pts = plotter.XYs{{0.5, 0.5}, {1, 0.5}}
line, err = plotter.NewLine(pts)
if err != nil {
    log.Panic(err)
}
scatter, err = plotter.NewScatter(pts)
if err != nil {
    log.Panic(err)
}
p.Add(line, scatter)

err = p.Save(100, 100, "plotLogo.png")
if err != nil {
    log.Panic(err)
}
}
