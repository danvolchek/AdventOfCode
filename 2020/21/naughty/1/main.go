package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path"
	"regexp"
	"strings"
)

func input() *os.File {
	input, err := os.Open(path.Join("2020", "21", "input.txt"))
	if err != nil {
		panic(err)
	}

	return input
}

var foodRegex = regexp.MustCompile(`^(.+) \(contains (.+)\)$`)

type constraint struct {
	ingredients map[string]bool
	allergens   map[string]bool
}

func (c constraint) invalid(assignments map[string]string) bool {
	applicableAllergens := make(map[string]bool)

	applicableIngredients := make(map[string]bool)
	for allergen, ingredient := range assignments {
		if !c.allergens[allergen] {
			continue
		}

		applicableAllergens[allergen] = true

		if c.ingredients[ingredient] {
			applicableIngredients[ingredient] = true
		}
	}

	// don't know yet
	if len(applicableAllergens) != len(c.allergens) {
		return false
	}

	// all allergens we want are assigned, yet our ingredient list doesn't contain all of them
	if len(applicableIngredients) < len(c.allergens) {
		return true
	}

	return false
}

func constraintsInvalid(constraints []constraint, assignments map[string]string) bool {
	for _, constraint := range constraints {
		if constraint.invalid(assignments) {
			return true
		}
	}

	return false
}

func assign(allIngredients map[string]bool, allAllergens []string, assignments map[string]string, unassignments map[string]map[string]bool, constraints []constraint) bool {
	if len(allAllergens) == 0 {
		return true
	}

	allergenToAssign := allAllergens[0]

	toTry := toSlice(allIngredients)

	for _, ingredient := range toTry {

		if unassignments[allergenToAssign][ingredient] {
			continue
		}

		assignments[allergenToAssign] = ingredient
		delete(allIngredients, ingredient)
		unassignments[allergenToAssign] = make(map[string]bool)

		for _, c := range constraints {
			if !c.allergens[allergenToAssign] || !c.ingredients[ingredient] {
				continue
			}

			for i := range c.ingredients {
				if i == ingredient {
					continue
				}

				unassignments[allergenToAssign][i] = true
			}
		}

		if len(unassignments[allergenToAssign]) != 0 {

			if !constraintsInvalid(constraints, assignments) {
				if assign(allIngredients, allAllergens[1:], assignments, unassignments, constraints) {
					return true
				}
			}

		}

		delete(assignments, allergenToAssign)
		delete(unassignments, allergenToAssign)
		allIngredients[ingredient] = true

	}

	return false
}

func solve(r io.Reader) {
	scanner := bufio.NewScanner(r)

	var constraints []constraint
	allIngredients := make(map[string]bool)
	allAllergens := make(map[string]bool)

	for scanner.Scan() {
		row := scanner.Text()

		parts := foodRegex.FindStringSubmatch(row)

		ingredients := strings.Split(parts[1], " ")
		allergens := strings.Split(parts[2], ", ")

		allergenMap := make(map[string]bool)
		for _, allergen := range allergens {
			allergenMap[allergen] = true

			allAllergens[allergen] = true
		}

		ingredientsMap := make(map[string]bool)
		for _, ingredient := range ingredients {
			ingredientsMap[ingredient] = true

			allIngredients[ingredient] = true
		}

		constraints = append(constraints, constraint{
			ingredients: ingredientsMap,
			allergens:   allergenMap,
		})
	}

	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	fmt.Println("ingredients", toSlice(allIngredients))
	fmt.Println("allergens", toSlice(allAllergens))

	assignments := make(map[string]string)
	if !assign(allIngredients, toSlice(allAllergens), assignments, make(map[string]map[string]bool), constraints) {
		panic("impossible")
	}

	fmt.Println(assignments)

	ingredientsWithAllergens := make(map[string]bool)

	for _, ingredient := range assignments {
		ingredientsWithAllergens[ingredient] = true
	}

	ingredientsWithoutAllergens := make(map[string]bool)

	for ingredient := range allIngredients {
		if !ingredientsWithAllergens[ingredient] {
			ingredientsWithoutAllergens[ingredient] = true
		}
	}

	fmt.Println(ingredientsWithoutAllergens)

	total := 0
	for _, constraint := range constraints {
		for ingredient := range ingredientsWithoutAllergens {
			if constraint.ingredients[ingredient] {
				total += 1
			}
		}
	}

	fmt.Println(total)
}

func toSlice(m map[string]bool) []string {
	ret := make([]string, len(m))

	i := 0

	for k := range m {
		ret[i] = k
		i += 1
	}

	return ret
}

func main() {
	solve(strings.NewReader("mxmxvkd kfcds sqjhc nhms (contains dairy, fish)\ntrh fvjkl sbzzf mxmxvkd (contains dairy)\nsqjhc fvjkl (contains soy)\nsqjhc mxmxvkd sbzzf (contains fish)"))

	solve(input())
}
