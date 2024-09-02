package variables

import (
	"fmt"
	"strconv"
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

func MostrarVariables() {
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
	// Tipos de datos básicos
	var entero int = 42
	var flotante float64 = 3.14
	var booleano bool = true
	var caracter rune = 'A'
	var cadena string = "Hola, Go!"

	// Tipos de datos compuestos
	var arreglo [3]int = [3]int{1, 2, 3}
	var slice []string = []string{"Go", "es", "genial"}
	var mapa map[string]int = map[string]int{"uno": 1, "dos": 2, "tres": 3}

	// Tipo struct
	type Persona struct {
		nombre string
		edad   int
	}
	var persona Persona = Persona{nombre: "Juan", edad: 30}

	// Tipo puntero
	var puntero *int = &entero

	// Impresión en consola de los tipos de datos
	fmt.Println("Entero:", entero)
	fmt.Println("Flotante:", flotante)
	fmt.Println("Booleano:", booleano)
	fmt.Println("Caracter:", string(caracter))
	fmt.Println("Cadena:", cadena)

	fmt.Println("Arreglo:", arreglo)
	fmt.Println("Slice:", slice)
	fmt.Println("Mapa:", mapa)

	fmt.Println("Struct Persona:", persona)
	fmt.Printf("Puntero (dirección de memoria): %p, Valor apuntado: %d\n", puntero, *puntero)

}
func ConverToText(numero int) (bool, string) {

	texto := strconv.Itoa(numero)
	return true, texto
}
