package main

import (
    //"unicode/utf8"
  	"image/color"
    //"math"
    "math/rand"
    //"fmt"
    "gonum.org/v1/gonum/stat/distuv"
    "gonum.org/v1/plot"
    "gonum.org/v1/plot/plotter"
    "gonum.org/v1/plot/vg"
)

const (
  //male advantage
  m = 5.

  //For nA and na. We must also choose values for N, B and b
  N = 10.
  B = 20.
  b = 1
)

func RangeInt(min float64, max float64, n int) []float64 {
     arr := make([]float64, n)
     for r := 0; r < n; r++ {
         arr[r] = min + rand.Float64() * (max - min)
     }
     return arr
}


func sexConflict(x float64) float64 {
  y := m*x/(m*x+1-x)
  pAA := x*y
  pAa := x*(1-y)
  paA := (1-x)*y
  paa := (1-x)*(1-y)
  nA := N*(pAA*b*B + 0.5*pAa*B + 0.5*paA*b*B)
  na := N*(0.5*pAa*B + 0.5*paA*b*B + paa*B)
  xnew := nA/(nA+na)
  return xnew
}

func main()  {

  //Plot
  p, err := plot.New()
  if err != nil {
    panic(err)
  }
  p.Title.Text = "Functions"
  p.X.Label.Text = "Current frequency of A allele, x"
  p.Y.Label.Text = "Future frequency of A, x'"

  n := plotter.NewFunction(sexConflict)
  //n.Dashes = []vg.Length{vg.Points(2), vg.Points(2)}
  n.Width = vg.Points(1)
  n.Color = color.RGBA{R: 255, A: 255}

  //var DistriGamma = distuv.Gamma{Alpha: 0.2, Beta: 0.4}
  //g := plotter.NewFunction(DistriGamma.Prob)

  var DistriLaplace = distuv.Laplace{Mu: 0.5, Scale: 0.8}
  l := plotter.NewFunction(DistriLaplace.Prob)
  l.Width = vg.Points(1)
  l.Color = color.RGBA{G: 255, A: 255}
  //var DistriGamma = distuv.Gamma{Alpha: 0.2, Beta: 0.4}


  e := plotter.NewFunction(func (x float64) float64 { return x })
  e.Dashes = []vg.Length{vg.Points(2), vg.Points(2)}
  e.Width = vg.Points(1)

  p.Add(n, e, l)
	p.Legend.Add("sex-conflict", n)
  p.Legend.Add("Laplace(mu:0.5, scale:0.8)", l)
	p.Legend.ThumbnailWidth = 0.5 * vg.Inch

  p.X.Min = 0
  p.X.Max = 1
  p.Y.Min = 0
  p.Y.Max = 1

  // Save the plot to a PNG file.
  if err := p.Save(4*vg.Inch, 4*vg.Inch, "functions.png"); err != nil {
    panic(err)
  }
}
