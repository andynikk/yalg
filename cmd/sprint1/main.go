package main

import (
	"yalg/basic"
	"yalg/cmd/sprint1"
)

func main() {
	res, err := basic.InputTxt()
	if err != nil {
		panic(err)
	}

	r := sprint1.Task1(&res)

	err = basic.OutputTxt(&r)
}
