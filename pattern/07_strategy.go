package strategy

import "fmt"

/*
Стратегия - это поведенческий паттерн проектирования, который определяет семейство схожих алгоритмов и помещает каждый из них в собственный класс,
после чего алгоритмы можно взаимозаменять прямо во время исполнения программы.

Паттерн может применяться:
- Когда вам нужно использовать разные вариации какого-то алгоритма внутри одного объекта.
Стратегия позволяет варьировать поведение объекта во время выполнения программы, подставляя в него различные объекты-поведения

- Когда у вас есть множество похожих классов, отличающихся только некоторым поведением.
Стратегия позволяет вынести отличающееся поведение в отдельную иерархию классов, а затем свести первоначальные классы к одному,
сделав поведение этого класса настраиваемым.

- Когда вы не хотите обнажать детали реализации алгоритмов для других классов.
Стратегия позволяет изолировать код, данные и зависимости алгоритмов от других объектов, скрыв эти детали внутри классов-стратегий.

- Когда различные вариации алгоритмов реализованы в виде развесистого условного оператора. Каждая ветка такого оператора представляет собой вариацию алгоритма.

Преимущества:
- Горячая замена алгоритмов на лету.
- Изолирует код и данные алгоритмов от остальных классов.
- Уход от наследования к делегированию.
- Реализует принцип открытости/закрытости.

Недостатки:
- Усложняет программу за счёт дополнительных классов.
- Клиент должен знать, в чём состоит разница между стратегиями, чтобы выбрать подходящую.
Изменение поведения, после отправки запроса на сервер без его перезапуска

Пример:
Выставление счета: стратегия выставления счета за использование чего-либо на основании календаря, индивидуальных показателей, прайс-листа и бонусов.
*/


// Интерфейс стратегии
type evictionAlgo interface {
	evict(c *cache)
}

// Конкретная стратегия
type fifo struct {
}

func (l *fifo) evict(c *cache) {
	fmt.Println("Evicting by fifo strtegy")
}

// Конкретная стратегия
type lru struct {
}

func (l *lru) evict(c *cache) {
	fmt.Println("Evicting by lru strtegy")
}

//  Конкретная стратегия

type lfu struct {
}

func (l *lfu) evict(c *cache) {
	fmt.Println("Evicting by lfu strtegy")
}

//  Контекст

type cache struct {
	storage      map[string]string
	evictionAlgo evictionAlgo
	capacity     int
	maxCapacity  int
}

func initCache(e evictionAlgo) *cache {
	storage := make(map[string]string)
	return &cache{
		storage:      storage,
		evictionAlgo: e,
		capacity:     0,
		maxCapacity:  2,
	}
}

func (c *cache) setEvictionAlgo(e evictionAlgo) {
	c.evictionAlgo = e
}

func (c *cache) add(key, value string) {
	if c.capacity == c.maxCapacity {
		c.evict()
	}
	c.capacity++
	c.storage[key] = value
}

func (c *cache) get(key string) {
	delete(c.storage, key)
}

func (c *cache) evict() {
	c.evictionAlgo.evict(c)
	c.capacity--
}

// func main() {
// 	lfu := &lfu{}
// 	cache := initCache(lfu)

// 	cache.add("a", "1")
// 	cache.add("b", "2")

// 	cache.add("c", "3")

// 	lru := &lru{}
// 	cache.setEvictionAlgo(lru)

// 	cache.add("d", "4")

// 	fifo := &fifo{}
// 	cache.setEvictionAlgo(fifo)

// 	cache.add("e", "5")

// }