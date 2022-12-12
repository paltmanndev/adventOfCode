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

	str := strings.Split(string(content), "\r")
	//fmt.Println(str[0])
	for i := range str {
		str = strings.Split(str[i], ",")
		fmt.Println(str[i])
		fmt.Println(str[i+1])
	}
}
