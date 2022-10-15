package main

import (
	"flag"
	"fmt"
	"go_study_1/sort"
)

var infile *string = flag.String("i", "unsorte.dat", "File contains value for sorting values")
var outfile *string = flag.String("o", "sorted.dat", "File contains value for sortied values")
var algorithm *string = flag.String("a", "qsort", "sort algorithm")

func main() {
	flag.Parse()
	if infile != nil {
		fmt.Println("infile=", infile, ",outfile=", outfile, ",algorithm=", algorithm)
	}
	// 测试排序算法
	sort.TestSort(infile, outfile, algorithm)
}