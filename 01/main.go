package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	content, err := ioutil.ReadFile("input.txt")
	//var zahlen = make(map[int]int)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("strings.Split(): %#v\n", strings.Split(string(content), "\r\n\r\n"))
}
