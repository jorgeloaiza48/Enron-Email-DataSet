//https://www.youtube.com/watch?v=h62mSkCjaEM     Data Analysis with Golang ( Go4DataScience )

package main

import (
	"fmt"
	"log"
	"os"

	//"github.com/kniren/gota"
	"github.com/go-gota/gota/dataframe"
)

func main(){
	csvfile, err := os.Open("emails.csv")
	if err != nil{
	log.Fatal(err)
	}
	df := dataframe.ReadCSV(csvfile)
	fmt.Println(df)
	fmt.Println("The file was read successfully")	
}