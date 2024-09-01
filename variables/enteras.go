package variables

import (
	"fmt"
	"time"
)

// Constantes a nivel de paquet
const Pi float32 = 3.14
const (
	x = 100
	y = 0b1010
	z = 0o2
)
const (
	Domingo = iota + 1 // Palabra reservada iota
	Lunes
	Martes
	Miercoles
	Jueves
	Virenes
)

// Variables a nivel de paquet
var estado bool = false
var sueldo float32 = 100000

func MostrarEnteros() {
	// Variables
	var intcomun int
	var fecha = time.Now()
	intde32 := int32(10)
	intde64 := int64(100)
	var firtName, lastName string
	var (
		otronombre = "Erick"
	)
	var message = "Hola mundo"
	firtName = "Henry"
	lastName = "rincon"
	fmt.Println("incomun = ", intcomun)
	fmt.Println("intde32 = ", intde32)
	fmt.Println("intde64 = ", intde64)
	fmt.Println("firtName = ", firtName)
	fmt.Println("lastName = ", lastName)
	fmt.Println("message = ", message)
	fmt.Println("otronombre = ", otronombre)
	fmt.Println("estado = ", estado)
	fmt.Println("sueldo = ", sueldo)
	fmt.Println("fecha = ", fecha)
	// Constantes
	fmt.Println("PI =", Pi)
	fmt.Println("Virenes =", Virenes)

}
