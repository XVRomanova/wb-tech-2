package facade

/*
Фасад скрывает сложные детали реализация для простоты ипользования клиентом, предосталяет простой интерфейс к сложной системе.
Когда применять: Когда нужно представить простой или урезанный интерфейс к сложной подсистеме. Когда нужно разложить подсистему на отдельные слои.

+
Изоляция клиентов от компонентов сложной подсистемы

-
Фасад может стать божественным объектом

Пример:
мессенджер
 */

import (
	"strings"
)

// NewMan creates man.
func NewMan() *Man {
	return &Man{
		house: &House{},
		tree:  &Tree{},
		child: &Child{},
	}
}

// Man implements man and facade.
type Man struct {
	house *House
	tree  *Tree
	child *Child
}

// Todo returns that man must do.
func (m *Man) Todo() string {
	result := []string{
		m.house.Build(),
		m.tree.Grow(),
		m.child.Born(),
	}
	return strings.Join(result, "\n")
}

// House implements a subsystem "House"
type House struct {
}

// Build implementation.
func (h *House) Build() string {
	return "Build house"
}

// Tree implements a subsystem "Tree"
type Tree struct {
}

// Grow implementation.
func (t *Tree) Grow() string {
	return "Tree grow"
}

// Child implements a subsystem "Child"
type Child struct {
}

// Born implementation.
func (c *Child) Born() string {
	return "Child born"
}
