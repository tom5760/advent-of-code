package main

import (
	"fmt"
	"strings"
)

const (
	chemStart = "FUEL"
	chemEnd   = "ORE"
)

// Ingredient represents a piece of a recipe.
type Ingredient struct {
	Name  string
	Count int
}

type Recipe struct {
	Inputs []Ingredient
	Output Ingredient
}

// ChemDB contains recipes for chemical formulas.
type ChemDB struct {
	Recipes map[string]Recipe
}

// Parse reads lines of recipes into a database.
func Parse(lines []string) (*ChemDB, error) {
	db := &ChemDB{
		Recipes: make(map[string]Recipe, len(lines)),
	}

	for _, line := range lines {
		parts := strings.Split(line, " => ")
		if len(parts) != 2 {
			return nil, fmt.Errorf("unexpected input/output format")
		}

		inputParts := strings.Split(parts[0], ", ")
		if len(inputParts) == 0 {
			return nil, fmt.Errorf("unexpected input format")
		}

		recipe := Recipe{
			Inputs: make([]Ingredient, len(inputParts)),
		}

		_, err := fmt.Sscanf(parts[1], "%d %s", &recipe.Output.Count, &recipe.Output.Name)
		if err != nil {
			return nil, fmt.Errorf("failed to scan output: %w", err)
		}

		for i := range recipe.Inputs {
			input := &recipe.Inputs[i]
			_, err := fmt.Sscanf(inputParts[i], "%d %s", &input.Count, &input.Name)
			if err != nil {
				return nil, fmt.Errorf("failed to scan input: %w", err)
			}
		}

		db.Recipes[recipe.Output.Name] = recipe
	}

	return db, nil
}

func (db *ChemDB) Totals(start string) int {
	work := map[string]Recipe{}

	cur := db.Recipes[start].Output

	for {
		recipe := db.Recipes[cur.Name]

		for _, ing := range recipe.Inputs {
			work[ingName].Count += (ing.Count * cur.Count)
		}

		for ing := range work {
			if ing.Name != chemEnd {
				cur = ing
				break
			}
		}

		if cur.Name == chemEnd {
			break
		}

		delete(work, cur)
	}

	return cur.Count
}
