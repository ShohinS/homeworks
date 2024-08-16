//1. Сериализация структуры в JSON
//Описание: Напишите программу, которая сериализует структуру Person в формат
//JSON и выводит результат на экран.
//Структура:
//type Person struct {
//Name string `json:"name"`
//Age int `json:"age"`
//Email string `json:"email"`
//}

//package main

//import (
//"encoding/json"
//"fmt"

//)
//func main(){
//per := Person{
//"Romish",
//35,
//"Romish1992mail.ru",

//}
//jcode,err:=json.Marshal(per)
//if err !=nil{
//fmt.Println(err)
//}
//fmt.Println(string(jcode))
//}

//type Person struct{
// Name string
//Age  int
//Email string

//}

//2. Десериализация JSON в структуру

//package main

//import (
//	"encoding/json"
//	"fmt"

//)
//func main(){
//per :=`{ "name":"Roma",
//"age":30,
//"email":"Romish1992mail.ru"
//}`

// var pers Person
//err:=json.Unmarshal([]byte(per),&pers)
//if err !=nil{
//fmt.Println(err)
// }
//fmt.Println(pers.Name, pers.Age, pers.Email)
//}

//type Person struct {
//	Name string `json:"name"`
//Age int `json:"age"`
//Email string `json:"email"`
//}

//3. Сериализация карты в JSON

//package main

//import (
	//"encoding/json"
	//"fmt"
	
//)
//func main(){
	//data := map[string]int{
		//"apples": 5,
		//"oranges": 10,
	   //}
  //jdata,err:=json.Marshal(data)
  //if err != nil{
	//fmt.Println(err)
  //}	   
   // fmt.Println(string(jdata))
 //}



 //4. Десериализация JSON в карту
 package main
 import (
	"encoding/json"
	"fmt"
	
)
func main(){
	 per:=`{
		"apples": 5,
		"oranges": 10
	   }`

	 var d map[string]int
	 err := json.Unmarshal([]byte(per),&d)
	 if err !=nil{
		fmt.Println(err)
	 }  
	   fmt.Println(d)

	 }

	 