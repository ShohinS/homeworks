//Задания для 13 – го урока
// 1.  Длина строки и Количество слов
//   - Описание: Реализуйте интерфейс StringProcessor, который будет содержать методы Length() и WordCount().
//   - Реализуйте структуру MyString, которая будет работать с текстом и реализует этот интерфейс.
//   - Методы:
//   - Length() int для получения длины строки.
//   - WordCount() int для подсчета количества слов
// package main

// import (
// 	"fmt"
// 	"strings"
// )

// type StringProcessor interface {
// 	Length()
// 	WordCount()
// }
// type MyString struct {
// 	str string
// }

//	func (m MyString) Length() int {
//		return len(m.str)
//	}
//
//	func (m MyString) WordCount() int {
//		return len(strings.Fields(m.str))
//	}
//
//	func main() {
//		myString := MyString{
//			str: "Humolab",
//		}
//		fmt.Println(myString.Length())
//		fmt.Println(myString.WordCount())
//	}
//
// -------------------------------------------------------------------------------------------------
// 2. Форматирование строки
//   - Описание: Реализуйте интерфейс Formatter с методами ToUpper() и ToLower().
//   - Реализуйте структуру MyFormatter, которая будет работать со строкой и реализует этот интерфейс.
//   - Методы:
//   - ToUpper() string для преобразования строки в верхний регистр.
//   - ToLower() string для преобразования строки в нижний регистр.
// package main

// import (
//
//	"fmt"
//	"strings"
//
// )
//
//	type Formatter interface{
//		ToUpper()
//		ToLower()
//	}
//
//	type MyFormatter struct{
//		str string
//	}
//
//	func (m MyFormatter) ToUpper()  {
//		fmt.Println(strings.ToUpper(m.str))
//	}
//
//	func (m MyFormatter) ToLower()  {
//		fmt.Println(strings.ToLower(m.str))
//	}
//
//	func main() {
//		variable:=MyFormatter{
//			str:"nekboy",
//		}
//		variable.ToUpper()
//		variable.ToLower()
//	}
//
// -------------------------------------------------------------------------------------------------
// 3. Работа с указателями на числа
//   - Описание: Реализуйте интерфейс PointerOperations с методами Increment() и Decrement().
//   - Реализуйте структуру IntPointer с указателем на число, которая реализует этот интерфейс.
//   - Методы:
//   - Increment() для увеличения значения числа на 1.
//   - Decrement() для уменьшения значения числа на 1.
// package main

// import (
// 	"fmt"
// )

// type PointerOperations interface {
// 	Increment()
// 	Decrement()
// }
// type IntPointer struct {
// 	intt *int
// }

// 	func (i IntPointer) Increment() {
// 		*i.intt++
// 		fmt.Println(*i.intt)
// 	}

// 	func (i IntPointer) Decrement() {
// 		fmt.Println(*i.intt - 1)
// 	}

// 	func main() {
// 		number := 1
// 		v := IntPointer{
// 			intt: &number,
// 		}
// 		v.Increment()
// 		v.Decrement()
// 	}
//
// -------------------------------------------------------------------------------------------------
// 4. Удаление пробелов в строке
//   - Описание: Реализуйте интерфейс StringCleaner с методами TrimSpaces() и RemoveSpaces().
//   - Реализуйте структуру TextCleaner, которая будет работать со строками и реализует этот интерфейс.
//   - Методы:
//   - TrimSpaces() string для удаления пробелов с начала и конца строки.
//   - RemoveSpaces() string для удаления всех пробелов из строки.
// package main

// import (
// 	"fmt"
// 	"strings"
// )

// type StringCleaner interface {
// 	TrimSpaces()
// 	RemoveSpaces()
// }
// type TextCleaner struct {
// 	str string
// }

//	func (t TextCleaner) TrimSpaces() {
//		fmt.Println(strings.TrimSpace(t.str))
//	}
//
//	func (t TextCleaner) RemoveSpaces() {
//		fmt.Println(strings.ReplaceAll(t.str, " ", ""))
//	}
//
//	func main() {
//		v := TextCleaner{
//			str: " N E K R U Z ",
//		}
//		v.TrimSpaces()
//		v.RemoveSpaces()
//	}
//
// -------------------------------------------------------------------------------------------------
// 5. Конкатенация строк
//   - Описание: Реализуйте интерфейс StringConcatenator с методами Concat() и Join().
//   - Реализуйте структуру StringJoiner, которая будет работать с массивами строк и реализует этот интерфейс.
//   - Методы:
//   - Concat() string для конкатенации всех строк в массиве.
//   - Join(separator string) string для объединения всех строк с заданным разделителем.
package main

import (
	"fmt"
	"strings"
)

// Определяем интерфейс StringConcatenator
type StringConcatenator interface {
	Concat() string
	Join(separator string) string
}

// Определяем структуру StringJoiner
type StringJoiner struct {
	strings []string
}

// Реализуем метод Concat для структуры StringJoiner
func (sj *StringJoiner) Concat() string {
	var result string
	for _, s := range sj.strings {
		result += s
	}
	return result
}

// Реализуем метод Join для структуры StringJoiner
func (sj *StringJoiner) Join(separator string) string {
	return strings.Join(sj.strings, separator)
}

func main() {
	sj := &StringJoiner{
		strings: []string{"Hello", "World", "Go", "Programming"},
	}

	fmt.Println("Concat:", sj.Concat())        // ВывоД: HelloWorldGoProgramming
	fmt.Println("Join (', '):", sj.Join(", ")) // Вывод: Hello, World, Go, Programming
}
