/*
Цепочка обязанностей — это поведенческий паттерн проектирования, который позволяет передавать запросы последовательно
по цепочке потенциальных обработчиков, пока один из них не обработает запрос.

Применяется:
- Когда программа должна обрабатывать разнообразные запросы несколькими способами, но заранее неизвестно,
какие конкретно запросы будут приходить и какие обработчики для них понадобятся.
- Когда важно, чтобы обработчики выполнялись один за другим в строгом порядке.

Преимущества:
- Ослабление связанности между объектами. Отправителю и получателю запроса ничего не известно друг о друге.
Клиенту неизветна цепочка объектов, какие именно объекты составляют ее, как запрос в ней передается
- Реализует принцип единственной обязанности.

Недостатки:
- Запрос может остаться никем не обработанным.

Пример:
контекстная справка в любом приложении
 */

package chain_of_resp

import "fmt"


type patient struct {
	name              string
    registrationDone  bool
    doctorCheckUpDone bool
    medicineDone      bool
    paymentDone       bool
}

//интерфейс обработчика
type department interface {
	execute(*patient)
	setNext(department)
}

// Конкретный обработчик
type doctor struct {
	next department
}

func (d *doctor) execute(p *patient) {
	if p.doctorCheckUpDone {
		fmt.Println("Doctor checkup already done")
		d.next.execute(p)
		return
	}
	fmt.Println("Doctor checking patient")
	p.doctorCheckUpDone = true
	d.next.execute(p)
}

func (d *doctor) setNext(next department) {
	d.next = next
}

// Конкретный обработчик
type medical struct {
	next department
}

func (m *medical) execute(p *patient) {
	if p.medicineDone {
		fmt.Println("Medicine already given to patient")
		m.next.execute(p)
		return
	}
	fmt.Println("Medical giving medicine to patient")
	p.medicineDone = true
	m.next.execute(p)
}

func (m *medical) setNext(next department) {
	m.next = next
}

// Конкретный обработчик
type reception struct {
    next department
}

func (r *reception) execute(p *patient) {
    if p.registrationDone {
        fmt.Println("Patient registration already done")
        r.next.execute(p)
        return
    }
    fmt.Println("Reception registering patient")
    p.registrationDone = true
    r.next.execute(p)
}

func (r *reception) setNext(next department) {
    r.next = next
}

// Конкретный обработчик
type cashier struct {
    next department
}

func (c *cashier) execute(p *patient) {
    if p.paymentDone {
        fmt.Println("Payment Done")
    }
    fmt.Println("Cashier getting money from patient patient")
}

func (c *cashier) setNext(next department) {
    c.next = next
}


//func main() {
//
//    cashier := &cashier{}
//
//    //Set next for medical department
//    medical := &medical{}
//    medical.setNext(cashier)
//
//    //Set next for doctor department
//    doctor := &doctor{}
//    doctor.setNext(medical)
//
//    //Set next for reception department
//    reception := &reception{}
//    reception.setNext(doctor)
//
//    patient := &patient{name: "abc"}
//    //Patient visiting
//    reception.execute(patient)
//}


