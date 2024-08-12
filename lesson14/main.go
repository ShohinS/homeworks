package main
import ("fmt"
    "os")
func main(){


	file, err := os.OpenFile("example.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
if err != nil {
fmt.Println("Error opening file:", err)
return
}
defer file.Close()
fmt.Println("File opened successfully")

}
