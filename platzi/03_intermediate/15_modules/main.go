package main

import (
	"github.com/donvito/hellomod"
	hm2 "github.com/donvito/hellomod/v2"
)

func main() {
	hellomod.SayHello()
	hm2.SayHello("Platzi")
}

// go mod init github.com/dazulu4/test
// go mod tidy => eliminar dependencias
// go mod download --json => ver ruta almacenamiento paquete
// go get github.com/donvito/hellomod => instalar un paquete externo
