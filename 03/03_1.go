package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func partOne() {
	alphabet := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	str := strings.Split(string(content), "\n")
	sum := 0
	for i := range str {
		n := len(str[i])
		str[i] = strings.TrimSuffix(str[i], "\r")
		firstHalf := str[i][:n/2]
		secondHalf := str[i][n/2:]
		for _, v := range firstHalf {
			if strings.ContainsAny(string(v), secondHalf) {
				//fmt.Printf("Found %s in %s\n", string(v), secondHalf)
				pos := strings.Index(alphabet, string(v)) + 1
				//fmt.Println(pos)
				sum += pos
				break
			}
		}
	}
	fmt.Println(sum)
}
