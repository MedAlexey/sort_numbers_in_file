package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
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
	for _, letter := range fileText {
		if string(letter) != " " {
			if number, err := strconv.Atoi(string(letter)); err == nil {
				result = append(result, number)
			} else {
				fmt.Println("The file contains a foreign character:", string(letter))
			}

		}
	}

	sort(result)

	for _, number := range result {
		fmt.Print(number, " ")
	}
}
