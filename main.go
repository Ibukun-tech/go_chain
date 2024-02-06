package main

import (
	"encoding/json"
	"fmt"
)

type Value struct {
	V string
	W string
}

func main() {
	var vs []*Value
	a := &Value{
		V: "WXY",
		W: "Yx",
	}
	b := &Value{
		V: "sr",
		W: "ANSN",
	}
	c := &Value{
		V: "sd",
		W: "sdesx",
	}
	d := &Value{
		V: "Oyetunji",
		W: "Favour",
	}
	vs = append(vs, a, b, c, d)
	for _, v := range vs {
		fmt.Println(v, "first range")
	}
	e, _ := json.MarshalIndent(vs, "", "")
	fmt.Println(string(e), "Second range")

}
