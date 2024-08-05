package main

func main() {

}

//lib:=Library{}
//	book := Book{
//		title:  "Shodmon",
//		author: "Go",
//	}
//	fmt.Println(book.GetDetails())
//
//}
//
//type Book struct {
//	title  string
//	author string
//}
//
//func (b Book) GetDetails() string {
//
//	return fmt.Sprint(b.title, b.author)
//
//}
//
//type Library struct {
//	Book []Book
//}
//
//func (l Library) AddBook(book Book) {
//	l.Book = append(l.Book, book)
//}

//⦁ Оценки студента
//⦁ Описание: Реализуйте структуру Student с полем grades (срез
//оценок). Реализуйте метод AddGrade, добавляющий оценку, и метод
//AverageGrade, возвращающий среднее значение оценок.
//⦁ Методы:
//⦁ AddGrade(grade int)
//⦁ AverageGrade() float64

type Student struct {
	grades []int
}

func (s *Student) Add(grade int) {
	s.grades = append(s.grades, grade)
}

func (s *Student) Average() float64 {
	sum := 0

	for _, grade := range s.grades {
		sum += grade
	}

	return float64(sum) / float64(len(s.grades))
}
