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
  Ref string
  Alt1 string
  Alt2 string
  Alt3 string
  AltHetP1 float64
  AltHetP2 float64
  AltHetP3 float64
  AltHomP1 float64
  AltHomP2 float64
  AltHomP3 float64
  RefHomP float64
	NumVar int
	H float64
	Singleton bool
}

type Gene struct {
  Name  string
  S []Polymorphism //Segregating sites
  //N int //Number of sequences in sample
  Eta1 int //Number of singletons
  //Eta2 int //Number of doubletons
  Xi int //Number of derived singletons
  ThetaPi float64
  ThetaW float64
  ThetaH float64 //Dervied allele needed
  TajimaD float64
  FuLiEtaD float64
  FuLiEtaF float64
  FuLiXiD float64
  FayWuF float64
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
    firstChar := strings.Split(line, "\t")[0]

    if firstChar != "CHR" {

      //SNP part
      polymorphism := Polymorphism{"", 0, "", "", "", "", 0, 0, 0, 0, 0, 0, 0, 0, 0, false} //initialize polymorphism

      entry := strings.Split(line, "\t")
      polymorphism.Chr = entry[0]
      polymorphism.Pos, _ = strconv.Atoi(entry[1])
      polymorphism.Ref = entry[3]

      alternative :=  strings.Split(entry[4], ",")

      altHeterozygoteCountList := strings.Split(entry[5], ",")

      altHomozygoteCountList := strings.Split(entry[6], ",")

      numAlt := len(alternative)

      numSeq := 0

      for i := 0; i < numAlt; i++ {
        /*hetero, _ := strconv.ParseFloat(altHeterozygoteCountList[i], 64)
        homo, _ := strconv.ParseFloat(altHomozygoteCountList[i], 64)*/
				hetero, _ := strconv.Atoi(altHeterozygoteCountList[i])
				homo, _ := strconv.Atoi(altHomozygoteCountList[i])
        iAdd :=  0
        iAdd = hetero + homo
        numSeq += iAdd
      }

      refHomo, _ := strconv.Atoi(entry[7])
      numSeq += refHomo

      switch {
      case numAlt == 1:
        polymorphism.Alt1 = alternative[0]
        //HetP1, _ := strconv.Atoi(altHeterozygoteCountList[0], 64)
				HetP1, _ := strconv.Atoi(altHeterozygoteCountList[0])
        polymorphism.AltHetP1= float64(HetP1)/float64(numSeq)
        HomP1, _ := strconv.Atoi(altHomozygoteCountList[0])
        polymorphism.AltHomP1= float64(HomP1)/float64(numSeq)
      case numAlt == 2:
        polymorphism.Alt1 = alternative[0]
        HetP1, _ := strconv.Atoi(altHeterozygoteCountList[0])
        polymorphism.AltHetP1 = float64(HetP1)/float64(numSeq)
        HomP1, _ := strconv.Atoi(altHomozygoteCountList[0])
        polymorphism.AltHomP1= float64(HomP1)/float64(numSeq)
        polymorphism.Alt2 = alternative[1]
        HetP2, _ := strconv.Atoi(altHeterozygoteCountList[1])
        polymorphism.AltHetP2 = float64(HetP2)/float64(numSeq)
        HomP2, _ := strconv.Atoi(altHomozygoteCountList[1])
        polymorphism.AltHomP2 = float64(HomP2)/float64(numSeq)
      case numAlt == 3:
        polymorphism.Alt1 = alternative[0]
        HetP1, _ := strconv.Atoi(altHeterozygoteCountList[0])
        polymorphism.AltHetP1 = float64(HetP1)/float64(numSeq)
        HomP1, _ := strconv.Atoi(altHomozygoteCountList[0])
        polymorphism.AltHomP1 = float64(HomP1)/float64(numSeq)
        polymorphism.Alt2 = alternative[1]
        HetP2, _ := strconv.Atoi(altHeterozygoteCountList[1])
        polymorphism.AltHetP2 = float64(HetP2)/float64(numSeq)
        HomP2, _ := strconv.Atoi(altHomozygoteCountList[1])
        polymorphism.AltHomP2 = float64(HomP2)/float64(numSeq)
        polymorphism.Alt3 = alternative[2]
        HetP3, _ := strconv.Atoi(altHeterozygoteCountList[2])
        polymorphism.AltHetP3 = float64(HetP3)/float64(numSeq)
        HomP3, _ := strconv.Atoi(altHomozygoteCountList[2])
        polymorphism.AltHomP3 = float64(HomP3)/float64(numSeq)
      }

      polymorphism.RefHomP = float64(refHomo)/float64(numSeq)

			polymorphism.NumVar = numSeq

			polymorphism.H = polymorphism.Heterozigosity()

      p = append(p, polymorphism)
    }
  }
  return p
}

func PopulateSingleton(f *os.File, p []Polymorphism) []Polymorphism {
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
}

func PopulateGenes(p []Polymorphism, g []Gene) []Gene {
	pInG := make([]Polymorphism, 0)
	geneName := ""
	gene := Gene{"", pInG, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

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
		case i == len(p)-1:
			pInG = append(pInG, p[i])
			gene.Name = geneName
			gene.S = pInG
			g = append(g, gene)
		}
		/*if i == 0 {
			geneName = p[i].Chr
			pInG = append(pInG, p[i])
		} else if p[i].Chr == p[i-1].Chr {
			pInG = append(pInG, p[i])
		} else if  p[i].Chr != p[i-1].Chr {
			gene.Name = geneName
			gene.S = pInG
			g = append(g, gene)
			geneName = p[i].Chr
			pInG = make([]Polymorphism, 0)
			pInG = append(pInG, p[i])
		} else if i == len(p)-1 {
			pInG = append(pInG, p[i])
			gene.Name = geneName
			gene.S = pInG
			g = append(g, gene)
		}*/
	}
	return g
}

//WRITE METHODS FOR VARIATION MEASURES

func (p Polymorphism) Heterozigosity() float64 {
		n := float64(p.NumVar)/float64(p.NumVar-1)
		het1 := math.Pow(p.AltHetP1, 2)
		het2 := math.Pow(p.AltHetP2, 2)
		het3 := math.Pow(p.AltHetP3, 2)
		hom1 := math.Pow(p.AltHomP1, 2)
		hom2 := math.Pow(p.AltHomP2, 2)
		hom3 := math.Pow(p.AltHomP3, 2)
		refHom := math.Pow(p.RefHomP, 2)
		sumP := het1+het2+het3+hom1+hom2+hom3+refHom
		alSum := 1.0 - sumP
		H := n*alSum
		//fmt.Printf("gene[%s] pos[%d] Heterozigosity[%f]\n", p.Chr, p.Pos, p.H)

		return H
}

func (g Gene) CalculateThetaPi() float64 {
	ThetaPi := 0.0
	for _, p := range g.S {
		ThetaPi += p.H
	}
	fmt.Printf("ThetaPi[%f]\n", ThetaPi)
	return ThetaPi
}

func (g Gene) CalculateThetaW() float64 {
	ThetaW := 0.0
	S := len(g.S)
	n := g.S[0].NumVar
	fmt.Printf("S[%d]\n", S)
	a := 0.0
	for i := 1; i < n; i++ {
		a += 1.0/float64(i)
	}
	ThetaW = float64(S)/a
	fmt.Printf("ThetaW[%f]\n", ThetaW)
	return ThetaW
}

func (g Gene) CalculateTajimaD() float64 {
	TajimaD := 0.0
	TajVar := 1.0
	TPi := g.ThetaPi
	TW := g.ThetaW
	S := float64(len(g.S))
	n := float64(g.S[0].NumVar)
	n2 := math.Pow(n, 2)
	a := 0.0
	a2 := 0.0
	for i := 1.0; i < n; i++ {
		a += 1.0/float64(i)
		i2 := math.Pow(i, 2)
		a2 += 1.0/float64(i2)
	}
	a22 := math.Pow(a, 2)
	TajVar = ((n+1)/(3*(n-1))-1/a)*S + ((2*(n2+n+3)/9*n*(n-1) - (n+2)/n*a + a2/a22)/(a2+a22))*S*(S-1)
	TajimaD = (TPi-TW)/TajVar
	fmt.Printf("TajimaD[%f]\n", TajimaD)
	return TajimaD
}

func main() {

  filePoly, _ := os.Open(os.Args[1])

	fileSingle, _ := os.Open(os.Args[2])

  dataPolymorphisms := make([]Polymorphism, 0)

	dataGenes := make([]Gene, 0)

  dataPolymorphisms = PopulatePolymorphisms(filePoly, dataPolymorphisms)

	dataPolymorphisms = PopulateSingleton(fileSingle, dataPolymorphisms)

	dataGenes = PopulateGenes(dataPolymorphisms, dataGenes)

	for i := 0; i < len(dataGenes); i++ {
		dataGenes[i].ThetaPi = dataGenes[i].CalculateThetaPi()
		dataGenes[i].ThetaW = dataGenes[i].CalculateThetaW()
		dataGenes[i].TajimaD= dataGenes[i].CalculateTajimaD()
	}

  for _, g := range dataGenes {
			fmt.Printf("gene[%s] ThetaPi[%f] ThetaW[%f] TajimaD[%f]\n", g.Name, g.ThetaPi, g.ThetaW, g.TajimaD)

  }
}
