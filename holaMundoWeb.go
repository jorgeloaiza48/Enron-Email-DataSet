//https://www.youtube.com/watch?v=qtE6jXpuMdI&list=PLdKnuzc4h6gFmPLeous4S0xn0j9Ik2s3Y&index=23

package main

import (
	"net/http"
	"io"
	"fmt"
	
)

func main(){
	http.HandleFunc("/",handler)	
	http.ListenAndServe(":8000",nil) //servidor escuchando en puerto 8000
}
func handler(w http.ResponseWriter,r *http.Request){
	fmt.Println("Hay una nueva petición")
	io.WriteString(w,"Hola Mundo WEB")
}