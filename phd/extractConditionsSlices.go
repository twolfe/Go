//Spinner displays an animation while computing the 45th Fibonacci number.
//Make a package from this and formal tests
package main

import (
	//"bufio"
	"fmt"
	"os"
	//"strings"
	//"strconv"
	"phd/polymorphism"
	//"github.com/biogo/boom"
)

func main() {
	// Use _, ok := m["gene"] to check if gene is already in the map
	//Data: a map of polymorphisms with entries true for dignostic SNP and false for no-diagnotic SNP
	//Change from map to array, slices. No searching, just running through + order to check previous polymorphism
	//Comand-line arguments
	file, _ := os.Open(os.Args[1])
	//parent is the parental genome of particular interest
	parent := os.Args[2]
	var data polymorphism.Data
	data = polymorphism.Populate(file, data)
	file.Close()

	data.ParentAssignAllele()
	dataD := polymorphism.MakeDiagnosticData(data)
	//fmt.Printf("%s\n", data[:])
	//bam stuff

	/*bf, err := boom.OpenBAMFile("/home/thomas/Documents/phd/research/snp/testing.bam")
	      if err != nil {
	  		panic(err)
	  	}
	      r, _, err := bf.Read()
	  		if err != nil {
	              fmt.Printf("%s\n", "error")
	  		}
	      fmt.Printf("%s\n", r.Seq())
	*/
	/*for i := range dataD {
		if dataD[i].Diagnostic == true && i != (len(dataD)-1) {
			dataD[i].DisplayPolymorphism(parent)
			fmt.Printf("%s", "|\n")
		} else {
			dataD[i].DisplayPolymorphism(parent)
		}
	}*/
	//for i := range dataD {
		//if dataD[i].Diagnostic == true {
			//dataD.DisplayDiagnosisData(parent)
        //}
    //}
    slice := polymorphism.MakeSliceUniqueChromo(dataD)
    //fmt.Printf("%s", slice)
    for s := range(slice){
        d := polymorphism.MakeDataUniqueChromo(slice[s], dataD)
        d.DisplayDiagnosisData(parent)
    }
    
    fmt.Printf("%s", "foo = ")
    for s := range(slice){
        if s < (len(slice)-1){
        fmt.Printf("%s %s", slice[s], "| ")
        } else {
            fmt.Printf("%s %s", slice[s], ";")
        }
    }
    
}
