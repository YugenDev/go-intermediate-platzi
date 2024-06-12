package main

import (
	"github.com/donvito/hellomod"
	v2 "github.com/donvito/hellomod/v2"
)

func main() {
	hellomod.SayHello()
	v2.SayHello("Juan!")
}
