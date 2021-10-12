//2. Что выведет программа? Объяснить вывод программы. Объяснить как работают defer’ы и их порядок вызовов.

package listing

import "fmt"

/*
Отложенные анонимные функции могут получать доступ и изменять именованные возвращаемые параметры внешней функции
1. Запускаем defer
2. Присваиваем х = 1
3. return
4. обновляем x на x++
*/
func test() (x int) {
	defer func() { // откладываем вызов функции до выхода из внешней функции test()
		x++
	}()
	x = 1
	return
}

/*
В случае с неименованным возвращаемым значением значение уже было скопировано с х.
Модицифкация локальной переменной х после этого не повлияет на уже скопированное возвращаемое значение.
Т.к. defer владеет ссылкой на локальную переменную, значение которой уже было скопировано и вот-вот вернётся вызывающей стороне.
*/
func anotherTest() int {
	var x int
	defer func() {
		x++
	}()
	x = 1
	return x
}

func main() {
	fmt.Println(test())        // 2
	fmt.Println(anotherTest()) // 1
}