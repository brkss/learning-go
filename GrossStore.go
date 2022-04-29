package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
    var units =  make(map[string]int);

    units["quarter_of_a_dozen"] = 3;
    units["half_of_a_dozen"] = 6;
    units["dozen"] = 12;
    units["small_gross"] = 120;
    units["gross"] = 144;
    units["great_gross"] = 1728;
	return (units);
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	var bill = make(map[string]int);
    
    return bill;
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {

    un, unExist := units[unit];
    if !unExist {
        return (false);
    }
	score, itExist := bill[item];
    if itExist {
        bill[item] = score + un;
    } else {
    	bill[item] = un
    }
	return true;    
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	it, itExist := bill[item];
    un, unExist := units[unit];
    if !itExist || !unExist {
        return false;
    }
	newQuantity := it - un;
    if newQuantity < 0 {
        return false;
    } else if newQuantity == 0 {
    	delete(bill, item);
        return true;
    }
	bill[item] -= un;
    return true;
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	it, exist := bill[item];
	return it, exist;
}

