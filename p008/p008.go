package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func find_product(x string) int {
	product := 1
	x_split := strings.Split(x, "")
	for i := 0; i < len(x_split); i += 1 {
		asInt, _ := strconv.Atoi(x_split[i])
		product *= asInt
	}
	return product
}

func main() {

	dat, _ := ioutil.ReadFile("data.txt")
	dat_str := string(dat)

	max_set_length := 13
	max_product := 0
	for i := 0; i < len(dat_str)-13; i += 1 {
		next_subset := dat_str[i : i+max_set_length]
		if p := find_product(next_subset); p > max_product {
			max_product = p
		}
	}
	fmt.Println("Answer:", max_product)
}
