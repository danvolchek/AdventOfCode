package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path"
	"regexp"
	"sort"
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

	// not all allergens we care about have assignments yet, so we can't tell if it's invalid
	// TODO: this can be expanded: e.g. if 1 is left to be assigned and we have 2 missing, it is invalid
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

		// prepare for future assignment
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

		// only keep going if we:
		//  - picked an assignment that's possible according to one of the rules
		//  - didn't invalidate any constraints with this assignment
		if len(unassignments[allergenToAssign]) != 0 {
			if !constraintsInvalid(constraints, assignments) {
				if assign(allIngredients, allAllergens[1:], assignments, unassignments, constraints) {
					return true
				}
			}

		}

		// unsuccessful, so undo
		delete(assignments, allergenToAssign)
		delete(unassignments, allergenToAssign)
		allIngredients[ingredient] = true

	}

	return false
}

func solve(r io.Reader) {
	scanner := bufio.NewScanner(r)

	var constriants []constraint
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

		constriants = append(constriants, constraint{
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
	if !assign(allIngredients, toSlice(allAllergens), assignments, make(map[string]map[string]bool), constriants) {
		panic("impossible")
	}

	fmt.Println(assignments)

	sortedAllergens := toSlice(allAllergens)
	sort.Strings(sortedAllergens)

	var res strings.Builder
	for i, allergen := range sortedAllergens {
		res.WriteString(assignments[allergen])
		if i != len(sortedAllergens)-1 {
			res.WriteString(",")
		}
	}

	fmt.Println(res.String())
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
