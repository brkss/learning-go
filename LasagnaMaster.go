package lasagna

import "fmt"

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, minutes int) int {
    var prepTime int;

    if(minutes == 0){
        prepTime = len(layers) * 2;
    }else {
    	prepTime = len(layers) * minutes;
    }
    return (prepTime);
}

// TODO: define the 'Quantities()' function

func Quantities(layers []string) (int, float64) {
    var noodels int;
    var water float64;

    for i := 0; i < len(layers); i++ {
        if layers[i] == "noodles" {
            noodels += 50;
        } else if layers[i] == "sauce" {
        	water += 0.2;
        }
    }
	return noodels, water; 
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendIngredients, myIngredrients []string) {
    myIngredrients[len(myIngredrients) - 1] = friendIngredients[len(friendIngredients) - 1];
}
// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(portions []float64, portionWanted int) []float64 {
	var scaler float64;
    var val float64;
    var scaledPortions []float64;

    scaler = float64(portionWanted) / 2.0;
	for i := 0; i < len(portions); i++ {
        val = portions[i] * scaler;
        fmt.Printf("val : %f \nportion : %f\n scaler: %f", val, portions[i], scaler);
        scaledPortions = append(scaledPortions, val);
    }
	return (scaledPortions);
}

