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
	//fmt.Println(string(content[1]))
	text := string(content)
	// for i := 0; i < (len(text) - 3); i++ {
	// 	puffer := string(text[i]) + string(text[i+1]) + string(text[i+2]) + string(text[i+3])
	// 	fmt.Println(puffer)
	// 	break
	// }
	//fmt.Println(text)
	for i := 0; i < (len(text) - 3); i++ {
		puffer := string(text[i]) + string(text[i+1]) + string(text[i+2]) + string(text[i+3])
		//fmt.Println(puffer)
		num := 0
		num = strings.Count(puffer, string(puffer[0]))
		num += strings.Count(puffer, string(puffer[1]))
		num += strings.Count(puffer, string(puffer[2]))
		num += strings.Count(puffer, string(puffer[3]))
		//fmt.Printf("%d ", num)
		if num == 4 {
			fmt.Println(i + 4)
			break
		}
	}

}
