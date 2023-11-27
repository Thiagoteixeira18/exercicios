package exercicios

import "fmt"

func FavoriteCards() []int {
	return []int{2, 6, 9}
}

func GetItem(cards []int, index int) int {
	if index < 0 || index >= len(cards) {
		return -1
	}
	return cards[index]
}

func SetItem(cards []int, index, newCard int) []int {
	if index < 0 || index >= len(cards) {
		cards = append(cards, newCard)
	} else {
		cards[index] = newCard
	}
	return cards
}

func PrependItems(cards []int, values ...int) []int {
	if len(values) == 0 {
		return cards
	}
	return append(values, cards...)
}

func RemoveItem(cards []int, index int) []int {
	if index < 0 || index >= len(cards) {
		return cards
	}
	return append(cards[:index], cards[index+1:]...)
}

func main3() {
	cards := FavoriteCards()
	fmt.Println(cards)

	card := GetItem([]int{1, 2, 4, 1}, 2)
	fmt.Println(card)

	index := 2
	newCard := 6
	cards = SetItem([]int{1, 2, 4, 1}, index, newCard)
	fmt.Println(cards)

	slice := []int{3, 2, 6, 4, 8}
	cards = PrependItems(slice, 5, 1)
	fmt.Println(cards)

	cards = RemoveItem([]int{3, 2, 6, 4, 8}, 2)
	fmt.Println(cards)
}