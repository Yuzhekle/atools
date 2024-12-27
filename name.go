package main

import (
	"github.com/brianvoe/gofakeit/v6"
)

func main() {
	p := gofakeit.Person()
	println(p.FirstName)
	println(p.LastName)

}
