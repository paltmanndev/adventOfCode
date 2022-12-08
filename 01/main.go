package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"reflect"
	"strconv"
	"strings"
)

func main() {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	str := strings.Split(string(content), "\r\n\r\n")
	max := len(str)
	fmt.Println(reflect.TypeOf(max))
	output := make([]int, 0)

	for i := 0; i < len(str); i++ {
		erg := 0
		res := strings.Trim(str[i], "[")
		res = strings.Trim(res, "]")
		value := strings.Split(res, "\r")

		for j := 0; j < len(value); j++ {
			value := strings.Split(res, "\r\n")
			if s, err := strconv.Atoi(value[j]); err == nil {
				erg = erg + s
			} else {
				fmt.Printf("error")
			}
		}
		output = append(output, erg)
	}
	value1, value2, value3 := 0, 0, 0
	for i := 0; i < max-1; i++ {
		if value1 < output[i] {
			value1 = output[i]
		}
	}
	fmt.Printf("value 1 is %d\n", value1)

	for i := 0; i < max-1; i++ {
		if value2 < output[i] && output[i] != value1 {
			value2 = output[i]
		}
	}
	fmt.Printf("value 2 is %d\n", value2)

	for i := 0; i < max-1; i++ {
		if value3 < output[i] && output[i] != value1 && output[i] != value2 {
			value3 = output[i]
		}
	}
	fmt.Printf("value 3 is %d\n", value3)

	fmt.Println(value1 + value2 + value3)
}
