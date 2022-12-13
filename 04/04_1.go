package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func partOne() {
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
		if (numArray[0] <= numArray[2] && numArray[1] >= numArray[3]) || (numArray[2] <= numArray[0] && numArray[3] >= numArray[1]) {
			//fmt.Printf("Punkt\n")
			erg++

		}

		//fmt.Printf("I = %d, num = %s, str = %s\n", i, str)
		//fmt.Println(str[i+1])
		/*for i := 0; i < len(str); i++ {
			str := strings.Split(str[i], "-")
			fmt.Printf("I = %d, str = %s\n", i, str)
			numArray[i], _ = strconv.Atoi(str[0])
			numArray[i+1], _ = strconv.Atoi(str[1])
			fmt.Println(numArray)
		}*/
	}
	fmt.Println(erg)
}
