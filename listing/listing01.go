// 1. Что выведет программа? Объяснить вывод программы.

package listing

import "fmt"

func main() {

	a := [5]int{76, 77, 78, 79, 80}
	var b []int = a[1:4] // создаем слайс, включающий элементы массива а с 1 по 3 индекс
	fmt.Println(b)       //  [77 78 79]
}
