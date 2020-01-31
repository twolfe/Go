package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"math"
	//"sort"
	//"flag"
	//"text/tabwriter"
	//"phd/polymorphism"
	//"github.com/biogo/boom"
)

type Polymorphism struct {
  Chr string
  Pos int
  Ancestral string
  A int
  C int
  G int
  T int
  //H float64
  NumVar int
	DAN int //derived allele number
}

type Gene struct {
  Name  string
  S []Polymorphism //Segregating sites
  //N int //Number of sequences in sample
  Xi []int //Derived Allele Frequencies
	ThetaPi float64
  ThetaW float64
  ThetaH float64 //Dervied allele needed
  TajimaD float64
	TajVar float64
  FuLiEtaD float64
  FuLiEtaF float64
  FuLiXiD float64
  FayWuH float64
}

/*func CalculateHet(p []Polymorphism) []Polymorphism {
  for _, p := range p {
    fmt.Printf("gene[%s] pos[%d] altHetero[%d] altHom[%d] refHom[%d]\n", d.Chr, d.Pos, d.AltHetCount1, d.AltHomCount1, d.RefHomCount)
  }
}*/

//type DataGene []Gene

func PopulatePolymorphisms(f *os.File, p []Polymorphism) []Polymorphism {
  input := bufio.NewScanner(f)
  for input.Scan() {
    line := input.Text()
    polymorphism := Polymorphism{"", 0, "", 0, 0, 0, 0, 0, 0} //initialize polymorphism
    inputList := strings.Split(line, " ")
    polyList := strings.Split(inputList[2], ",")
    alleleA, _ := strconv.Atoi(polyList[0])
    alleleC, _ := strconv.Atoi(polyList[1])
    alleleG, _ := strconv.Atoi(polyList[2])
    alleleT, _ := strconv.Atoi(polyList[3])
    alleleCount := float64(alleleA) + float64(alleleC) + float64(alleleG) + float64(alleleT)
		//fmt.Printf("%s\n", inputList[9])
		Alikelihood, _ := strconv.ParseFloat(inputList[9], 64)
		Clikelihood, _ := strconv.ParseFloat(inputList[10], 64)
		Glikelihood, _ := strconv.ParseFloat(inputList[11], 64)
		Tlikelihood, _ := strconv.ParseFloat(inputList[12], 64)
		ancestral := []float64{Alikelihood, Clikelihood, Glikelihood, Tlikelihood}
		derived := []float64{float64(alleleA), float64(alleleC), float64(alleleG), float64(alleleT)}
		ancestralMaxAllele:= MaxLikelihoodAllele(ancestral)
		derivedMaxAllele := MaxLikelihoodAllele(derived)

    hypo := math.Sqrt((math.Pow(float64(alleleA), 2) + math.Pow(float64(alleleC), 2) + math.Pow(float64(alleleG), 2) + math.Pow(float64(alleleT), 2)))

    // Enter and make polymorphism if heterozygote or if polyploid major is different from
    if alleleCount != hypo || ancestralMaxAllele != derivedMaxAllele {
      polymorphism.Chr = inputList[0]
      polymorphism.Pos, _ = strconv.Atoi(inputList[1])
      polymorphism.A = alleleA
      polymorphism.C = alleleC
      polymorphism.G = alleleG
      polymorphism.T = alleleT
      polymorphism.NumVar = alleleA + alleleC + alleleG + alleleT
			//Which alleles are derived?
			switch {
				case ancestralMaxAllele == "A":
					polymorphism.DAN = alleleC + alleleG + alleleT
				case ancestralMaxAllele == "C":
					polymorphism.DAN = alleleA + alleleG + alleleT
				case ancestralMaxAllele == "G":
					polymorphism.DAN = alleleA + alleleC + alleleT
				case ancestralMaxAllele == "T":
					polymorphism.DAN = alleleA + alleleG + alleleC
			}
			//fmt.Printf("%s %d %s %s %s %s %d\n", polymorphism.Chr, polymorphism.Pos, ancestral, ancestralMaxAllele, derived, derivedMaxAllele, polymorphism.DAN)
      //polymorphism.H = polymorphism.Heterozigosity()
      p = append(p, polymorphism)
    }
  }
  return p
}

func PopulateGenes(p []Polymorphism, g []Gene) []Gene {
	pInG := make([]Polymorphism, 0)
	Xi := make([]int, 0)
	geneName := ""
	gene := Gene{"", pInG, Xi, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	for i := 0; i < len(p); i++ {
		switch {
		case i == 0:
			geneName = p[i].Chr
			pInG = append(pInG, p[i])
		case p[i].Chr == p[i-1].Chr && i < len(p)-1:
			pInG = append(pInG, p[i])
		case p[i].Chr != p[i-1].Chr && i < len(p)-1:
			gene.Name = geneName
			gene.S = pInG
			g = append(g, gene)
			geneName = p[i].Chr
			pInG = make([]Polymorphism, 0)
			pInG = append(pInG, p[i])
		case p[i].Chr == p[i-1].Chr && i == len(p)-1:
			pInG = append(pInG, p[i])
			gene.Name = geneName
			gene.S = pInG
			g = append(g, gene)
    case p[i].Chr != p[i-1].Chr && i == len(p)-1:
      gene.Name = geneName
			gene.S = pInG
			g = append(g, gene)
			geneName = p[i].Chr
			pInG = make([]Polymorphism, 0)
			pInG = append(pInG, p[i])
		}
	}
	return g
}

/*func PopulateXi(g []Gene) []int{
	input := bufio.NewScanner(f)
	for input.Scan() {
		line := input.Text()
		firstChar := strings.Split(line, "\t")[0]
		if firstChar != "CHROM" {
			entry := strings.Split(line, "\t")
      chr := entry[0]
			pos, _ := strconv.Atoi(entry[1])
			sinDouble := entry[2]
			for i := 0; i < len(p); i++ {
				if p[i].Chr == chr && p[i].Pos == pos && sinDouble == "S" {
					p[i].Singleton = true
					//fmt.Printf("gene[%s] pos[%d] singleton[%s]\n", chr, pos, sinDouble)
				}
			}
		}
	}
	return p
}*/
//WRITE METHODS FOR VARIATION MEASURES
func MaxLikelihoodAllele(array []float64) (string) {
	var maxIndex int = 0
	var maxAllele string = ""
	for i := 0; i <= 3; i++ {
		if array[maxIndex] < array[i] {
			maxIndex = i
		}
	}
	switch {
	case maxIndex == 0:
		maxAllele = "A"
	case maxIndex == 1:
		maxAllele = "C"
	case maxIndex == 2:
		maxAllele = "G"
	case maxIndex == 3:
		maxAllele = "T"

	}
	return maxAllele
}

func MinMax(array []int) (int, int) {
	var max int = array[0]
	var min int = array[0]
	for _, i := range array {
		if max < i && i != 0 {
			max = i
		}
		if min > i && i != 0 {
			min = i
		}
	}
	return min, max
}

func (p Polymorphism) Heterozigosity() float64 {
		n := float64(p.NumVar)/float64(p.NumVar-1)
		slice := []int{p.A, p.C, p.G, p.T}
		min, max := MinMax(slice)
		tot := min + max
		sumP := 0.0
		if tot == p.NumVar {
			hetP := math.Pow((float64(max)/float64(p.NumVar)), 2)
			hetQ := math.Pow((float64(min)/float64(p.NumVar)), 2)
			sumP = hetP + hetQ
		} else {
			diff := p.NumVar - tot
			hetP := math.Pow((float64(tot)/float64(p.NumVar)), 2)
			hetQ := math.Pow((float64(diff)/float64(p.NumVar)), 2)
			sumP = hetP + hetQ
		}
		/*When more than 2 symbols exist at a site, the second most frequent is used to compute the frequency.
		All the other ones are pooled and considered as the most frequent one.*/
		alSum := 1.0 - sumP
		H := n*alSum
		//fmt.Printf("gene[%s] pos[%d] Heterozigosity[%f]\n", p.Chr, p.Pos, H)

		return H
}

func (g Gene) MakeXi() []int {
	n := g.S[0].NumVar
	Xi := make([]int, n)
	for _, p := range g.S {
		Xi[p.DAN-1] += 1
	}
	return Xi
}

func (g Gene) CalculateThetaPi() float64 {
	ThetaPi := 0.0
	for _, p := range g.S {
		ThetaPi += p.Heterozigosity()
	}
	//fmt.Printf("ThetaPi[%f]\n", ThetaPi)
	return ThetaPi
}

func (g Gene) CalculateThetaW() float64 {
	ThetaW := 0.0
	S := len(g.S)
	n := g.S[0].NumVar
	//fmt.Printf("S[%d]\n", S)
	a := 0.0
	for i := 1; i < n; i++ {
		a += 1.0/float64(i)
	}
	ThetaW = float64(S)/a
	//fmt.Printf("ThetaW[%f]\n", ThetaW)
	return ThetaW
}

func (g Gene) CalculateThetaH() float64 {
	ThetaH := 0.0
	n := g.S[0].NumVar
	for i := 1; i < n; i++ {
		ThetaH += math.Pow(float64(i), 2)*float64(g.Xi[i-1])
	}
	//fmt.Printf("ThetaW[%f]\n", ThetaW)
	return ThetaH/(float64(n)*(float64(n)-1.0)*1.0/2.0)
}

func (g Gene) CalculateTajimaD() (float64, float64) {
	tajimaD := 0.0
	tajVar := 1.0
	delta := g.ThetaPi - g.ThetaW
	S := float64(len(g.S))
	n := g.S[0].NumVar
	a1 := 0.0
	a2 := 0.0
	for i := 1; i < n; i++ {
		a1 += 1.0/float64(i)
		i2 := math.Pow(float64(i), 2)
		a2 += 1.0/float64(i2)
	}
	b1 := (float64(n) + 1.)/(3.*(float64(n) - 1.))
	b2 := (2.*(math.Pow(float64(n), 2) + float64(n) + 3))/(9.*float64(n)*(float64(n) - 1))
  c1 := b1 - 1./a1
	c2 := b2 - (float64(n) + 2.)/(a1*float64(n)) + a2/math.Pow(a1, 2)
	e1 := c1/a1
	e2 := c2/(math.Pow(a1, 2) + a2)
	tajVar =  e1*S + e2*S*(S-1.)
	tajimaD = delta/math.Sqrt(tajVar)
	//fmt.Printf("TajimaD[%f]\n", TajimaD)
	return tajimaD, tajVar
}

func (g Gene) CalculateFayWuH() (float64) {
	return g.ThetaPi - g.ThetaH
	}

func main() {

  filePoly, _ := os.Open(os.Args[1])

	//fileSingle, _ := os.Open(os.Args[2])

  dataPolymorphisms := make([]Polymorphism, 0)

	dataGenes := make([]Gene, 0)

  dataPolymorphisms = PopulatePolymorphisms(filePoly, dataPolymorphisms)

	//dataPolymorphisms = PopulateSingleton(fileSingle, dataPolymorphisms)

	dataGenes = PopulateGenes(dataPolymorphisms, dataGenes)

	for i := 0; i < len(dataGenes); i++ {
		dataGenes[i].Xi = dataGenes[i].MakeXi()
		dataGenes[i].ThetaPi = dataGenes[i].CalculateThetaPi()
		dataGenes[i].ThetaW = dataGenes[i].CalculateThetaW()
		dataGenes[i].ThetaH = dataGenes[i].CalculateThetaH()
		dataGenes[i].FayWuH = dataGenes[i].CalculateFayWuH()
		dataGenes[i].TajimaD, dataGenes[i].TajVar = dataGenes[i].CalculateTajimaD()
	}
	//Xi := make([]int, dataGenes[1].S[1].NumVar)
	fmt.Println("Gene Num_segregating ThetaPi ThetaW ThetaH TajimaD FayWuH")
  for _, g := range dataGenes{
		if len(g.S) >= 10 {
			fmt.Printf("%s %d %f %f %f %f %f\n", g.Name, len(g.S), g.ThetaPi, g.ThetaW, g.ThetaH, g.TajimaD, g.FayWuH)
			//fmt.Println("Xi = ", g.Xi)
			/*for i := 0; i < len(g.Xi); i++ {
				Xi[i] += g.Xi[i]
			}*/
		}
			//fmt.Println("Xi = ", Xi)
      //fmt.Printf("%s %d %f %f %f\n", g.Name, len(g.S), g.ThetaPi, g.ThetaW, g.TajimaD)
      //fmt.Printf("gene[%s] S[%d] Poly[%s]\n", g.Name, len(g.S), g.S)

  }
}
