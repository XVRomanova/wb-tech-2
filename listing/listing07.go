// 7. Что выведет программа? Объяснить вывод программы.
/*
Выведет числа от 1 до 8, после нули.
Функция merge() читает значения из каналов, которые были закрыты, когда в них перестали писать значения. Соответственно на строчках 36 и 38
мы читаем из закрытых каналов, значения из которых по дефолту равны нулю.
*/

package listing

import (
	"fmt"
	"math/rand"
	"time"
)

func asChan(vs ...int) <-chan int {
	c := make(chan int)

	go func() {
		for _, v := range vs {
			c <- v
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}

		close(c)
	}()
	return c
}

func merge(a, b <-chan int) <-chan int {
	c := make(chan int)
	go func() {
		for {
			select {
			case v := <-a:
				c <- v
			case v := <-b:
				c <- v
			}
		}
	}()
	return c
}

func main() {

	a := asChan(1, 3, 5, 7)
	b := asChan(2, 4, 6, 8)
	c := merge(a, b)
	for v := range c {
		fmt.Println(v)
	}
}
