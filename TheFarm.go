package thefarm

import "errors"
import "fmt"

// See types.go for the types defined for this exercise.

// TODO: Define the SillyNephewError type here.
type SillyNephewError interface {
    Error() string
}

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {
    var fodder float64;
    
    amount, err := weightFodder.FodderAmount();
    if err != nil && amount >= 0 {
		amount = amount * 2;
    }
	if err != nil && err != ErrScaleMalfunction {
        return 0, errors.New("non-scale error");
    } 
	if amount < 0 {
        return 0, errors.New("negative fodder");
    }
	if cows == 0 {
        return 0, errors.New("division by zero");
    }
	if cows < 0 {
        e := fmt.Sprintf("silly nephew, there cannot be %d cows", cows);
        return 0, errors.New(e);
;    }
	fodder = amount / float64(cows); 
    return fodder, nil;
}

