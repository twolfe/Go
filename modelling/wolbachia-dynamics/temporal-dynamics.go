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
	io.Pf("Barton and Turelli: Schraiber and Schreiber model")

	// constants
	d_I := 0.05 // Specifies the death rate of infected individuals
	s_h := 1.0 // Measures the intensity of CI
	s_f := 0.1 // Measure the reduction in fecundity due to Wol. infection
	s_v := 0.1 // Measures the decrease in mean lifetime due to Wol. infection
	p_hat := (s_f + s_v - s_f*s_v)/s_h
	dt := 1.875 / 50.0
	tf := 200.0
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
	fcn := func(f la.Vector, dt, t float64, p la.Vector) {
		//f[0] = λ*y[0] - λ*math.Exp(x)
		f[0] = s_h*d_I*p[0]*(1.0 - p[0])*(p[0] - p_hat)/(1.0 - s_f*p[0] - s_h*p[0]*(1.0 - p[0]))
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
	p1 := la.NewVector(ndim)
	p1[0] = 0.1

	p2 := la.NewVector(ndim)
	p2[0] = 0.4
	// FwEuler
	io.Pf("\n------------ Runge Kutta 2 ------------------\n")
	fixedStep := true
	stat1, out1 := ode.Solve("rk2", fcn, nil, p1.GetCopy(), tf, dt, atol, rtol, numJac, fixedStep, saveStep, saveDense)
	stat1.Print(false)

	stat2, out2 := ode.Solve("rk2", fcn, nil, p2.GetCopy(), tf, dt, atol, rtol, numJac, fixedStep, saveStep, saveDense)
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
	pp := utl.LinSpace(0, 1, npts)
	yy := utl.GetMapped(pp, func(p float64) float64 {
		return s_h*d_I*p*(1.0 - p)*(p - p_hat)/(1.0 - s_f* - s_h*p*(1.0 - p))
	})



	// plot
	plt.Reset(true, &plt.A{WidthPt: 1000, Prop: 1.1, FszLbl: 10.0, FszXtck: 10.0, FszLeg: 10.0, FszYtck: 10.0})
	plt.Title("AB", &plt.A{})
	plt.Subplot(4, 2, 1)

	//plt.Plot(tt, yy1, &plt.A{C: "grey", Ls: "-", Lw: 5, L: "analytical", NoClip: true})
	plt.Plot(out1.GetStepX(), out1.GetStepY(0), &plt.A{L: "p_0=0.1", C: "r", Ls: "-"})
	plt.Plot(out2.GetStepX(), out2.GetStepY(0), &plt.A{L: "p_0=0.4", C: "b", Ls: "-"})
  //plt.Plot(out2.GetStepX(), out2.GetStepY(0), &plt.A{L: "r=0.5, A=1000, K=2000", C: "r", Ls: "-"})
	plt.Gll("$t$", "$p$", nil)

	plt.Subplot(4, 2, 2)
	plt.Plot(pp, yy, &plt.A{C: "k", Ls: "-", Lw: 1, L: "Strong Allee Effect", NoClip: true})

	//plt.Plot(NN, yy2, &plt.A{C: "grey", Ls: "-", Lw: 1, L: "analytical", NoClip: true})
	//plt.Plot(out1.GetStepY(0), out1.GetStepY(1), &plt.A{L: "moeuler", C: "k", M: ".", Ls: ":"})
	//plt.Plot(out1.GetStepY(0), out1.GetStepY(1), &plt.A{L: "dopri5", C: "r", M: ".", Ls: "none"})
	plt.Gll("$p$", "$dp/dt$", nil)
  /*
	plt.Subplot(4, 2, 3)
	plt.Plot(NN, yy2, &plt.A{C: "k", Ls: "-", Lw: 1, L: "Weak Allee Effect (small A)", NoClip: true})
	//plt.Plot(out1.GetStepX(), out1.GetStepY(0), &plt.A{L: "prey", C: "k", M: ".", Ls: ":"})
	//plt.Plot(out1.GetStepX(), out1.GetStepY(1), &plt.A{L: "predator", C: "r", M: ".", Ls: "-", Void: true})
	plt.Gll("$N$", "$dN/dt$", nil)

	plt.Subplot(4, 2, 4)
	plt.Plot(NN, yy3, &plt.A{C: "k", Ls: "-", Lw: 1, L: "No Allee Effect = Logistic function", NoClip: true})
	//plt.Plot(out1.GetStepX(), out1.GetStepY(0), &plt.A{L: "prey", C: "k", M: ".", Ls: ":"})
	//plt.Plot(out1.GetStepX(), out1.GetStepY(1), &plt.A{L: "predator", C: "r", M: ".", Ls: "-", Void: true})
	plt.Gll("$N$", "$dN/dt$", nil)*/

	plt.Save("/tmp/gosl", "temporal-dynamics")
}
