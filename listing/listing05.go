// 5. Что выведет программа? Объяснить вывод программы.

/*
В err записывается указатель со значением nil. Это указатель на тип customError, который реализует интерфейс error, поэтому этот указатель можно записать
в переменную типа error - компилятор позволит нам это. Но при проверке на nil за nil засчитается только значение записанное
по типу интерфейса error - то есть если бы функция test в качестве возвращаемого значения в сигнатуре функции указывала бы интерфейс error,
тогда все бы работало как предполагается и в консоль было бы выведено ok.
*/
package listing

type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}

func test() *customError {
	{
		// do something
	}
	return nil
}

func main() {
	var err error
	err = test()
	if err != nil {
		println("error")
		return
	}
	println("ok")
}
