package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func partOne() {
	container := [][]string{
		{"L", "N", "W", "T", "D"},
		{"C", "P", "H"},
		{"W", "P", "H", "N", "D", "G", "M", "J"},
		{"C", "W", "S", "N", "T", "Q", "L"},
		{"P", "H", "C", "N"},
		{"T", "H", "N", "D", "M", "W", "Q", "B"},
		{"M", "B", "R", "J", "G", "S", "L"},
		{"Z", "N", "W", "G", "V", "B", "R", "T"},
		{"W", "G", "D", "N", "P", "L"},
	}
	fmt.Println(container[4][2])
	content, err := ioutil.ReadFile("input.txt")
	var input []string
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(content))
	for _, v := range strings.Split(string(content), "\n") {

		input = strings.Split(v, " ")
		fmt.Println(input)
		amount1, _ := strconv.Atoi(input[1])
		amount2, _ := strconv.Atoi(input[3])
		amount2--
		amount3, _ := strconv.Atoi(input[5])
		amount3--
		//Erste Zahle gibt an wie viele Buchstaben aus eine Reihe
		//Zweite Zahl gibt an welche Reihe(container[Reihe][Stelle])
		//Dritte Zahl gibt an in welche Reihe verschoben werdne soll
		move := container[amount2][len(container[amount2])-amount1:]
		container[amount2] = container[amount2][:len(container[amount2])-amount1]
		fmt.Println(len(move))
		for i := len(move); i > 0; i-- {
			container[amount3] = append(container[amount3], move[i-1])
			fmt.Println(container)
		}

		//container[4] = append(container[4], move...)
		//copy(move, container[5][:6])
		fmt.Println(move)
		fmt.Println(container)
	}
	for _, v := range container {
		fmt.Print(v[len(v)-1])
	}

}
