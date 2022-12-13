package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func partOne() {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	text := string(content)
	//fmt.Println(text)
	wald := strings.Split(text, "\n")
	//fmt.Println(wald)
	waldArray := make([][]string, 0)
	for _, v := range wald {
		waldReihe := strings.Split(v, "")
		waldArray = append(waldArray, waldReihe)
	}
	left := func(waldArray [][]string) {
		fmt.Print(waldArray)
	}
	left(waldArray)
}
