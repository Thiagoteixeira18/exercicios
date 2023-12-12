package exercicios

// Units armazena as medidas de unidades da Gross Store.
func Units() map[string]int {
	return map[string]int{
		"quarter_of_a_dozen": 3,
		"half_of_a_dozen":    6,
		"dozen":              12,
		"small_gross":        120,
		"gross":              144,
		"great_gross":        1728,
	}
}

// NewBill cria uma nova fatura (bill).
func NewBill() map[string]int {
	return make(map[string]int)
}

// AddItem adiciona um item à fatura do cliente.
func AddItem(bill, units map[string]int, item, unit string) bool {
	unitValue, exists := units[unit]
	if !exists {
		return false
	}

	if currentQuantity, exists := bill[item]; exists {
		bill[item] = currentQuantity + unitValue
	} else {
		bill[item] = unitValue
	}
	return true
}

// RemoveItem remove um item da fatura do cliente.
func RemoveItem1(bill, units map[string]int, item, unit string) bool {
	unitValue, exists := units[unit]
	if !exists {
		return false
	}

	currentQuantity, exists := bill[item]
	if !exists {
		return false
	}

	newQuantity := currentQuantity - unitValue
	if newQuantity < 0 {
		return false
	} else if newQuantity == 0 {
		delete(bill, item)
	} else {
		bill[item] = newQuantity
	}
	return true
}

// GetItem retorna a quantidade de um item que o cliente possui em sua fatura.
func GetItem1(bill map[string]int, item string) (int, bool) {
	quantity, exists := bill[item]
	return quantity, exists
}