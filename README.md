# Saludos en Golang

Este paquete proporciona una forma simple de obtener saludos personalizados en Go.

## Intalacion
Ejecuta el siguiente comando para instalar el paquete:

```bash
go get -u github.com/roviar04/greetings
```

## Uso
Aqui tienes un ejemplo de como utilizar el paquete en tu codigo:

```go
package main

import (
	"fmt"
	"log"

	"github.com/roviar04/greetings"
)

func main() {
	//Manejar el error del modulo importado
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	//Agregar varios nombres para mandar saludos
	names := []string{"Roger", "Ferny", "Cesar", "Cheva perras"}
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(messages)
}
```