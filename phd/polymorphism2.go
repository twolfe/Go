package polymorphism

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Polymorphism struct {
	Gene       string
	Pos        int
	Ref        string
	Alt        string
	Ch         string
	P1         string
	P2         string
	Diagnostic bool //for string organised code
}

type Data []Polymorphism

//Function: make a map containing polymorphisms as keys and dignosis as entries from HyLiTE input
func Populate(f *os.File, d Data) Data {
	input := bufio.NewScanner(f)
	for input.Scan() {
		line := input.Text()
		entry := strings.Split(line, "\t")
		p, err := strconv.Atoi(entry[1])
		// Check that postion does exist and is really a postion (header is not a position)
		if err == nil {
			polymorphism := Polymorphism{entry[0], p, entry[2], entry[3], entry[4], entry[5], entry[6], false}
			//polymorphism = Polymorphism{entry[0], p, entry[2], entry[3], entry[4], entry[5], entry[6], polymorphism.Diagnosis()}
			//polymorphism = polymorphism.ParentalAssign()
			d = append(d, polymorphism)
		}
	}
	return d
}

//Function: returns whether a polymorphism is diagnotic or not
//make a Method o polymorphism
func (p Polymorphism) Diagnosis() bool {
	diagnostic := false
	//Condition: SNP is fixed and ancestral in either parents
	if (p.P1 == "0,0" && p.P2 == "1,1") || (p.P1 == "1,1" && p.P2 == "0,0") {
		diagnostic = true
		//Condition: fixed but non-ancestral SNP
	}
	return diagnostic
}

//Problemaric conditions: when non-fixed parental SNP but different alleles between parents

//Method: assigns the right allele to each parent in polymorphism if it is diagnostic
func (p Polymorphism) ParentAssignAllele() Polymorphism {
	if p.Diagnosis() {
		p.Diagnostic = true
		if p.P1 == "1,1" {
			p.P1 = p.Alt
			p.P2 = p.Ref
		} else if p.P1 == "0,0" {
			p.P1 = p.Ref
			p.P2 = p.Alt
		}
	}
	return p
}

func (d Data) ParentAssignAllele() Data {
	for i := range d {
		d[i] = d[i].ParentAssignAllele()
	}
	return d
}

//Create a new Data string only containing diagnosis polymophisms
func MakeDiagnosticData(d Data) Data {
	var r Data
	for i := range d {
		if d[i].Diagnostic {
			r = append(r, d[i])
		}
	}
	return r
}

//Make slice of unique chromosomes
func MakeSliceUniqueChromo(d Data) []string {
	var slice []string
	for i := range d {
		if i == 0 {
			slice = append(slice, d[i].Gene)
		} else if i > 0 && d[i].Gene != d[i-1].Gene {
			slice = append(slice, d[i].Gene)
		}
	}
	return slice
}

func MakeDataUniqueChromo(s string, d Data) Data {
	var dout Data
	for i := range d {
		if d[i].Gene == s {
			dout = append(dout, d[i])
		}
	}
	return dout
}

// DisplayDiagnosis takes a Data string which only contain diagnosis polymorphisms & where all Polymorphisms are on the same chromosome
func (d Data) DisplayDiagnosisData(parent string) {
	if parent == "P1" {
		for i := 0; i <= (len(d) - 1); i++ {
			if len(d) == 1 {
				fmt.Printf("%s %s %s %s %d %s %s %s", d[i].Gene, "= (chr(", d[i].Gene, ") & nt_exact(", d[i].Pos, ",", d[i].P1, ") );\n")
			} else if len(d) != 1 && i == 0 {
				fmt.Printf("%s %s %s %s %d %s %s %s", d[i].Gene, "= (chr(", d[i].Gene, ") & ( nt_exact(", d[i].Pos, ",", d[i].P1, ") | ")
			} else if len(d) != 1 && i > 0 && i != (len(d)-1) {
				fmt.Printf("%s %d %s %s %s", "nt_exact(", d[i].Pos, ",", d[i].P1, ") | ")
			} else if len(d) != 1 && i == (len(d)-1) {
				fmt.Printf("%s %d %s %s %s", "nt_exact(", d[i].Pos, ",", d[i].P1, ") ) );\n")
			}
		}
	}

	if parent == "P2" {
		for i := 0; i <= (len(d) - 1); i++ {
			if len(d) == 1 {
				fmt.Printf("%s %s %s %s %d %s %s %s", d[i].Gene, "= (chr(", d[i].Gene, ") & nt_exact(", d[i].Pos, ",", d[i].P2, ") );\n")
			} else if len(d) != 1 && i == 0 {
				fmt.Printf("%s %s %s %s %d %s %s %s", d[i].Gene, "= (chr(", d[i].Gene, ") & ( nt_exact(", d[i].Pos, ",", d[i].P2, ") | ")
			} else if len(d) != 1 && i > 0 && i != (len(d)-1) {
				fmt.Printf("%s %d %s %s %s", "nt_exact(", d[i].Pos, ",", d[i].P2, ") | ")
			} else if len(d) != 1 && i == (len(d)-1) {
				fmt.Printf("%s %d %s %s %s", "nt_exact(", d[i].Pos, ",", d[i].P2, ") ) );\n")
			}
		}
	}
}

// Displayfull takes a Data string
func (p Polymorphism) DisplayPolymorphism(parent string) {
	if parent == "P1" {
		fmt.Printf("%s %s %s %d %s %s %s", "(chr(", p.Gene, ") & nt_exact(", p.Pos, ",", p.P1, "))")
	} else if parent == "P2" {
		fmt.Printf("%s %s %s %d %s %s %s", "(chr(", p.Gene, ") & nt_exact(", p.Pos, ",", p.P2, "))")
	}
}

//Use biogo for extracting reads
