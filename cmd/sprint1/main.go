package main

import (
	"yalg/basic"
	"yalg/sprint1"
)

func main() {
	res, err := basic.InputTxt()
	if err != nil {
		panic(err)
	}

	r, sep := sprint1.Task6(&res)

	err = basic.OutputTxt(&r, sep)
}
