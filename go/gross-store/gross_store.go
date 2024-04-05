package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	measurements := map[string]int{
		"quarter_of_a_dozen": 3,
		"half_of_a_dozen":    6,
		"dozen":              12,
		"small_gross":        120,
		"gross":              144,
		"great_gross":        1728,
	}
	return measurements
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return make(map[string]int)
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	pieces, exists := units[unit]
	if !exists {
		return false
	}
	bill[item] += pieces
	return true

}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, itemName, unitName string) bool {
	requiredPieces, unitExists := units[unitName]
	billPieces, itemExists := bill[itemName]

	if !unitExists || !itemExists || requiredPieces > billPieces {
		return false
	} else if requiredPieces == billPieces {
		delete(bill, itemName)
	} else {
		bill[itemName] -= requiredPieces
	}

	return true
}

// GetItem returns the quantity of an item that the customer has in their bill
func GetItem(bill map[string]int, itemName string) (int, bool) {
	quantity, itemExists := bill[itemName]
	return quantity, itemExists
}
