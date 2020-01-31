// Copyright 2016 The Gosl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

package main

import (
	//"math"

	"github.com/cpmech/gosl/io"
	"github.com/cpmech/gosl/la"
	"github.com/cpmech/gosl/ode"
	"github.com/cpmech/gosl/plt"
	"github.com/cpmech/gosl/utl"
)

func main() {

	// title
	io.Pf("Allee Effects")

	// constants
	r := 0.5
	A := 300.0
	K := 2000.0
	dt := 1.875 / 50.0
	tf := 5.0
	atol := 1e-4
	rtol := 1e-4
	numJac := false
	saveStep := true
	saveDense := false

	// ode function: f(x,y) = dy/dx
	/*fcn := func(f la.Vector, dx, x float64, y la.Vector) {
		f[0] = λ*y[0] - λ//*math.Cos(x)
		f[1] = λ*y[0] - λ*y[1]//*math.Sin(x)*y[1]
	}*/

	// ode function: f(x,y) = dy/dx
	fcn1 := func(f la.Vector, dt, t float64, N la.Vector) {
		//f[0] = λ*y[0] - λ*math.Exp(x)
		f[0] = r*N[0]*(N[0]/A - 1)*(1 - N[0]/K)
	}


	// Jacobian function: J = df/dy
	/*jac := func(dfdy *la.Triplet, dx, x float64, y la.Vector) {
		if dfdy.Max() == 0 {
			dfdy.Init(1, 1, 1)
		}
		dfdy.Start()
		dfdy.Put(0, 0, λ)
	}*/
	//var jac = nil

	// initial values
	ndim := 1
	y := la.NewVector(ndim)
	y[0] = 500.0

	// FwEuler
	io.Pf("\n------------ Runge Kutta 2 ------------------\n")
	fixedStep := true
	stat1, out1 := ode.Solve("rk2", fcn1, nil, y.GetCopy(), tf, dt, atol, rtol, numJac, fixedStep, saveStep, saveDense)
	stat1.Print(false)

	A = 1000.0
	K = 2000.0

	fcn2 := func(f la.Vector, dt, t float64, N la.Vector) {
		//f[0] = λ*y[0] - λ*math.Exp(x)
		f[0] = r*N[0]*(N[0]/A - 1)*(1 - N[0]/K)
	}
	stat2, out2 := ode.Solve("rk2", fcn2, nil, y.GetCopy(), tf, dt, atol, rtol, numJac, fixedStep, saveStep, saveDense)
	stat2.Print(false)

	// BwEuler
	/*io.Pf("\n------------ Backward-Euler ------------------\n")
	fixedStep = true
	stat2, out2 := ode.Solve("bweuler", fcn, nil, y.GetCopy(), xf, dx, atol, rtol, numJac, fixedStep, saveStep, saveDense)
	stat2.Print(false)

	// MoEuler
	io.Pf("\n------------ Modified-Euler ------------------\n")
	fixedStep = false
	stat3, out3 := ode.Solve("moeuler", fcn, nil, y.GetCopy(), xf, dx, atol, rtol, numJac, fixedStep, saveStep, saveDense)
	stat3.Print(true)

	// DoPri5
	io.Pf("\n------------ Dormand-Prince5 -----------------\n")
	fixedStep = false
	stat4, out4 := ode.Solve("dopri5", fcn, nil, y.GetCopy(), xf, dx, atol, rtol, numJac, fixedStep, saveStep, saveDense)
	stat4.Print(true)

	// DoPri8
	io.Pf("\n------------ Dormand-Prince8 -----------------\n")
	fixedStep = false
	stat5, out5 := ode.Solve("dopri8", fcn, nil, y.GetCopy(), xf, dx, atol, rtol, numJac, fixedStep, saveStep, saveDense)
	stat5.Print(true)

	// Radau5
	io.Pf("\n------------ Radau5 --------------------------\n")
	fixedStep = false
	stat6, out6 := ode.Solve("radau5", fcn, nil, y.GetCopy(), xf, dx, atol, rtol, numJac, fixedStep, saveStep, saveDense)
	stat6.Print(true)*/

	// analytical solution
	npts := 201
	NN := utl.LinSpace(0, K*1.1, npts)
	yy1 := utl.GetMapped(NN, func(N float64) float64 {
		return r*N*(N/A - 1)*(1 - N/K)
	})
	A = 100.0
	yy2 := utl.GetMapped(NN, func(N float64) float64 {
		return r*N*(N/A - 1)*(1 - N/K)
	})
	//A = 0.000
	yy3 := utl.GetMapped(NN, func(N float64) float64 {
		return r*N*(1 - N/K)
	})



	// plot
	plt.Reset(true, &plt.A{WidthPt: 1000, Prop: 1.1, FszLbl: 10.0, FszXtck: 10.0, FszLeg: 10.0, FszYtck: 10.0})
	plt.Title("AB", &plt.A{})
	plt.Subplot(4, 2, 1)

	//plt.Plot(tt, yy1, &plt.A{C: "grey", Ls: "-", Lw: 5, L: "analytical", NoClip: true})
	plt.Plot(out1.GetStepX(), out1.GetStepY(0), &plt.A{L: "r=0.5, A=300, K=2000", C: "b", Ls: "-"})
  plt.Plot(out2.GetStepX(), out2.GetStepY(0), &plt.A{L: "r=0.5, A=1000, K=2000", C: "r", Ls: "-"})
	plt.Gll("$t$", "$N$", nil)

	plt.Subplot(4, 2, 2)
	plt.Plot(NN, yy1, &plt.A{C: "k", Ls: "-", Lw: 1, L: "Strong Allee Effect (large A)", NoClip: true})
	//plt.Plot(NN, yy2, &plt.A{C: "grey", Ls: "-", Lw: 1, L: "analytical", NoClip: true})
	//plt.Plot(out1.GetStepY(0), out1.GetStepY(1), &plt.A{L: "moeuler", C: "k", M: ".", Ls: ":"})
	//plt.Plot(out1.GetStepY(0), out1.GetStepY(1), &plt.A{L: "dopri5", C: "r", M: ".", Ls: "none"})
	plt.Gll("$N$", "$dN/dt$", nil)
	plt.Subplot(4, 2, 3)
	plt.Plot(NN, yy2, &plt.A{C: "k", Ls: "-", Lw: 1, L: "Weak Allee Effect (small A)", NoClip: true})
	//plt.Plot(out1.GetStepX(), out1.GetStepY(0), &plt.A{L: "prey", C: "k", M: ".", Ls: ":"})
	//plt.Plot(out1.GetStepX(), out1.GetStepY(1), &plt.A{L: "predator", C: "r", M: ".", Ls: "-", Void: true})
	plt.Gll("$N$", "$dN/dt$", nil)

	plt.Subplot(4, 2, 4)
	plt.Plot(NN, yy3, &plt.A{C: "k", Ls: "-", Lw: 1, L: "No Allee Effect = Logistic function", NoClip: true})
	//plt.Plot(out1.GetStepX(), out1.GetStepY(0), &plt.A{L: "prey", C: "k", M: ".", Ls: ":"})
	//plt.Plot(out1.GetStepX(), out1.GetStepY(1), &plt.A{L: "predator", C: "r", M: ".", Ls: "-", Void: true})
	plt.Gll("$N$", "$dN/dt$", nil)

	plt.Save("/tmp/gosl", "allee-effect")
}
