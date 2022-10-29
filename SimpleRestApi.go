//https://www.youtube.com/watch?v=pQAV8A9KLwk       Golang, REST API CRUD

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

type task struct {
	ID      int    `json:ID`
	Name    string `json:name`
	Content string `json:Content`
}
type alltasks []task

var tasks = alltasks{
	{
		ID:      1,
		Name:    "Task  One",
		Content: "Some Content",
	},
	{
		ID:      2,
		Name:    "Task  Two",
		Content: "Some Content Two",
	},
}

func createTask(w http.ResponseWriter, r *http.Request) {
	var newTask task
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil{
		fmt.Fprintf(w,"Insert a valid task")
	}
	json.Unmarshal(reqBody, &newTask)
	newTask.ID = len(tasks)+1
	tasks = append(tasks, newTask)
	w.Header().Set("Content-Type", "application/json") //cabecera de un objeto json
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTask)

}
func updateTask(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	taskID, err := strconv.Atoi(vars["id"])
	var updateTask task
	if err != nil{
		fmt.Fprintf(w,"Invalid ID")
	}
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil{
		fmt.Fprintf(w, "Entre un dato válido")
	}
	json.Unmarshal(reqBody, &updateTask)
	for i, t := range tasks{
		if t.ID == taskID{
			tasks = append(tasks[:i], tasks[i+1:]...)
			updateTask.ID = taskID
			tasks = append(tasks, updateTask)
			fmt.Fprintf(w, "La tarea con ID %v ha sido actualizada satisfactoriamente ",taskID)
		}
	}
}
func deleteTask(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	taskID, err := strconv.Atoi(vars["id"])
	if err != nil{
		fmt.Fprintf(w, "Invalid ID")
		return
	}
	for i, task := range tasks{
		if task.ID == taskID{
			tasks = append(tasks[:i], tasks[i + 1:]...) //esta línea conserva los elementos que están antes y después del elemento encontrado. Ejemplo [a,b,c], si el elemento a eliminar es b entonces quedaría [a,c]
			fmt.Fprintf(w,"La tarea con el ID %v ha sido eliminada", taskID) //%v muestra la variable en cuestión
		}
	}
}
func getTask(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	taskID, err := strconv.Atoi(vars["id"]) //Atoi convierte un string en entero
	if err != nil{
		fmt.Fprintf(w, "Invalid ID")
		return
	}
	for _, task := range tasks{ //Este ciclo significa que recorre elemento a elemento de la lista "tasks"
		if task.ID == taskID {
			w.Header().Set("Content-Type", "application/json") //cabecera de un objeto json
			json.NewEncoder(w).Encode(task) //esta línea devuelve los datos en formato json
		}
	}	
}
func getTasks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(tasks)
}
func indexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to my API")
}
func main(){
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", indexRoute)
	router.HandleFunc("/tasks", getTasks).Methods("GET")
	router.HandleFunc("/tasks", createTask).Methods("POST")
	router.HandleFunc("/tasks/{id}", getTask).Methods("GET")
	router.HandleFunc("/tasks/{id}", deleteTask).Methods("DELETE")
	router.HandleFunc("/tasks/{id}", updateTask).Methods("PUT")
	log.Fatal(http.ListenAndServe(":3000", router))

}
