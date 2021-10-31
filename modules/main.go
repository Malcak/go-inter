package main

import (
	"github.com/donvito/hellomod"
	hellomod2"github.com/donvito/hellomod/v2"
)

func main() {
	hellomod.SayHello()
	hellomod2.SayHello("malcak ")
}

// go mod init github.com/username/module
// go get github.com/donvito/hellomod
// go mod tidy -> eliminar dependencias sin utilizar
// go mod download -json -> info de la cache de modulos
