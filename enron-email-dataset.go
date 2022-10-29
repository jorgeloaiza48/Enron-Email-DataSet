package main

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"

	//"reflect"

	//"time"
	//"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

type email struct {
	ID                        int    `json:"ID"`
	Message_ID                string `json:"Message-ID"`
	Date                      string `json:"Date"`
	From                      string `json:"from"`
	To                        string `json:"to"`
	Subject                   string `json:"subject"`
	Mime_Version              string `json:"Mime-Version"`
	Content_Type              string `json:"Content-Type"`
	Content_Transfer_Encoding string `json:"Content-Transfer-Encoding"`
	X_From                    string `json:"X-From"`
	X_To                      string `json:"X-To"`
	X_cc                      string `json:"X-cc"`
	X_bcc                     string `json:""X-bcc"`
	X_Folder                  string `json:"X-Folder"`
	X_Origin                  string `json:"X-Origin"`
	X_FileName                string `json:"X-FileName"`
	Cc                        string `json:"Cc"`
	Body                      string `json:"Body"`
}

var jSonFinal []string //array donde se guardará todos los correos en forma de objetos.

// List all folders
func list_all_folders(folder_name string) []string { //recibe como parámetro el folder "maildir".
	files, err := ioutil.ReadDir(folder_name) //"ioutil.ReadDir" extrae todos los subfolders y los guarda en "files"
	if err != nil {
		log.Fatal(err)
	}
	var list_folders []string //array donde se guardarán las subcarpetas de "maildir"
	for _, f := range files {
		list_folders = append(list_folders, f.Name()) //Guradmos el nombre de cada subfolder
	}
	return list_folders
}

// Lista cada uno de los archivos o correos
func list_files(folder_name string) []string {
	files, err := ioutil.ReadDir(folder_name) //https://golang.cafe/blog/how-to-list-files-in-a-directory-in-go.html
	if err != nil {
		log.Fatal(err)
	}
	var files_names []string //array donde se guardarán los nombres de los archivos contenidos en las subcarpetas.
	for _, file := range files {
		//files_names[i] = file.Name()
		files_names = append(files_names, file.Name())
	}
	return files_names
}
func parse_data(data_lines *bufio.Scanner, id int) email {

	var data email
	for data_lines.Scan() {
		data.ID = id
		if strings.Contains(data_lines.Text(), "Message-ID:") {
			data.Message_ID = data_lines.Text()[11:len(data_lines.Text())]
		} else if strings.Contains(data_lines.Text(), "Date:") {
			data.Date = data_lines.Text()[5:len(data_lines.Text())]
		} else if strings.Contains(data_lines.Text(), "From:") {
			data.From = data_lines.Text()[5:len(data_lines.Text())]
		} else if strings.Contains(data_lines.Text(), "To:") {
			data.To = data_lines.Text()[3:len(data_lines.Text())]
		} else if strings.Contains(data_lines.Text(), "Subject:") {
			data.Subject = data_lines.Text()[8:len(data_lines.Text())]
		} else if strings.Contains(data_lines.Text(), "Cc:") {
			data.Cc = data_lines.Text()[3:len(data_lines.Text())]
		} else if strings.Contains(data_lines.Text(), "Mime-Version:") {
			data.Mime_Version = data_lines.Text()[9:len(data_lines.Text())]
		} else if strings.Contains(data_lines.Text(), "Content-Type:") {
			data.Content_Type = data_lines.Text()[9:len(data_lines.Text())]
		} else if strings.Contains(data_lines.Text(), "Content-Transfer-Encoding:") {
			data.Content_Transfer_Encoding = data_lines.Text()[9:len(data_lines.Text())]
		} else if strings.Contains(data_lines.Text(), "X-From:") {
			data.X_From = data_lines.Text()[9:len(data_lines.Text())]
		} else if strings.Contains(data_lines.Text(), "X-To:") {
			data.X_To = data_lines.Text()[9:len(data_lines.Text())]
		} else if strings.Contains(data_lines.Text(), "X-cc:") {
			data.X_cc = data_lines.Text()[6:len(data_lines.Text())]
		} else if strings.Contains(data_lines.Text(), "X-bcc:") {
			data.X_bcc = data_lines.Text()[6:len(data_lines.Text())]
		} else if strings.Contains(data_lines.Text(), "X-Folder:") {
			data.X_Folder = data_lines.Text()[9:len(data_lines.Text())]
		} else if strings.Contains(data_lines.Text(), "X-Origin:") {
			data.X_Origin = data_lines.Text()[9:len(data_lines.Text())]
		} else if strings.Contains(data_lines.Text(), "X-FileName:") {
			data.X_FileName = data_lines.Text()[9:len(data_lines.Text())]
		} else {
			data.Body = data.Body + data_lines.Text()
		}
	}
	return data
}
func index_data(data email) {
	user := "admin"
	password := "Complexpass#123"
	auth := user + ":" + password
	bas64encoded_creds := base64.StdEncoding.EncodeToString([]byte(auth))
	index := "enronJELM"
	zinc_host := "http://localhost:4080"
	//zinc_host := "https://playground.dev.zincsearch.com"
	zinc_url := zinc_host + "/api/" + index + "/_doc"
	jsonData, _ := json.MarshalIndent(data, "", "   ") //esta línea muestra los resultados tipo JSON de forma ordenada(https://gosamples.dev/pretty-print-json/)
	jSonFinal = append(jSonFinal, string(jsonData))
	//fmt.Print("Estos es la variable datos ----->> ", jSonFinal)
	// fmt.Println(string(jsonData))
	// time.Sleep(2 * time.Second)
	req, err := http.NewRequest("POST", zinc_url, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatal("Error reading request. ", err)
	}
	// Set headers
	//req.SetBasicAuth(user, password)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Basic "+bas64encoded_creds)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	//log.Println(resp.StatusCode)
	// body, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	//fmt.Println(string(body))
	//fmt.Printf(res.Response.Status)

}
func JSONfinal(datos []string) {
	file, err := os.Create("jSonFinal.json")
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}
	file.WriteString("{")
	file.WriteString(" Enron-email : [") //investigar como colocar la palabra con las comillas.
	for index, _ := range datos {
		file.WriteString(datos[index])
		if index == len(datos)-1 {
			file.WriteString("]")
			file.WriteString("}")
		} else {
			file.WriteString(",")
		}
	}
	file.Close()
	fmt.Println("File written successfully")
}
func main() {
	path := "c:/Users/jelm4/Downloads/enron_mail_20110402/enron_mail_20110402/maildir2/"
	contador := 0 //esta variable es para crear el ID en el archivo JSON
	user_list := list_all_folders(path)
	for _, user := range user_list {
		folders := list_all_folders(path + user)
		for _, folder := range folders {
			mail_files := list_files(path + user + "/" + folder + "/")
			for _, mail_file := range mail_files {
				//fmt.Println("Indexing: " + user + "/" + folder + "/" + mail_file)
				sys_file, _ := os.Open(path + user + "/" + folder + "/" + mail_file) //abre el archivo
				lines := bufio.NewScanner(sys_file)                                  //Lee el archivo líea por línea (https://golangdocs.com/reading-files-in-golang)
				// var lineas []string
				// for lines.Scan() { //https://thedeveloperblog.com/bufio
				// 	//lineas = append(lineas, lines.Text())
				// 	fmt.Println("Líneas ---------->>> ",lines.Text())
				// }
				contador++ //cada vez que se invoque la función "parse_data" esta variable se pasa con un incremento de 1 para crear el ID de cada objeto en el JSON.
				index_data(parse_data(lines, contador))
				//defer sys_file.Close() //cierra el archivo
			}
		}
	}
	JSONfinal(jSonFinal)
	// fmt.Print("Longitud del array jSonFinal -------->>>  ", len(jSonFinal))
	// fmt.Print("Tipo de dato de jsonFinal ---->>  ", reflect.TypeOf(jSonFinal))
	//fmt.Print(list_all_folders(path))
	//fmt.Println(list_files(path))

}

//
