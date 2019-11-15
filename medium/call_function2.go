package main

import (
	"errors"
	"fmt"
	"reflect"
)

type stubMapping map[string]interface{}

var StubStorage = stubMapping{}

func main() {
	StubStorage = map[string]interface{}{
		"funcA": funcA,
		"funcB": funcB,
	}

	resA,_ := Call("funcA", "value", "keyword1")
	var prntA string
	prntA = resA.(string)
	fmt.Println(prntA)

	resB,_ := Call("funcB", 10)
	var prntB int
	prntB = resB.(int)
	fmt.Println(prntB)
}

func Call(funcName string, params ... interface{}) (result interface{}, err error) {
	f := reflect.ValueOf(StubStorage[funcName])
	if len(params) != f.Type().NumIn() {
		err = errors.New("The number of params is out of index.")
		return
	}
	in := make([]reflect.Value, len(params))
	for k, param := range params {
		in[k] = reflect.ValueOf(param)
	}
	var res []reflect.Value
	res = f.Call(in)
	result = res[0].Interface()
	return
}

func funcA(arg0 string, arg1 string) string {
	return fmt.Sprintf("Function one! %v %v", arg0, arg1)
}

func funcB(args0 int) int {
	return args0
}
