package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func partTwo() {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(string(content[1]))
	text := string(content)
	// for i := 0; i < (len(text) - 13); i++ {
	// 	puffer := string(text[i]) + string(text[i+1]) + string(text[i+2]) + string(text[i+3])
	// 	//fmt.Println(puffer)
	// 	break
	// }
	//fmt.Println(text)
	for i := 0; i < (len(text) - 14); i++ {
		puffer := string(text[i]) + string(text[i+1]) + string(text[i+2]) + string(text[i+3]) + string(text[i+4]) + string(text[i+5]) + string(text[i+6]) + string(text[i+7]) + string(text[i+8]) + string(text[i+9]) + string(text[i+10]) + string(text[i+11]) + string(text[i+12]) + string(text[i+13])
		//fmt.Println(puffer)
		num := 0
		num = strings.Count(puffer, string(puffer[0]))
		num += strings.Count(puffer, string(puffer[1]))
		num += strings.Count(puffer, string(puffer[2]))
		num += strings.Count(puffer, string(puffer[3]))
		num += strings.Count(puffer, string(puffer[4]))
		num += strings.Count(puffer, string(puffer[5]))
		num += strings.Count(puffer, string(puffer[6]))
		num += strings.Count(puffer, string(puffer[7]))
		num += strings.Count(puffer, string(puffer[8]))
		num += strings.Count(puffer, string(puffer[9]))
		num += strings.Count(puffer, string(puffer[10]))
		num += strings.Count(puffer, string(puffer[11]))
		num += strings.Count(puffer, string(puffer[12]))
		num += strings.Count(puffer, string(puffer[13]))
		//fmt.Printf("%d ", num)
		if num == 14 {
			fmt.Println(i + 14)
			break
		}
	}

}
