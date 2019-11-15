package main

import (
	"fmt"
)

func main() {
	m := map[string]func(string, string){
		"keyword1": func1,
		"keyword2": func2,
	}

	someOtherFunction("value", "keyword1", m["keyword1"])

	for k, v := range m {
		v("value", k)
	}
}

func someOtherFunction(a, b string, f func(string, string)) {
	f(a, b)
}

func func1(arg0 string, arg1 string) {
	fmt.Println("Function one!", arg0, arg1)
}

func func2(arg0 string, arg1 string) {
	fmt.Println("Function two!", arg0, arg1)
}
