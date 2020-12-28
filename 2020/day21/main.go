package main

import (
	"fmt"
	"regexp"
	"sort"
	"strings"

	"github.com/mbonnafon/AdventOfCode/helpers"
)

var r = regexp.MustCompile(`([a-z\s]+)\(contains ([a-z,\s]+)\)`)

type food struct {
	ingredients []string
	allergens   []string
}

func (f food) contains(a string) bool {
	for _, alr := range f.allergens {
		if a == alr {
			return true
		}
	}
	return false
}

func main() {
	lines, _ := helpers.StringLines("./input.txt")
	fmt.Println("Part 1. ingredients that cannot possibly contain any of the allergens in your list appears:", pt1(lines))
	fmt.Println("Part 2. your canonical dangerous ingredient list is:", pt2(lines))
}

func pt1(lines []string) int {
	var foods []food
	for _, l := range lines {
		foods = append(foods, parseFood(l))
	}
	count, _ := listNonAllergenIngredients(foods)
	return count
}

func pt2(lines []string) string {
	var foods []food
	for _, l := range lines {
		foods = append(foods, parseFood(l))
	}
	_, nonAllergen := listNonAllergenIngredients(foods)
	return findAllergenCorrespondance(foods, nonAllergen)
}

func parseFood(line string) food {
	line = strings.ReplaceAll(line, ",", "")
	l := r.FindStringSubmatch(line)
	l[1] = strings.TrimSpace(l[1])
	l[2] = strings.TrimSpace(l[2])
	return food{ingredients: strings.Split(l[1], " "), allergens: strings.Split(l[2], " ")}
}

func listNonAllergenIngredients(foods []food) (int, map[string]bool) {
	// count how many times an ingrendient can be an allergen
	ingredients := map[string]map[string]int{}
	for _, food := range foods {
		for _, ing := range food.ingredients {
			for _, all := range food.allergens {
				if ingredients[ing] == nil {
					ingredients[ing] = make(map[string]int)
				}
				ingredients[ing][all]++
			}
		}
	}

	// count allergens occurences
	allergensOccurences := map[string]int{}
	for _, food := range foods {
		for _, all := range food.allergens {
			allergensOccurences[all]++
		}
	}
	// consider an ingredient if it's present every time an allergen is listed
	ingAll := map[string]string{}
	for ing, allergen := range ingredients {
		for all, occ := range allergen {
			if allergensOccurences[all] == occ {
				ingAll[ing] = all
			}
		}
	}

	// count all non allergens occurence
	var count int
	nonAllergens := make(map[string]bool)
	for _, food := range foods {
		for _, ing := range food.ingredients {
			if ingAll[ing] == "" {
				nonAllergens[ing] = true
				count++
			}
		}
	}

	return count, nonAllergens
}

func findAllergenCorrespondance(foods []food, nonAllergen map[string]bool) string {
	// count how many times an allergen in listed per aliments
	allergen := map[string]map[string]int{}
	for _, food := range foods {
		for _, all := range food.allergens {
			for _, ing := range food.ingredients {
				if nonAllergen[ing] == true {
					continue
				}
				if allergen[all] == nil {
					allergen[all] = make(map[string]int)
				}
				allergen[all][ing]++
			}
		}
	}

	correspondance := make(map[string]string)
	alreadyFound := make(map[string]bool)
	for len(allergen) > 0 {
		// try to find correspondance
		for all, ing := range allergen {
			if alreadyFound[all] {
				continue
			}
			if s := findMax(ing); s != "" {
				correspondance[all] = s
				delete(allergen, all)
			}
		}
		// remove already found
		for _, ing := range allergen {
			for _, found := range correspondance {
				delete(ing, found)
			}
		}
	}

	// sort them
	var allergens []string
	for s := range correspondance {
		allergens = append(allergens, s)
	}
	sort.Strings(allergens)

	var canonDangerIngr string
	for _, v := range allergens {
		canonDangerIngr += correspondance[v] + ","
	}
	return strings.TrimSuffix(canonDangerIngr, ",")
}

func findMax(m map[string]int) string {
	var max string
	var maxInt, previous int
	for k, v := range m {
		if len(m) == 1 {
			return k
		}
		if v == maxInt {
			previous = maxInt
		}
		if v > maxInt {
			previous = maxInt
			maxInt = v
			max = k
		}
	}
	if maxInt == previous {
		return ""
	}
	return max
}
