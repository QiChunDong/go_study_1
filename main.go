package main

import (
	"flag"
	"fmt"
	"go_study_1/music"
	"go_study_1/sort"
)

var infile *string = flag.String("i", "unsorte.dat", "File contains value for sorting values")
var outfile *string = flag.String("o", "sorted.dat", "File contains value for sortied values")
var algorithm *string = flag.String("a", "qsort", "sort algorithm")

var lib *music.MusicManager
var id int = 1
var ctrl, signal chan int

func handleLibCommands(tokens []string) {
	switch tokens[1] {
	case "list":
		for i := 0; i < lib.Len(); i++ {
			e, _ := lib.Get(i)
			fmt.Println(i + 1, ":", e.Name, e.Artist, e.Source, e.Type)
		}
		// TODO 命令处理到此为止
	}
}

func main() {
	flag.Parse()
	if infile != nil {
		fmt.Println("infile=", infile, ",outfile=", outfile, ",algorithm=", algorithm)
	}
	// 测试排序算法
	sort.TestSort(infile, outfile, algorithm)
}