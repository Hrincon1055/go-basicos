package main

import (
	"fmt"

	"github.com/golang/variables"
)

func main() {
	estato, texto := variables.ConverToText(1222)
	fmt.Println("estado = ", estato, "text = ", texto)
}
