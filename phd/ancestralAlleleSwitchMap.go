package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	//"phd/polymorphism"
	//"github.com/biogo/boom"
)

type Key struct {
	Gene string
	Pos  int
}

type Allele struct {
	Freq float64
	Snp  string
}

type Polymorphism struct {
	Nallele int
	Nchr    int
	AA      Allele
	DA1     Allele
}

type Data map[Key]Polymorphism

// input has to be a comma seperated .frq file (from vcftools and no nan filtering)
func Populate(f *os.File, d Data) Data {
	input := bufio.NewScanner(f)
	for input.Scan() {
		line := input.Text()
		entry := strings.Split(line, ",")

		gene := entry[0]
		pos, err := strconv.Atoi(entry[1])

		key := Key{gene, pos}
		nAll, _ := strconv.Atoi(entry[2])
		nChr, _ := strconv.Atoi(entry[3])

		aa := strings.Split(entry[4], ":")
		if len(aa) > 1 {
			//fmt.Printf("%s\n", aa)
			aaFreq, _ := strconv.ParseFloat(aa[1], 64)
			da1 := strings.Split(entry[5], ":")
			da1Freq, _ := strconv.ParseFloat(da1[1], 64)
			aaS := Allele{aaFreq, aa[0]}
			da1S := Allele{da1Freq, da1[0]}

			// Check that postion does exist and is really a postion (header is not a position)
			if err == nil {
				polymorphism := Polymorphism{nAll, nChr, aaS, da1S}
				d[key] = polymorphism
			}
		}
	}
	return d
}

func main() {

	fileAA, _ := os.Open(os.Args[1]) // ancestral allele
	fileChange, _ := os.Open(os.Args[2])

	dataAA := make(Data)
	dataChange := make(Data)

	dataAA = Populate(fileAA, dataAA)
	dataChange = Populate(fileChange, dataChange)

	fileAA.Close()
	fileChange.Close()
	//fmt.Printf("%s\n", data[:])

	//Make switches in ancestral alleles here
	for k, p := range dataChange {
		// where the magic happens
		if dataAA[k].DA1.Freq == 1 {

			aaFreqChange := dataChange[k].AA.Freq
			da1FreqChange := dataChange[k].DA1.Freq
			aaSnpChange := dataChange[k].AA.Snp
			da1SnpChange := dataChange[k].DA1.Snp
			AAnew := Allele{da1FreqChange, da1SnpChange}
			DA1new := Allele{aaFreqChange, aaSnpChange}
			p = Polymorphism{dataChange[k].Nallele, dataChange[k].Nchr, AAnew, DA1new}
			dataChange[k] = p

		}
	}
	fmt.Printf("%s,%s,%s,%s,%s,%s,%s,%s\n", "CHROM", "POS", "N_ALLELES", "N_CHR", "AA", "DA1", "DA2", "DA3")
	for k, p := range dataChange {
		fmt.Printf("%s,%d,%d,%d,%f,%f,,\n", k.Gene, k.Pos, p.Nallele, p.Nchr, p.AA.Freq, p.DA1.Freq)
	}

}
