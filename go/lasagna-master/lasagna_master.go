package lasagna

// TODO: define the 'PreparationTime()' function

// TODO: define the 'Quantities()' function

// TODO: define the 'AddSecretIngredient()' function

// TODO: define the 'ScaleRecipe()' function

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.

// PreparationTime() takes a slice of layers with the average preparation time per layer, then returns
// a total estimate time.
func PreparationTime(layers []string, timePerLayer int) (totalTime int) {
	if timePerLayer == 0 {
		timePerLayer = 2
	}
	return len(layers) * timePerLayer
}

// Quantities() takes a slice of layers and returns the quantity of noodles and sauce needed to make your meal.
// For each noodle layer in your lasagna, you will need 50 grams of noodles. For each sauce layer in your lasagna, you will need 0.2 liters of sauce.
func Quantities(layers []string) (noodles int, sauce float64) {
	for i := 0; i < len(layers); i++ {
		if layers[i] == "noodles" {
			noodles += 50
		} else if layers[i] == "sauce" {
			sauce += 0.2
		}
	}
	return noodles, sauce
}

// AddSecretIngredient() takes two slices of ingredients and replaces the last item of the second list with the last
// item of the first list.
func AddSecretIngredient(gordonsRecipe, myRecipe []string) {
	if len(myRecipe) > 0 && len(gordonsRecipe) > 0 {
		myRecipe[len(myRecipe)-1] = gordonsRecipe[len(gordonsRecipe)-1]
	}
}

// ScaleRecipe() takes a slice of ingredient quantities for 2 people with the number of portions you want to cook.
// It returns the updated slice of ingredient quantities according to the number of portions desired.
func ScaleRecipe(currentQuantities []float64, portions int) (updatedQuantities []float64) {
	if currentQuantities == nil {
		return nil
	}

	updatedQuantities = make([]float64, len(currentQuantities))
	var onePortion float64

	for i := 0; i < len(currentQuantities); i++ {
		onePortion = currentQuantities[i] / 2
		updatedQuantities[i] = onePortion * float64(portions)
	}
	return updatedQuantities
}
