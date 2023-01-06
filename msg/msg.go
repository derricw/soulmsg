/*
Copyright © 2023 Derric Williams <balls@balls.balls>

*/
package msg

import (
	"fmt"
	"math/rand"
	"strings"
)

func RandomTemplate() string {
	return TEMPLATES[rand.Intn(len(TEMPLATES))]
}

func RandomCategoryName() string {
	return CATEGORY_NAMES[rand.Intn(len(CATEGORY_NAMES))]
}

func RandomCategory() []string {
	name := RandomCategoryName()
	return CATEGORIES[name]
}

func RandomThing() string {
	category := RandomCategory()
	return category[rand.Intn(len(category))]
}

func RandomConjunction() string {
	return CONJUNCTIONS[rand.Intn(len(CONJUNCTIONS))]
}

func RandomClause() string {
	template := RandomTemplate()
	thing := RandomThing()
	return strings.Replace(template, "****", thing, 1)
}

func RandomMessage(allowConjuctions bool) string {
	clause0 := RandomClause()
	if !allowConjuctions {
		return clause0
	}
	conjuction := RandomConjunction()
	clause1 := RandomClause()
	return fmt.Sprintf("%s %s %s", clause0, conjuction, clause1)
}
