package sort

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"
)

// 测试排序
// infile 输入文件
// outfile 输出文件
// algorithm 算法
func readValues(infile string)(values []int, err error) {
	file, err := os.Open(infile)
	if err != nil {
		fmt.Println("打开文件失败:", infile)
		return
	}

	defer file.Close()

	br := bufio.NewReader(file)

	values = make([]int, 0)

	for {
		line, inprefix, err1 := br.ReadLine()

		if err1 != nil {
			if err1 != io.EOF {
				err = err1
			}
			break
		}

		if inprefix {
			fmt.Println("行太长了")
			break
		}

		str  := string(line)

		value, err1 := strconv.Atoi(str)
		if err1 != nil {
			err = err1
			return
		}

		values = append(values, value)
	}
	return
}

func writeValues(values []int, outfile string) error {
	file, err := os.Create(outfile)
	if err != nil {
		fmt.Println("创建输出文件失败:", outfile)
		return err
	}

	defer file.Close()

	for _, value := range values {
		str := strconv.Itoa(value)
		file.WriteString(str + "\n")
	}
	return nil
}

func TestSort(inputfile *string, outfile *string, algorithm *string) {
	values, err := readValues(*inputfile)
	if err != nil {
		fmt.Println(err)
		return
	}

	t1 := time.Now()

	switch *algorithm {
	case "qsort":
		QuickSort(values)
	case "bubblesort":
		Bubblesort(values)
	default:
		fmt.Println("输入的排序算法未知：", algorithm)
	}
	t2 := time.Now()

	fmt.Println("排序完事了， 耗时：", t2.Sub(t1))

	writeValues(values, *outfile)
}