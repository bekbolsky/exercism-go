package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	grossUnit := map[string]int{
		"quarter_of_a_dozen": 3,
		"half_of_a_dozen":    6,
		"dozen":              12,
		"small_gross":        120,
		"gross":              144,
		"great_gross":        1728,
	}
	return grossUnit
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	billMap := make(map[string]int)
	return billMap
}

/*
AddItem adds an item to customer bill.

Returns false if the given unit is not in the units map.

If the item is already present in the bill,
increases its quantity by the amount that belongs to the provided unit.
Otherwise adds the item to the customer bill, indexed by the item name, then returns true.
*/
func AddItem(bill, units map[string]int, item, unit string) bool {
	unitValue, unitExists := units[unit]
	if unitExists {
		bill[item] += unitValue
		return true
	}
	return false
}

/*
RemoveItem removes an item from customer bill.

Returns false if the given item is not in the bill.
Returns false if the given unit is not in the units map.
Returns false if the new quantity would be less than 0.

If the new quantity is 0, completely removes the item from the bill then returns true.
Otherwise, reduce the quantity of the item and return true.
*/
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	itemValue, ItemExists := bill[item]
	unitValue, UnitExists := units[unit]
	if !ItemExists || !UnitExists {
		return false
	}
	if newQty := itemValue - unitValue; newQty == 0 {
		delete(bill, item)
	} else if newQty < 0 {
		return false
	} else {
		bill[item] = newQty
	}
	return true
}

/*
GetItem returns the quantity of an item that the customer has in his/her bill.

Returns 0 and false if the item is not in the bill.

Returns the quantity of the item in the bill and true.
*/
func GetItem(bill map[string]int, item string) (int, bool) {
	itemValue, itemExists := bill[item]
	return itemValue, itemExists
}
