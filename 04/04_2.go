package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func partTwo() {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	erg := 0
	str := strings.Split(string(content), "\n")
	//fmt.Println(str[0])
	for i := range str {
		str := strings.Split(str[i], ",")
		numArray := make([]int, 0)

		for i := range str {
			str := strings.Split(str[i], "-")
			for _, s := range str {
				num, _ := strconv.Atoi(s)
				numArray = append(numArray, num)
			}
			//fmt.Println(str)
		}
		//fmt.Println(numArray)
		if (numArray[1] >= numArray[2]) && (numArray[3] >= numArray[0]) {
			//fmt.Printf("Punkt\n")
			erg++

		}
	}
	fmt.Println(erg)
}
