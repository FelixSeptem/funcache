package main

import (
	"fmt"
	"github.com/FelixSeptem/funcache"
)

func fib(in int) int {
	if in <= 1 {
		return 1
	}
	return fib(in-2) + fib(in-1)
}

// a stupid and useless function,just to simulate a time expensive function
func someExpensiveFun(in1, in2, in3 int) (out1, out2, out3 int, err error) {
	return fib(in1), fib(in2), fib(in3), nil
}

func UseWithCache() {
	wrapped := func(in []interface{}) ([]interface{}, error) {
		out1, out2, out3, err := someExpensiveFun(in[0].(int), in[1].(int), in[2].(int))
		return []interface{}{out1, out2, out3}, err
	}
	wrapper := funcache.CacheFunWithLru(128, wrapped, 3) // similar to use CacheFunWithLfu or CacheFunWithArc

	_, err := wrapper([]interface{}{1, 2, 3}) // equal someExpensiveFun(1,2,3)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func main() {
	UseWithCache()
}
