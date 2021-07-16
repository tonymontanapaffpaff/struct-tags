package main

import (
	"encoding/json"
	"fmt"
)

func main()  {
	type T1 struct {
		F int `json:"foo"`
	}
	type T2 struct {
		F int `json:"bar"`
	}
	t1 := T1{10}
	var t2 T2
	t2 = T2(t1)
	bytes, err := json.Marshal(t2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s\n", bytes)
}
