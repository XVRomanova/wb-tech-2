// 4. Что выведет программа? Объяснить вывод программы.

package listing

/*
Программа выведет числа от 0 до 9. После будет deadlock, поскольку мы не закрыли канал. На строчке 17 из канала продолжают читать, но туда ничего не пишется.
*/

func main() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()

	for n := range ch {
		println(n)
	}
}
