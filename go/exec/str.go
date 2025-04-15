package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	arr := [...]int{3, 4, 5, 2, 7, 3, 888, 55, 78}
	//news := arr[:]
	n := make([]int, len(arr))
	copy(n, arr[:])
	sort.Ints(n)
	fmt.Println(n)
	fmt.Println(arr)

	str := "aaabbddcc"
	var result strings.Builder
	count := 1
	for i := 1; i < len(str); i++ {
		if str[i] == str[i-1] {
			count++
		} else {
			v := str[i-1]
			result.WriteByte(v)
			cou := strconv.Itoa(count)
			result.WriteString(cou)
			count = 1
		}
	}
	last := str[len(str)-1]
	result.WriteByte(last)
	cous := strconv.Itoa(count)
	result.WriteString(cous)
	fmt.Println(result.String())
}
