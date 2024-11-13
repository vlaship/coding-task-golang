package problem2115

import "container/list"

func findAllRecipes(recipes []string, ingredients [][]string, supplies []string) []string {
	//create map ingredient to recipe
	ingredientToRecipe := make(map[string][]string)
	for i, recipe := range recipes {
		for _, ingredient := range ingredients[i] {
			ingredientToRecipe[ingredient] = append(ingredientToRecipe[ingredient], recipe)
		}
	}

	//create map recipe to ingredients
	recipeToIngredients := make(map[string][]string)
	for i, recipe := range recipes {
		recipeToIngredients[recipe] = ingredients[i]
	}

	//create map recipe to degree
	recipeToDegree := make(map[string]int)
	for r, i := range recipeToIngredients {
		recipeToDegree[r] = len(i)
	}

	//put supplies to deque
	deque := list.New()
	for _, supply := range supplies {
		deque.PushBack(supply)
	}

	result := []string{}

	//bfs
	for deque.Len() > 0 {
		s := deque.Front()
		deque.Remove(s)
		for _, r := range ingredientToRecipe[s.Value.(string)] {
			recipeToDegree[r]--
			if recipeToDegree[r] == 0 {
				deque.PushBack(r)
				result = append(result, r)
			}
		}
	}

	return result
}
