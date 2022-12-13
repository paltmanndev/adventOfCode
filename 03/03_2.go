package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func partTwo() {
	alphabet := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	str := strings.Split(string(content), "\n")
	sum := 0
	for i := 0; i < len(str); i = i + 3 {
		//fmt.Printf("%s\n%s\n%s\n", str[i], str[i+1], str[i+2])
		for _, v := range str[i] {
			if strings.ContainsAny(string(v), str[i+1]) && strings.ContainsAny(string(v), str[i+2]) {
				//fmt.Printf("Found %s in %s and %s\n", string(v), str[i+1], str[i+2])
				pos := strings.Index(alphabet, string(v)) + 1
				//fmt.Println(pos)
				sum += pos
				break

			}

		}
	}
	fmt.Println(sum)
}
