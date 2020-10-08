package main

import (
	//"bufio"
	"fmt"
	//"os"
	//"strconv"
	//"strings"
)

type Organism interface {
	Children() bool
}

type Host struct {
	sex      bool
	infected Wolbachia
	resitant bool
}

func (h Host) Children() bool {
	return h.resitant
}

type Wolbachia struct {
	ci bool
}

func (w Wolbachia) Children() bool {
	return w.ci
}

func replicate(o Organism) {
	fmt.Println(o)
	fmt.Println(o.Children())
}

func migrate(h Host) {
	fmt.Println(h)
}

func main() {
	jane := Wolbachia{
		ci: true,
	}
	//jane := o
	harry := Host{
		sex:      false,
		infected: jane,
		resitant: false,
	}

	replicate(harry)
	replicate(jane)
	migrate(harry)
}

//interface mutate

//method migrate
