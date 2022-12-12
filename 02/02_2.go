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
	const X = 1
	const Y = 2
	const Z = 3
	const win = 6
	const lost = 0
	const draw = 3
	punkte_ges := 0

	str := strings.Split(string(content), "\n")
	//fmt.Println(len(str))
	for i := 0; i < len(str); i++ {
		//fmt.Printf(str[i])
		if strings.Contains(str[i], "A") && strings.Contains(str[i], "X") {
			punkte := Z + lost
			punkte_ges = punkte_ges + punkte
			//fmt.Printf(" Verloren %d(%d + %d) \n", punkte, Z, lost)
		}
		if strings.Contains(str[i], "A") && strings.Contains(str[i], "Y") {
			punkte := X + draw
			punkte_ges = punkte_ges + punkte
			//fmt.Printf(" Unentschieden %d(%d + %d) \n", punkte, X, draw)
		}
		if strings.Contains(str[i], "A") && strings.Contains(str[i], "Z") {
			punkte := Y + win
			punkte_ges = punkte_ges + punkte
			//fmt.Printf(" Gewonnen %d(%d + %d) \n", punkte, Y, win)
		}
		if strings.Contains(str[i], "B") && strings.Contains(str[i], "X") {
			punkte := X + lost
			punkte_ges = punkte_ges + punkte
			//fmt.Printf(" Verloren %d(%d + %d) \n", punkte, X, lost)
		}
		if strings.Contains(str[i], "B") && strings.Contains(str[i], "Y") {
			punkte := Y + draw
			punkte_ges = punkte_ges + punkte
			//fmt.Printf(" Unentschieden %d(%d + %d) \n", punkte, Y, draw)
		}
		if strings.Contains(str[i], "B") && strings.Contains(str[i], "Z") {
			punkte := Z + win
			punkte_ges = punkte_ges + punkte
			//fmt.Printf(" Gewonnen %d(%d + %d) \n", punkte, Z, win)
		}
		if strings.Contains(str[i], "C") && strings.Contains(str[i], "X") {
			punkte := Y + lost
			punkte_ges = punkte_ges + punkte
			//fmt.Printf(" Verloren %d(%d + %d) \n", punkte, Y, lost)
		}
		if strings.Contains(str[i], "C") && strings.Contains(str[i], "Y") {
			punkte := Z + draw
			punkte_ges = punkte_ges + punkte
			//fmt.Printf(" Unentschieden %d(%d + %d) \n", punkte, Z, draw)
		}
		if strings.Contains(str[i], "C") && strings.Contains(str[i], "Z") {
			punkte := X + win
			punkte_ges = punkte_ges + punkte
			//fmt.Printf(" Gewonnen %d(%d + %d) \n", punkte, X, win)
		}
	}
	fmt.Println(punkte_ges)
}
