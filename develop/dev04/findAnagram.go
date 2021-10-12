/*
Поиск анаграмм по словарю.

Напишите функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Входные данные для функции: ссылка на массив - каждый элемент которого -
слово на русском языке в кодировке utf8.

Выходные данные: Ссылка на мапу множеств анаграмм.

Ключ - первое встретившееся в словаре слово из множества
Значение - ссылка на массив, каждый элемент которого, слово из множества.

Массив должен быть отсортирован по возрастанию.

Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.
 */
package main

import (
	"sort"
	"strings"
)

func findAnagram(input []string) map[string][]string {

	// инициализация map
	groupAnagrams := make(map[string][]string)
	output := make(map[string][]string)

	// Приводим все слова к нижнему регистру
	for i, word := range input {
		input[i] = strings.ToLower(word)
	}

	// Итерируемся по списку слов и добавляем в groupAnagrams
	for _, currentWord := range input {
		sortedWord := SortString(currentWord)               //сортируем символы в слове
		if wordList, ok := groupAnagrams[sortedWord]; !ok { //если map не содежит ключ sortedWord, добавляем
			wordList = append(wordList)
		} else if contains(wordList, currentWord) { // в противном случае проверяем, есть ли значение currentWord уже по такому ключу sortedWord
			continue // если да - идем на след. итерацию
		}

		groupAnagrams[sortedWord] = append(groupAnagrams[sortedWord], currentWord) // добавлем по ключу значение
	}

	// Итерируемся по groupAnagrams
	for key, value := range groupAnagrams {
		if len(value) == 1 {
			delete(groupAnagrams, key) // удаляем значение, если слово в множестве, всего одно
		} else {
			output[value[0]] = value // добавляем в output значение с ключом, который является первым встретившемся в словаре слово из множества
			sort.Strings(value)      // сортируем значения по возрастанию
		}
	}

	return output
}

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func SortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}

func contains(wordList []string, word string) bool {
	for _, currentWord := range wordList {
		if currentWord == word {
			return true
		}
	}

	return false
}