// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 231.

// Pipeline3 demonstrates a finite 3-stage pipeline
// with range, close, and unidirectional channel types.
package main

import (
	"gonum.org/v1/gonum/stat/distuv"
	"fmt"
	)

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
	dist := distuv.Exponential{
    //Mu:    9,
    //Sigma: 1,
		Rate:	0.01,
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
		out <- v//dist.Rand()
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
			out <- v1 * v2//dist.Rand()
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
	naturals1 := make(chan float64)
	naturals2 := make(chan float64)
	squares1 := make(chan float64)
	squares2 := make(chan float64)
	quads := make(chan float64)

	go counter(naturals1)
	go counter(naturals2)
	go squarer1(squares1, naturals1)
	go squarer2(squares2, naturals2)
	go quad(quads, squares1, squares2)
	printer(naturals1)
}
