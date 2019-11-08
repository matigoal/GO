package main

import "fmt"


func uniqueNames(second []string, third []string) []string {

	check := make(map[string]int)
	first := append(second, third...)
	result := make([]string,0)
	for _, val := range first {
		check[val] = 1
	}

	for names, _ := range check {
		result = append(result,names)
	}

	return result
}

func main() {
	second := []string{"Ava", "Emma", "Olivia"}
	third := []string{"Olivia", "Sophia", "Emma"}
	last := uniqueNames(second, third)
	fmt.Println(last)

}
