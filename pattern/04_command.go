package command

import "fmt"

/*
Команда - поведенческий паттерн проектирования, который превращает запросы в объекты,
позволяя передавать их как аргументы при вызове методов, ставить запросы в очередь, логировать их, а также поддерживать отмену операций.

Применение:
- когда небходимо инкапсулирование запроса в виде объекта

Преимущества:
- Убирает прямую зависимость между объектами, вызывающими операции, и объектами, которые их непосредственно выполняют.
- Позволяет реализовать простую отмену и повтор операций.
- Позволяет реализовать отложенный запуск операций.
- Позволяет собирать сложные команды из простых.
- Реализует принцип открытости/закрытости.

Недостатки:
- Усложняет код программы из-за введения множества дополнительных классов.

Пример:
Шаблон «Команда» можно использовать и для реализации системы на основе транзакций.
То есть системы, в которой вы сохраняете историю команд по мере их выполнения. Если последняя команда выполнена успешно, то всё хорошо.
В противном случае система итерирует по истории и делает undo для всех выполненных команд.
 */

// интерфейс команды
type command interface {
    execute()
}

// отправитель
type button struct {
    command command
}

func (b *button) press() {
    b.command.execute()
}

// интерфейс получателя
type device interface {
    on()
    off()
}

// Конкретная команда
type onCommand struct {
    device device
}

func (c *onCommand) execute() {
    c.device.on()
}

 // Конкретная команда
type offCommand struct {
    device device
}

func (c *offCommand) execute() {
    c.device.off()
}

// конкретный получатель
type tv struct {
    isRunning bool
}

func (t *tv) on() {
    t.isRunning = true
    fmt.Println("Turning tv on")
}

func (t *tv) off() {
    t.isRunning = false
    fmt.Println("Turning tv off")
}

//func main() {
//    tv := &tv{}
//
//    onCommand := &onCommand{
//        device: tv,
//    }
//
//    offCommand := &offCommand{
//        device: tv,
//    }
//
//    onButton := &button{
//        command: onCommand,
//    }
//    onButton.press()
//
//    offButton := &button{
//        command: offCommand,
//    }
//    offButton.press()
//}


