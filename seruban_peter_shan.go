package main

import (
	"fmt"
	"strings"
)

func SerubanPeterShan(i int) {
	//Top half of the diamond
	for j := 1; j <= i; j++ {
		fmt.Println(strings.Repeat(" ", i-j) + strings.Repeat("*", 2*j-1))
	}
	//Bottom half of the diamond
	for j := i - 1; j > 0; j-- {
		fmt.Println(strings.Repeat(" ", i-j) + strings.Repeat("*", 2*j-1))
	}
}
