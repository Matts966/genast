package main

import "fmt"

func a() {
	var x = []interface{}{"A", "b"}
	var y = make([]interface{}, 2, 4)
	fmt.Println(x, y, []interface{}{1, 3, 3})
}
