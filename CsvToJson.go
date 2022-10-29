//https://www.golanglearn.com/golang-tutorials/converting-csv-data-to-json-in-golang/

package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

type Employee struct {
	ID             string `json:"id"`
	EmployeeName   string `json:"employee_name"`
	EmployeeSalary int    `json:"employee_salary"`
	EmployeeAge    int    `json:"employee_age"`
}

func main() {
	csv_file, err := os.Open("sample.csv")
	if err != nil {
		fmt.Println(err)
	}
	defer csv_file.Close()

	r := csv.NewReader(csv_file)
	records, err := r.ReadAll()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var emp Employee
	var employees []Employee

	for _, rec := range records {
		emp.ID = rec[0]
		emp.EmployeeName = rec[1]
		emp.EmployeeSalary, _ = strconv.Atoi(rec[2])
		emp.EmployeeAge, _ = strconv.Atoi(rec[3])
		employees = append(employees, emp)
	}
	// Convert to JSON
	json_data, err := json.Marshal(employees)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	//print json data
	fmt.Println(string(json_data))

	//create json file
	json_file, err := os.Create("sample.json")
	if err != nil {
		fmt.Println(err)
	}
	defer json_file.Close()

	json_file.Write(json_data)
	json_file.Close()
}
