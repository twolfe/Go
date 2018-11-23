package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sort"
	"flag"
	//"text/tabwriter"
	//"phd/polymorphism"
	//"github.com/biogo/boom"
)

//Usage: go run formatEST.go focal.species.vcf ancestral1.species.vcf ancestral2.species.vcf all.species.vcf "synonymous" > out.formatted.est
//focal.species.vcf: This file contains annotated focal snps with no missing data

//ancestral1.species.vcf: This file contains annotated snp data for the parent of interest. It contains all the positions

//ancestral2.species.vcf: This file contains annotated snp data for the other parent (outgroup). It contains all the positions
//Careful, these snps can often not be in the the file from the focal species because they are mapping to the other part of the sub-genome...

//all.species.vcf: This file is the same as ancestral.species.vcf. Can contain all species and must contain all postions.
//(used to assess the the reference snp, also outgroup)

type Key struct {
	Gene string
	Pos  int
}

type SNP struct {
	NumA int
	NumC int
	NumG int
	NumT int
	Effect string //The effect (synonymous, missense ...) of the snp
}

type Poly struct {
	Name string
	Number  int
}

type Data map[Key]SNP

// input has to be a comma seperated .frq file (from vcftools and no nan filtering)
func PopulateFocal(f *os.File, d Data) Data {
	input := bufio.NewScanner(f)
	for input.Scan() {

		line := input.Text()
		firstChar := strings.Split(line, "")[0]
		entry := strings.Split(line, "\t")
		var alt string
		if (firstChar != "#"){
			alt = entry[4]
		}

		//the firstChar part is to make sure we work on lines that are not comments.
		//the len part is to make sure there are less than 3 snp, no indels. Should do for now
		if (firstChar != "#") && (len(alt) <= 5) {

			//SNP part
			snp := SNP{0,0,0,0,""} //initialize snp

			ref := entry[3]

			//Possibility of two alternatives
			altSplit := strings.Split(alt, ",")
			//could have up to 3 alternative snps (eg. A,T,G with ref C)
			//Only 1 for the moment
			alt1 := ""
			alt2 := ""
			alt3 := ""
			numAlt := len(altSplit)
			switch numAlt {
			case 1:
				alt1 = altSplit[0]
			case 2:
				alt1 = altSplit[0]
				alt2 = altSplit[1]
			case 3:
				alt1 = altSplit[0]
				alt2 = altSplit[1]
				alt3 = altSplit[2]
			}

			//Work on the assigning the snps
			for i := 9; i <= (len(entry)-1); i++ {
				//Note: Have not found any 1/0 nor 1/2 nor 0/1
				polymorphism := strings.Split(entry[i], ":")[0]

				if polymorphism == "0/0" {
					switch ref {
					case "A":
							snp.NumA = snp.NumA + 2
					case "C":
							snp.NumC = snp.NumC + 2
					case "G":
							snp.NumG = snp.NumG + 2
					case "T":
							snp.NumT = snp.NumT + 2
					}
				} else if polymorphism == "1/1" {
					switch alt1 {
					case "A":
							snp.NumA = snp.NumA + 2
					case "C":
							snp.NumC = snp.NumC + 2
					case "G":
							snp.NumG = snp.NumG + 2
					case "T":
							snp.NumT = snp.NumT + 2
					}
				} else if polymorphism == "2/2" {
					switch alt2 {
					case "A":
							snp.NumA = snp.NumA + 2
					case "C":
							snp.NumC = snp.NumC + 2
					case "G":
							snp.NumG = snp.NumG + 2
					case "T":
							snp.NumT = snp.NumT + 2
					}
				} else if polymorphism == "3/3" {
					switch alt3 {
					case "A":
							snp.NumA = snp.NumA + 2
					case "C":
							snp.NumC = snp.NumC + 2
					case "G":
							snp.NumG = snp.NumG + 2
					case "T":
							snp.NumT = snp.NumT + 2
					}
				} else if polymorphism == "0/1"{
					switch ref {
					case "A":
							snp.NumA = snp.NumA + 1
					case "C":
							snp.NumC = snp.NumC + 1
					case "G":
							snp.NumG = snp.NumG + 1
					case "T":
							snp.NumT = snp.NumT + 1
					}
					switch alt1 {
					case "A":
							snp.NumA = snp.NumA + 1
					case "C":
							snp.NumC = snp.NumC + 1
					case "G":
							snp.NumG = snp.NumG + 1
					case "T":
							snp.NumT = snp.NumT + 1
					}
				} else if polymorphism == "0/2"{
					switch ref {
					case "A":
							snp.NumA = snp.NumA + 1
					case "C":
							snp.NumC = snp.NumC + 1
					case "G":
							snp.NumG = snp.NumG + 1
					case "T":
							snp.NumT = snp.NumT + 1
					}
					switch alt2 {
					case "A":
							snp.NumA = snp.NumA + 1
					case "C":
							snp.NumC = snp.NumC + 1
					case "G":
							snp.NumG = snp.NumG + 1
					case "T":
							snp.NumT = snp.NumT + 1
					}
				} else if polymorphism == "0/3"{
					switch ref {
					case "A":
							snp.NumA = snp.NumA + 1
					case "C":
							snp.NumC = snp.NumC + 1
					case "G":
							snp.NumG = snp.NumG + 1
					case "T":
							snp.NumT = snp.NumT + 1
					}
					switch alt3 {
					case "A":
							snp.NumA = snp.NumA + 1
					case "C":
							snp.NumC = snp.NumC + 1
					case "G":
							snp.NumG = snp.NumG + 1
					case "T":
							snp.NumT = snp.NumT + 1
					}
				}
			}

			//Find the effect (synonymous, missense ...) of the snp
			/*info := entry[7]
			//fmt.Printf("%s", info)
			annStr := strings.Split(info, ";")
			ann := annStr[len(annStr)-1]
			effect := strings.Split(ann, "|")[1] //OK effect is obtained
			snp.Effect = effect
			//fmt.Printf("%s", effect)
			*/
			//Key part
			gene := entry[0]
			pos, _ := strconv.Atoi(entry[1])

			key := Key{gene, pos}

			d[key] = snp
			//nAll, _ := strconv.Atoi(entry[2])
			//nChr, _ := strconv.Atoi(entry[3])
		}
	}
	return d
}

func PopulateAncestral(f *os.File, d Data) Data {
	input := bufio.NewScanner(f)
	for input.Scan() {

		line := input.Text()
		firstChar := strings.Split(line, "")[0]
		entry := strings.Split(line, "\t")
		var alt string
		if (firstChar != "#"){
			alt = entry[4]
		}

		//the firstChar part is to make sure we work on lines that are not comments.
		//the len part is to make sure there are less than 3 snp, no indels. Should do for now
		if (firstChar != "#") && (len(alt) <= 5) {

			//SNP part
			snp := SNP{0,0,0,0,""} //initialize snp

			ref := entry[3]

			//Possibility of two alternatives
			altSplit := strings.Split(alt, ",")
			//could have up to 3 alternative snps (eg. A,T,G with ref C)
			//Only 1 for the moment
			alt1 := ""
			alt2 := ""
			alt3 := ""
			numAlt := len(altSplit)
			switch numAlt {
			case 1:
				alt1 = altSplit[0]
			case 2:
				alt1 = altSplit[0]
				alt2 = altSplit[1]
			case 3:
				alt1 = altSplit[0]
				alt2 = altSplit[1]
				alt3 = altSplit[2]
			}

			//Work on the assigning the snps
			for i := 9; i <= (len(entry)-1); i++ {
				//Note: Have not found any 1/0 nor 1/2 nor 0/1
				polymorphism := strings.Split(entry[i], ":")[0]

				if polymorphism == "0/0" {
					switch ref {
					case "A":
							snp.NumA = snp.NumA + 2
					case "C":
							snp.NumC = snp.NumC + 2
					case "G":
							snp.NumG = snp.NumG + 2
					case "T":
							snp.NumT = snp.NumT + 2
					}
				} else if polymorphism == "1/1" {
					switch alt1 {
					case "A":
							snp.NumA = snp.NumA + 2
					case "C":
							snp.NumC = snp.NumC + 2
					case "G":
							snp.NumG = snp.NumG + 2
					case "T":
							snp.NumT = snp.NumT + 2
					}
				} else if polymorphism == "2/2" {
					switch alt2 {
					case "A":
							snp.NumA = snp.NumA + 2
					case "C":
							snp.NumC = snp.NumC + 2
					case "G":
							snp.NumG = snp.NumG + 2
					case "T":
							snp.NumT = snp.NumT + 2
					}
				} else if polymorphism == "3/3" {
					switch alt3 {
					case "A":
							snp.NumA = snp.NumA + 2
					case "C":
							snp.NumC = snp.NumC + 2
					case "G":
							snp.NumG = snp.NumG + 2
					case "T":
							snp.NumT = snp.NumT + 2
					}
				} else if polymorphism == "0/1"{
					switch ref {
					case "A":
							snp.NumA = snp.NumA + 1
					case "C":
							snp.NumC = snp.NumC + 1
					case "G":
							snp.NumG = snp.NumG + 1
					case "T":
							snp.NumT = snp.NumT + 1
					}
					switch alt1 {
					case "A":
							snp.NumA = snp.NumA + 1
					case "C":
							snp.NumC = snp.NumC + 1
					case "G":
							snp.NumG = snp.NumG + 1
					case "T":
							snp.NumT = snp.NumT + 1
					}
				} else if polymorphism == "0/2"{
					switch ref {
					case "A":
							snp.NumA = snp.NumA + 1
					case "C":
							snp.NumC = snp.NumC + 1
					case "G":
							snp.NumG = snp.NumG + 1
					case "T":
							snp.NumT = snp.NumT + 1
					}
					switch alt2 {
					case "A":
							snp.NumA = snp.NumA + 1
					case "C":
							snp.NumC = snp.NumC + 1
					case "G":
							snp.NumG = snp.NumG + 1
					case "T":
							snp.NumT = snp.NumT + 1
					}
				} else if polymorphism == "0/3"{
					switch ref {
					case "A":
							snp.NumA = snp.NumA + 1
					case "C":
							snp.NumC = snp.NumC + 1
					case "G":
							snp.NumG = snp.NumG + 1
					case "T":
							snp.NumT = snp.NumT + 1
					}
					switch alt3 {
					case "A":
							snp.NumA = snp.NumA + 1
					case "C":
							snp.NumC = snp.NumC + 1
					case "G":
							snp.NumG = snp.NumG + 1
					case "T":
							snp.NumT = snp.NumT + 1
					}
				}
			}

			//sort the snps :D
			snps := []Poly{
				{"NumA", snp.NumA},
				{"NumC", snp.NumC},
				{"NumG", snp.NumG},
				{"NumT", snp.NumT},
			}
			//Closure
			sort.Slice(snps, func(i, j int) bool {
				return snps[i].Number > snps[j].Number //Change > < when major minor is wanted
			})

			highestName := snps[0].Name
			highestNumber := snps[0].Number
			//WHAT HAPPENS WHEN NONE ARE HIGHER THAN THE OTHER (BOUNDERY CONDITION)
			//Enter only if there is one type of snp that has a higher count.
			//Otherwise do not do anything, we are already at 0 everywhere
			if highestNumber > 0 {
				switch highestName {
				case "NumA":
					snp.NumA = 1
					snp.NumC = 0
					snp.NumG = 0
					snp.NumT = 0
				case "NumC":
					snp.NumA = 0
					snp.NumC = 1
					snp.NumG = 0
					snp.NumT = 0
				case "NumG":
					snp.NumA = 0
					snp.NumC = 0
					snp.NumG = 1
					snp.NumT = 0
				case "NumT":
					snp.NumA = 0
					snp.NumC = 0
					snp.NumG = 0
					snp.NumT = 1
				}
				//fmt.Println(snps[0])
			}

			//Find the effect (synonymous, missense ...) of the snp
			/*info := entry[7]
			//fmt.Printf("%s", info)
			annStr := strings.Split(info, ";")
			ann := annStr[len(annStr)-1]
			effect := strings.Split(ann, "|")[1] //OK effect is obtained
			snp.Effect = effect
			//fmt.Printf("%s", effect)
			*/
			//Key part
			gene := entry[0]
			pos, _ := strconv.Atoi(entry[1])

			key := Key{gene, pos}

			d[key] = snp
			//nAll, _ := strconv.Atoi(entry[2])
			//nChr, _ := strconv.Atoi(entry[3])
		}
	}
	return d
}

func PopulateReference(f *os.File, d Data) Data {
	input := bufio.NewScanner(f)
	for input.Scan() {

		line := input.Text()
		firstChar := strings.Split(line, "")[0]

		if firstChar != "#" {

			//SNP part
			snp := SNP{0,0,0,0,""} //initialize snp

			entry := strings.Split(line, "\t")
			ref := entry[3]

			//assign the ref
			switch ref {
			case "A":
					snp.NumA = snp.NumA + 1
			case "C":
					snp.NumC = snp.NumC + 1
			case "G":
					snp.NumG = snp.NumG + 1
			case "T":
					snp.NumT = snp.NumT + 1
			}

			//Find the effect (synonymous, missense ...) of the snp
			//info := entry[7]
			//fmt.Printf("%s", info)
			//annStr := strings.Split(info, ";")
			//ann := annStr[len(annStr)-1]
			//effect := strings.Split(ann, "|")[1] //OK effect is obtained
			//snp.Effect = effect
			//fmt.Printf("%s", effect)

			//Key part
			gene := entry[0]
			pos, _ := strconv.Atoi(entry[1])

			key := Key{gene, pos}

			d[key] = snp
			//nAll, _ := strconv.Atoi(entry[2])
			//nChr, _ := strconv.Atoi(entry[3])
		}
	}
	return d
}

func main() {
	fPtr := flag.String("f", "nothing", "file containing focal SNPs; no missing data")
	a1Ptr := flag.String("a1", "nothing", "file containing outgroup 1 SNPs; all data for that outgroup must be present")
	a2Ptr := flag.String("a2", "nothing", "file containing outgroup 2 SNPs; all data for that outgroup must be present")
	a3Ptr := flag.String("a3", "nothing", "file containing outgroup 3 SNPs; all data must be present")
	//typePtr := flag.String("type", "synonymous_variant", "what type of SNP are you interested in? (eg. missense_variant)")
	//Parse the flags before doing anything!!!
	flag.Parse()

	//variant := *typePtr

	fileFocal, _ := os.Open(*fPtr) // focal allele
	fileAncestral1, _ := os.Open(*a1Ptr) // ancestral 1 allele
	fileAncestral2, _ := os.Open(*a2Ptr) // ancestral 2 allele
	fileAncestralRef, _ := os.Open(*a3Ptr) // ancestral allele

	dataFocal := make(Data)
	dataAncestral1 := make(Data)
	dataAncestral2 := make(Data)
	dataAncestralRef := make(Data)

	dataFocal = PopulateFocal(fileFocal, dataFocal)
	dataAncestral1 = PopulateAncestral(fileAncestral1, dataAncestral1)
	dataAncestral2 = PopulateAncestral(fileAncestral2, dataAncestral2)
	dataAncestralRef = PopulateReference(fileAncestralRef, dataAncestralRef)

	fileFocal.Close()
	fileAncestral1.Close()
	fileAncestral2.Close()
	fileAncestralRef.Close()


	//fileAncestralRef.Close()
	//fileChange.Close()
	/*for k, v := range dataAncestralRef {
    fmt.Printf("key[%s] value[%s]\n", k, v)
	}*/

	/*for k, v := range dataFocal {
		fmt.Printf("%d,%d,%d,%d %d,%d,%d,%d %d,%d,%d,%d %d,%d,%d,%d \n", v.NumA, v.NumC, v.NumG, v.NumT,
			dataAncestral1[k].NumA, dataAncestral1[k].NumC, dataAncestral1[k].NumG, dataAncestral1[k].NumT,
			dataAncestral2[k].NumA, dataAncestral2[k].NumC, dataAncestral2[k].NumG, dataAncestral2[k].NumT,
			dataAncestralRef[k].NumA, dataAncestralRef[k].NumC, dataAncestralRef[k].NumG, dataAncestralRef[k].NumT)
	}*/
	//Format so that SNPs are in order.

	var keys []Key
	for k := range dataFocal {
	    keys = append(keys, k)
	}

	sort.Slice(keys, func(i, j int) bool {
	    if keys[i].Gene < keys[j].Gene {
	        return true
	    }
	    if keys[i].Gene > keys[j].Gene {
	        return false
	    }
	    return keys[i].Pos < keys[j].Pos
	})
	for _, k := range keys {
		//if dataFocal[k].Effect == variant {
			fmt.Printf("%d,%d,%d,%d %d,%d,%d,%d %d,%d,%d,%d %d,%d,%d,%d \n",
				dataFocal[k].NumA, dataFocal[k].NumC, dataFocal[k].NumG, dataFocal[k].NumT,
				dataAncestral1[k].NumA, dataAncestral1[k].NumC, dataAncestral1[k].NumG, dataAncestral1[k].NumT,
				dataAncestral2[k].NumA, dataAncestral2[k].NumC, dataAncestral2[k].NumG, dataAncestral2[k].NumT,
				dataAncestralRef[k].NumA, dataAncestralRef[k].NumC, dataAncestralRef[k].NumG, dataAncestralRef[k].NumT)
				//k.Gene, k.Pos)
			//fmt.Printf("gene[%s] pos[%d]\n", k.Gene, k.Pos)
		//}
	}
}
