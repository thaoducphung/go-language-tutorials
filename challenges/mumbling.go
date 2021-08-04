package main

import ("strings","fmt")

func Accum(s string) string {
	result :=  ""
	for index, c := range s {
		result += strings.ToUpper(string(c))
		result += "-"

	}
}
