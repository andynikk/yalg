package basic

import (
	"os"
	"strings"
)

func InputTxt() ([]byte, error) {
	res, err := os.ReadFile("input.txt")
	return res, err
}

func OutputTxt(r *[]string, sep string) error {

	file, err := os.Create("output.txt")
	if err != nil {
		return err
	}
	file.WriteString(strings.Join(*r, sep))
	return nil

}

func OutputStr(r *string) error {

	file, err := os.Create("output.txt")
	if err != nil {
		return err
	}
	file.WriteString(*r)
	return nil

}
