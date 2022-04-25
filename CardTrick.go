package cards

// FavoriteCards returns a slice with the cards 2, 6 and 9 in that order.
func FavoriteCards() []int {
    var cards []int;

    cards = append(cards, 2, 6, 9);
    return (cards);
}

// GetItem retrieves an item from a slice at given position.
// If the index is out of range, we want it to return -1.
func GetItem(slice []int, index int) int {
	var value int;

    if index < 0 || index > len(slice) - 1 {
        return (-1);
    }
    value = slice[index];
    return (value);
}

// SetItem writes an item to a slice at given position overwriting an existing value.
// If the index is out of range the value needs to be appended.
func SetItem(slice []int, index, value int) []int {
    if index < 0 || index > len(slice) -1 {
        slice = append(slice, value);
        return (slice);
    }
	slice[index] = value;
    return (slice);
}

// PrependItems adds an arbitrary number of values at the front of a slice.
func PrependItems(slice []int, value ...int) []int {
	/*
    for _, v := range value {
        slice = append([]int{v}, slice...);
    }
	*/
    slice = append(value, slice...)
	return (slice);
}

// RemoveItem removes an item from a slice by modifying the existing slice.
func RemoveItem(slice []int, index int) []int {
	var newSlice []int;

    for i, v := range slice {
        if i != index {
            newSlice = append(newSlice, v);
        }
    }
	return (newSlice);
}

