package lasagna

/*
PreparationTime accepts a slice of layers as a []string and
the average preparation time per layer in minutes as an int.

The function should return the estimate for the total preparation time
based on the number of layers as an int.

If the average preparation time is passed as 0 (the default initial value for an int),
then the default value of 2 should be used.
*/
func PreparationTime(layers []string, avgPrepTime int) int {
	totalLayers := len(layers)
	if avgPrepTime == 0 {
		return totalLayers * 2
	}
	return totalLayers * avgPrepTime
}

/*
Quantities function takes a slice of layers as parameter as a []string.
The function will then determine the quantity of noodles and sauce needed to make your meal.

The result should be returned as two values of noodles as an int and sauce as a float64.

For each noodle layer in your lasagna, you will need 50 grams of noodles.
For each sauce layer in your lasagna, you will need 0.2 liters of sauce.
*/
func Quantities(layers []string) (int, float64) {
	const noodlesGrams int = 50
	const sauceLiters float64 = 0.2
	var noodlesLayer, sauceLayer int
	for _, layer := range layers {
		if layer == "noodles" {
			noodlesLayer++
		} else if layer == "sauce" {
			sauceLayer++
		}
	}
	return noodlesLayer * noodlesGrams, float64(sauceLayer) * sauceLiters
}

/*
AddSecretIngredient accepts two slices of ingredients of type []string as parameters.

The first parameter is the list your friend sent you, the second is the ingredient list of your own recipe.
The last element in your ingredient list is always "?".

The function should replace it with the last item from your friends list.

AddSecretIngredient does not return anything - you should modify the list of your ingredients directly.
The list with your friend's ingredients should not be modified.
*/
func AddSecretIngredient(friendIngredients, myIngredients []string) {
	myIngredients[len(myIngredients)-1] = friendIngredients[len(friendIngredients)-1]
}

/*
ScaleRecipe takes two parameters:

- A slice of float64 amounts needed for 2 portions.

- The number of portions you want to cook.

The amounts listed in your cookbook only yield enough lasagna for two portions.

Since you want to cook for more people next time,
you want to calculate the amounts for different numbers of portions.

The function should return a slice of float64 of the amounts needed for the desired number of portions.

You want to keep the original recipe though.
This means the amounts argument should not be modified in this function.
*/
func ScaleRecipe(amounts []float64, portions int) []float64 {
	newAmounts := make([]float64, len(amounts))
	for _, amount := range amounts {
		newAmounts = append(newAmounts, (amount/2)*float64(portions))
	}
	return newAmounts
}
