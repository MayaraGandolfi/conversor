package main

import "fmt"

const conversaoK = 273

func main() {

	tempEbulicaoK := 373
	tempEbulicaoC := tempEbulicaoK - conversaoK
	fmt.Printf("A temperatura de ebulição em Kelvin é %d K e a temperatura de ebulicao em Celsius é %dº C.", tempEbulicaoK, tempEbulicaoC)

}
