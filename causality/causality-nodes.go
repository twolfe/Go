// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 231.

// Pipeline3 demonstrates a finite 3-stage pipeline
// with range, close, and unidirectional channel types.
package main

import (
	"fmt"

	"gonum.org/v1/gonum/stat/distuv"
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

type F func(out chan<- float64, in <-chan float64)
type Node struct {
	input    chan float64
	output   chan float64
	parents  *Node // note that this is now a pointer
	children *Node // note that this is now a pointer
	function F
}

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
	in := make(chan float64)
	//naturals2 := make(chan float64)
	out := make(chan float64)
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
	node := Node{
		input:    in,
		output:   out,
		parents:  nil,
		children: nil,
		function: func(out chan<- float64, in <-chan float64) {
			for v := range in {
				out <- v / v
			}
			close(out)
		},
	}

	go counter(in)
	//go counter(naturals2)
	//go squarer1(squares1, naturals1)
	//go squarer2(squares2, naturals1)
	//go quad(quads, squares1, squares2)
	//v := node.function
	go node.function(node.output, node.input)
	//fmt.Println(v)
	printer(node.output)
}
