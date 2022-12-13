package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func partOne() {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	text := string(content)
	fmt.Println(text)



}
