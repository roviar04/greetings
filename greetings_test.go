package greetings

import (
	"regexp"
	"testing"
)

// el objeto con puntero de tipo testing sirve para reportar el
// resultado de la prueba
func TestHelloName(t *testing.T) {
	/*
		Este codigo analiza si la variable name coincide con
		el texto asignado
	*/
	name := "Juan"

	//Expresion regular
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("Juan")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Juan") = %q, %v, quiere coincidencia para %#q, nil`, msg, err, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")

	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, quiere "", error`, msg, err)
	}
}
