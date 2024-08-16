package main
import (
	"fmt"
    "encoding/json"
    "os"
)
func main(){
		book1:=book{
		auther: "Shodmon",
		title: "30",
	}
		data,err:=Json.Marshal(book)
		if err!=nil{
		return fmt.Print(err)
	}
     
     fmt.Println(data)
	 
}

//func countCharacters(fileName string) (int, error){

	//file, err := os.Open("example.txt")
	//if err != nil {
	//fmt.Println("ощибка ", err)
	//return
	//}
	//defer file.Close()


//}


type book struct{
	title string
	auther string
} 


