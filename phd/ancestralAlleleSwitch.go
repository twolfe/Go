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

type Allele struct {
	Freq float64
	Snp  string
}

type Polymorphism struct {
	Gene    string
	Pos     int
	Nallele int
	Nchr    int
	AA      Allele
	DA1     Allele
}

type Data []Polymorphism

// input has to be a comma seperated .frq file (from vcftools and no nan filtering)
func Populate(f *os.File, d Data) Data {
	input := bufio.NewScanner(f)
	for input.Scan() {
		line := input.Text()
		entry := strings.Split(line, ",")

		pos, err := strconv.Atoi(entry[1])
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
				polymorphism := Polymorphism{entry[0], pos, nAll, nChr, aaS, da1S}
				d = append(d, polymorphism)
			}
		}
	}
	return d
}

func main() {

	fileAA, _ := os.Open(os.Args[1]) // ancestral allele
	fileChange, _ := os.Open(os.Args[2])

	var dataAA Data
	var dataChange Data

	dataAA = Populate(fileAA, dataAA)
	dataChange = Populate(fileChange, dataChange)

	fileAA.Close()
	fileChange.Close()
	//fmt.Printf("%s\n", data[:])

	//Make switches in ancestral alleles here
	for pol := range dataAA {
		// where the magic happens
		if dataAA[pol].DA1.Freq == 1 {
			aaSnpAA := dataAA[pol].AA.Snp
			aaSnpDA1 := dataAA[pol].DA1.Snp
			dataAA[pol].AA = Allele{1, aaSnpDA1}
			dataAA[pol].DA1 = Allele{0, aaSnpAA}
			aaFreqChange := dataChange[pol].AA.Freq
			da1FreqChange := dataChange[pol].DA1.Freq
			aaSnpChange := dataChange[pol].AA.Snp
			da1SnpChange := dataChange[pol].DA1.Snp
			dataChange[pol].AA = Allele{da1FreqChange, da1SnpChange}
			dataChange[pol].DA1 = Allele{aaFreqChange, aaSnpChange}
		}
	}
	fmt.Printf("%s,%s,%s,%s,%s,%s,%s,%s\n", "CHROM", "POS", "N_ALLELES", "N_CHR", "AA", "DA1", "DA2", "DA3")
	for pol := range dataChange {
		fmt.Printf("%s,%d,%d,%d,%s:%f,%s:%f,,\n", dataChange[pol].Gene, dataChange[pol].Pos, dataChange[pol].Nallele, dataChange[pol].Nchr, dataChange[pol].AA.Snp, dataChange[pol].AA.Freq, dataChange[pol].DA1.Snp, dataChange[pol].DA1.Freq)
	}

}
