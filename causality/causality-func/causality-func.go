// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 231.

// Pipeline3 demonstrates a finite 3-stage pipeline
// with range, close, and unidirectional channel types.
package main

import (
	"fmt"
	//"os"
	//"runtime/trace"
	"gonum.org/v1/gonum/stat/distuv"
	"math"
)

/*
######## ##    ## ########  ########  ######
   ##     ##  ##  ##     ## ##       ##    ##
   ##      ####   ##     ## ##       ##
   ##       ##    ########  ######    ######
   ##       ##    ##        ##             ##
   ##       ##    ##        ##       ##    ##
   ##       ##    ##        ########  ######
*/

type F func(out chan<- float64, in ...<-chan float64)

/*
	 ######   #######  ##     ## ##    ## ######## ######## ########
	##    ## ##     ## ##     ## ###   ##    ##    ##       ##     ##
	##       ##     ## ##     ## ####  ##    ##    ##       ##     ##
	##       ##     ## ##     ## ## ## ##    ##    ######   ########
	##       ##     ## ##     ## ##  ####    ##    ##       ##   ##
	##    ## ##     ## ##     ## ##   ###    ##    ##       ##    ##
	 ######   #######   #######  ##    ##    ##    ######## ##     ##
*/

func counter(out chan<- float64) {
	dist := distuv.Normal{
		Mu:    9,
		Sigma: 1,
		//Rate: 0.01,
	}
	for x := 0; x <= 10000; x++ {
		out <- dist.Rand()
	}
	close(out)
}

/*
 ######   #######  ##     ##    ###    ########  ######## ########     ##
##    ## ##     ## ##     ##   ## ##   ##     ## ##       ##     ##  ####
##       ##     ## ##     ##  ##   ##  ##     ## ##       ##     ##    ##
 ######  ##     ## ##     ## ##     ## ########  ######   ########     ##
      ## ##  ## ## ##     ## ######### ##   ##   ##       ##   ##      ##
##    ## ##    ##  ##     ## ##     ## ##    ##  ##       ##    ##     ##
 ######   ##### ##  #######  ##     ## ##     ## ######## ##     ##  ######
*/

func squarer1(out chan<- float64, in <-chan float64) {
	/*dist := distuv.Exponential{
		Rate:    -2,
	}*/
	for v := range in {
		//var r float64 = -1.0 * v
		/*dist := distuv.Normal{
			Mu:    r,
			Sigma:	1,
		}*/
		v = 1.0
		out <- v //dist.Rand()
	}
	close(out)
}

/*
 ######   #######  ##     ##    ###    ########  ######## ########   #######
##    ## ##     ## ##     ##   ## ##   ##     ## ##       ##     ## ##     ##
##       ##     ## ##     ##  ##   ##  ##     ## ##       ##     ##        ##
 ######  ##     ## ##     ## ##     ## ########  ######   ########   #######
      ## ##  ## ## ##     ## ######### ##   ##   ##       ##   ##   ##
##    ## ##    ##  ##     ## ##     ## ##    ##  ##       ##    ##  ##
 ######   ##### ##  #######  ##     ## ##     ## ######## ##     ## #########
*/

func squarer2(out chan<- float64, in <-chan float64) {
	dist := distuv.Normal{
		Mu:    9,
		Sigma: 1,
		//Rate:	0.01,
	}
	for v := range in {
		out <- v * dist.Rand()
	}
	close(out)
}

/*
 #######  ##     ##    ###    ########
##     ## ##     ##   ## ##   ##     ##
##     ## ##     ##  ##   ##  ##     ##
##     ## ##     ## ##     ## ##     ##
##  ## ## ##     ## ######### ##     ##
##    ##  ##     ## ##     ## ##     ##
 ##### ##  #######  ##     ## ########
*/

func quad(out chan<- float64, in1 <-chan float64, in2 <-chan float64) {
	/*dist := distuv.Exponential{
		Rate:    -2,
	}*/
	for v1 := range in1 {
		for v2 := range in2 {
			out <- v1 * v2 //dist.Rand()
		}
	}
	close(out)
}

/*
########  ########  #### ##    ## ######## ######## ########
##     ## ##     ##  ##  ###   ##    ##    ##       ##     ##
##     ## ##     ##  ##  ####  ##    ##    ##       ##     ##
########  ########   ##  ## ## ##    ##    ######   ########
##        ##   ##    ##  ##  ####    ##    ##       ##   ##
##        ##    ##   ##  ##   ###    ##    ##       ##    ##
##        ##     ## #### ##    ##    ##    ######## ##     ##
*/

func printer(in <-chan float64) {
	for v := range in {
		fmt.Println(v)
	}
}

/*
##     ##    ###    #### ##    ##
###   ###   ## ##    ##  ###   ##
#### ####  ##   ##   ##  ####  ##
## ### ## ##     ##  ##  ## ## ##
##     ## #########  ##  ##  ####
##     ## ##     ##  ##  ##   ###
##     ## ##     ## #### ##    ##
*/

func main() {
	//trace.Start(os.Stderr)
	quit := make(chan bool)
	//in1 := make(chan float64)
	in := make(chan float64)
	out1 := make(chan float64)
	//out12 := make(chan float64)
	out2 := make(chan float64)
	//out3 := make(chan float64)
	out4 := make(chan float64)
	//squares2 := make(chan float64)
	//quads := make(chan float64)
	//var chansIn ch //[1]ch
	/*for i := range chans {
	   chans[i] = make(ch)
	}*/
	//var chansOut <-chan float64 //[1]ch
	/*for i := range chans {
		 chans[i] = make(ch)
	}*/
	//var n1 F
	n1 := func(out chan<- float64, in <-chan float64, q chan bool) {
		dist := distuv.Normal{
			Mu:    1,
			Sigma: 1,
			//Rate: 0.01,
		}
		for x := 0; x <= 200000; x++ {
			out <- dist.Rand()
		}
		close(out)
		q <- true
	}

	n2 := func(out chan<- float64, in <-chan float64) {
		/*dist := distuv.Normal{
			Mu:    0.1,
			Sigma: 0.1,
		}*/
		for v := range in {
			out <- 30*v/v//3*v + dist.Rand()
		}
		close(out)
	}

	/*n3 := func(out chan<- float64, in <-chan float64) {
		dist := distuv.Normal{
			Mu:    0.1,
			Sigma: 0.1,
		}
		for v := range in {
			out <- v + dist.Rand()
		}
		close(out)
	}*/

	n4 := func(out chan<- float64, in1 <-chan float64, in2 <-chan float64, q chan bool) {
		dist := distuv.Normal{
			Mu:    0.1,
			Sigma: 0.1,
		}
		//for v1, v2 := range in1, in2, {
		//		out <- v1/v2 + dist.Rand(),
			//close(out)
		//}

		for {
		      v := 0.0
		      select {
		      case <-in1:
						v = <-in2*math.Cos(<-in1) + dist.Rand()

		      case <-in2:
		        v = <-in2*math.Cos(<-in1) + dist.Rand()
					//default:
					case <- q:
						close(out)
						return

		      }
		      out <- v// + dist.Rand()
		    }
			}

	go n1(out1, in, quit)
	//go n1(out1, in, quit)
	go n2(out2, out1)
	//go n3(out3, out1)
	go n4(out4, out2, out1, quit)
	//go counter(in)
	//go counter(naturals2)
	//go squarer1(squares1, naturals1)
	//go squarer2(squares2, naturals1)
	//go quad(quads, squares1, squares2)
	//v := node.function
	//fmt.Println(v)
	printer(out2)
	//trace.Stop()
}
