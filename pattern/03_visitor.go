/*
Visitor -  это поведенческий паттерн, который позволяет добавлять в программу новые операции, не изменяя классы объектов,
над которыми эти операции могут выполняться.

Используется, когда:
- Когда над объектами сложной структуры объектов надо выполнять некоторые не связанные между собой операции,
но вы не хотите «засорять» классы такими операциями.
- Когда новое поведение имеет смысл только для некоторых классов из существующей иерархии

Преимущества:
Упрощает добавление операций, работающих со сложными структурами объектов.
Объединяет родственные операции в одном классе.
Посетитель может накапливать состояние при обходе структуры элементов

Недостатки:
усложняет расширение иерархии классов, поскольку новые классы обычно требуют добавления нового метода visit для каждого посетителя


 */

package visitor

import "fmt"

// элемент
type shape interface {
    getType() string
    accept(visitor)
}

// конкретный элемент
type square struct {
    side int
}

func (s *square) accept(v visitor) {
    v.visitForSquare(s)
}

func (s *square) getType() string {
    return "Square"
}

// конкретный элемент

type circle struct {
    radius int
}

func (c *circle) accept(v visitor) {
    v.visitForCircle(c)
}

func (c *circle) getType() string {
    return "Circle"
}

// конкретный элемент
type rectangle struct {
    l int
    b int
}

func (t *rectangle) accept(v visitor) {
    v.visitForrectangle(t)
}

func (t *rectangle) getType() string {
    return "rectangle"
}

//Посетитель
type visitor interface {
    visitForSquare(*square)
    visitForCircle(*circle)
    visitForrectangle(*rectangle)
}


// Конкретный посетитель
type areaCalculator struct {
    area int
}

func (a *areaCalculator) visitForSquare(s *square) {
    // Calculate area for square.
    // Then assign in to the area instance variable.
    fmt.Println("Calculating area for square")
}

func (a *areaCalculator) visitForCircle(s *circle) {
    fmt.Println("Calculating area for circle")
}
func (a *areaCalculator) visitForrectangle(s *rectangle) {
    fmt.Println("Calculating area for rectangle")
}

// Конкретный посетитель
type middleCoordinates struct {
    x int
    y int
}

func (a *middleCoordinates) visitForSquare(s *square) {
    // Calculate middle point coordinates for square.
    // Then assign in to the x and y instance variable.
    fmt.Println("Calculating middle point coordinates for square")
}

func (a *middleCoordinates) visitForCircle(c *circle) {
    fmt.Println("Calculating middle point coordinates for circle")
}
func (a *middleCoordinates) visitForrectangle(t *rectangle) {
    fmt.Println("Calculating middle point coordinates for rectangle")
}

//func main() {
//    square := &square{side: 2}
//    circle := &circle{radius: 3}
//    rectangle := &rectangle{l: 2, b: 3}
//
//    areaCalculator := &areaCalculator{}
//
//    square.accept(areaCalculator)
//    circle.accept(areaCalculator)
//    rectangle.accept(areaCalculator)
//
//    fmt.Println()
//    middleCoordinates := &middleCoordinates{}
//    square.accept(middleCoordinates)
//    circle.accept(middleCoordinates)
//    rectangle.accept(middleCoordinates)
//}



