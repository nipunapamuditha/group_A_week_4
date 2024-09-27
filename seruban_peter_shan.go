package main

import (
	"fmt"
	"strings"
)

func SerubanPeterShan(i int) {
	for j := 0; j < i; j++ {
		fmt.Println(strings.Repeat("*", j))
	}
}
