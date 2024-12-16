package problem_2115

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindAllRecipes(t *testing.T) {
	tests := []struct {
		name        string
		recipes     []string
		ingredients [][]string
		supplies    []string
		expected    []string
	}{
		{
			name:        "Example 1",
			recipes:     []string{"bread"},
			ingredients: [][]string{{"yeast", "flour"}},
			supplies:    []string{"yeast", "flour", "corn"},
			expected:    []string{"bread"},
		},
		{
			name:        "Example 2",
			recipes:     []string{"bread", "sandwich"},
			ingredients: [][]string{{"yeast", "flour"}, {"bread", "meat"}},
			supplies:    []string{"yeast", "flour", "meat"},
			expected:    []string{"bread", "sandwich"},
		},
		{
			name:        "Example 3",
			recipes:     []string{"bread", "sandwich", "burger"},
			ingredients: [][]string{{"yeast", "flour"}, {"bread", "meat"}, {"sandwich", "meat", "bread"}},
			supplies:    []string{"yeast", "flour", "meat"},
			expected:    []string{"bread", "sandwich", "burger"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findAllRecipes(tt.recipes, tt.ingredients, tt.supplies)
			assert.ElementsMatch(t, tt.expected, result)
		})
	}
}
