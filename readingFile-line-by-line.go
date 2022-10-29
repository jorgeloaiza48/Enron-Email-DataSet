package main
 
import (
    "fmt"
    "os"
    "bufio"
)
 
func main() {
    file, err := os.Open("texto.txt")
    if err != nil {
        fmt.Println(err)
    }
    defer file.Close()
 
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        fmt.Println(scanner.Text())
    }
 
    if err := scanner.Err(); err != nil {
        fmt.Println(err)
    }
}