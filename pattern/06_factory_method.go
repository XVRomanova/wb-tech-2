package factory_method

import "fmt"

/*

В Golang можно реализовать лишь базовую версию паттерна - простая фабрика.

Фабричный метод —это класс, в котором есть один метод с большим условным оператором, выбирающим создаваемый продукт.

Применимость:
- Когда заранее неизвестны типы и зависимости объектов, с которыми должен работать ваш код.
 Фабричный метод отделяет код производства продуктов от остального кода, который эти продукты использует.

- Когда вы хотите дать возможность пользователям расширять части вашего фреймворка или библиотеки.

Преимущества:
- Избавляет класс от привязки к конкретным классам продуктов.
- Выделяет код производства продуктов в одно место, упрощая поддержку кода.
- Упрощает добавление новых продуктов в программу.
- Реализует принцип открытости/закрытости.


*/

//  Интерфейс продукта
type iGun interface {
    setName(name string)
    setPower(power int)
    getName() string
    getPower() int
}

type gun struct {
    name  string
    power int
}

func (g *gun) setName(name string) {
    g.name = name
}

func (g *gun) getName() string {
    return g.name
}

func (g *gun) setPower(power int) {
    g.power = power
}

func (g *gun) getPower() int {
    return g.power
}

// Конкретный продукт
type ak47 struct {
    gun
}

func newAk47() iGun {
    return &ak47{
        gun: gun{
            name:  "AK47 gun",
            power: 4,
        },
    }
}

// Конкретный продукт
type musket struct {
    gun
}

func newMusket() iGun {
    return &musket{
        gun: gun{
            name:  "Musket gun",
            power: 1,
        },
    }
}

// Фабрика
func getGun(gunType string) (iGun, error) {
    if gunType == "ak47" {
        return newAk47(), nil
    }
    if gunType == "musket" {
        return newMusket(), nil
    }
    return nil, fmt.Errorf("Wrong gun type passed")
}
