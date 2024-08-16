package main
import (
	"bufio"
    "fmt"
    "io"
    "os"
	 )
func main(){
	n, err := countLines("test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(n)
}

func countLines(fileName string) (int, error){
	file,err := os.Open("fileName.txt")
  	if err != nil {
   		return 0, fmt.Println("Ошибка",err)
	}
  	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	k := 0

	for scanner.Scan() {
		k++
	}

	return k, nil
}

