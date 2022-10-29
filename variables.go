package main

//Paquetes
import (
	"fmt"
	"strconv"
)

func main() {
	//var x,y,z int
	edad := 23                     //declara la variable de tipo entero y se le asigna un entero
	nombre := "Jorge"              //declara la variable de tipo cadena y le asigna un valor.
	edad_str := strconv.Itoa(edad) //strconv.Itoa sirve para convertir un entero a cadena de texto
	//strconv.Atoi convierte de texto a entero. Esto retorna dos valores entonces se debe colocar un guión bajo para ignorar ese segundo valor(edad_int,_:=strconv.Atoi(edad))
	//edad:="22" 						//es texto
	//edad_int,_:= strconv.Atoi(edad) //edad_int queda como entero
	fmt.Println("Soy " + nombre + " y tengo " + edad_str)

}
