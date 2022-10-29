package main

import(
	"fmt"
	"bufio"
	"os"
)

func main() {
	// nombre:="Jorge"
	// fmt.Println("Mi nombre es ", nombre)
	
	// var edad int
	// fmt.Println("Ingresa tu edad: ")
	// fmt.Scanf("%d\n",&edad)
	// fmt.Printf("Tu edad es %d ",edad)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Ingresa tu nombre: ")
	nombre,err:= reader.ReadString('\n') //la variale err devuelve un error si hay uno.
	if err !=nil{
		fmt.Println(err)
	}else{ fmt.Println("Hola "+nombre) }
}