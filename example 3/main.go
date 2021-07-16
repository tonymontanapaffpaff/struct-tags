package main

import (
	"encoding/json"
	"fmt"
)

type T struct {
	F1 int `json:"f_1"`
	F2 int `json:"f_2,omitempty"`
	F3 int `json:"f_3,omitempty"`
	F4 int `json:"-"`
}

func main() {
	t := T{
		F1: 1,
		F2: 0,
		F3: 2,
		F4: 3,
	}

	b, err := json.Marshal(t)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", b)
}
