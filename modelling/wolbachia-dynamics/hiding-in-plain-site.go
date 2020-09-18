package main

import (
	//"bufio"
	"fmt"
	//"os"
	//"strconv"
	//"strings"
)

type Host struct {
	sex      string
	infected Wolbachia
	resitant bool
}

type Wolbachia struct {
	ci bool
}

func main() {
	jane := Wolbachia{
		ci: true,
	}
	harry := Host{
		sex:      "male",
		infected: jane,
		resitant: false,
	}

	fmt.Println(harry)
}

//interface mutate

//method migrate
