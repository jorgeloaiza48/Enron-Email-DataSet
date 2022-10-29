package main

import 
	"net/http" //paquete para http

func main(){
	
	//routes
	http.HandleFunc("/",homeHandler)
	http.HandleFunc("/contact",contactHandler)
	
	http.ListenAndServe(":3000",nil) //servidor escuchando en puerto 3000
}

func homeHandler(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Mi primer servidor con go"))
}
func contactHandler(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Contact Page"))
}