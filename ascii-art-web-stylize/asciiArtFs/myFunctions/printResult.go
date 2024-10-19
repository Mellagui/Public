package myfunctions

import "fmt"

func PrintResult(result []string) {
	//Printing line by line the result.
	for _, v := range result {
		fmt.Println(v)
	}
}
