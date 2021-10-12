/*
Создать программу печатающую точное время с использованием NTP библиотеки.
Инициализировать как go module.
Использовать библиотеку github.com/beevik/ntp.
Написать программу печатающую текущее время / точное время с использованием этой библиотеки.

1. Программа должна быть оформлена с использованием как go module.
2. Программа должна корректно обрабатывать ошибки библиотеки: распечатывать их в STDERR и
возвращать ненулевой код выхода в OS.
3. Программа должна проходить проверки go vet и golint.
 */

package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"os"
)

const host = "0.beevik-ntp.pool.ntp.org"

func main() {
	if err := printCurrentTime(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

//printCurrentTime() печатает текущее время, в случае возникновения ошибок возвращает их
func printCurrentTime() error {
	time, err := ntp.Time(host)
	if err != nil {
		return err
	}

	fmt.Println(time)

	return nil
}
