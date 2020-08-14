package main

import "fmt"

type error interface {
	Error() string
}

type MyType string

func (m MyType) Error() string {
	return string(m)
}

func main() {
	var err error
	err = MyType("This is an error.")

	fmt.Println(err.Error())

}
