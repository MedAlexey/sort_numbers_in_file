package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func checkErr(e error) {
	if e != nil {
		fmt.Println("Error opening file")
		os.Exit(1)
	}
}

func sort(result []int) {
	for j := 0; j < len(result)-1; j++ {
		for i := 0; i < len(result)-1-j; i++ {
			if result[i] > result[i+1] {
				tmp := result[i]
				result[i] = result[i+1]
				result[i+1] = tmp
			}
		}
	}
}

func main() {

	var fileName string
	fmt.Scan(&fileName)
	file, err := ioutil.ReadFile(fileName)
	checkErr(err)

	var result []int
	fileText := string(file)
	numbers := strings.Split(fileText, " ")

	for _, number := range numbers {
		if number, err := strconv.Atoi(number); err == nil {
			result = append(result, number)
		}
	}

	sort(result)

	for _, number := range result {
		fmt.Print(number, " ")
	}
}
