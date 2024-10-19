package myfunctions

import "strings"

func BytesToAsciiMap(style []byte) map[int]string {
	chars := make(map[int]string)
	intValue := 32
	splitedStyle := strings.Split(string(style[1:]), "\n\n")
	//Range char by char standard to fill each ascii separatly in the map.
		for _, v := range splitedStyle {
			chars[intValue] = v
			intValue++
		}
	return chars
}
