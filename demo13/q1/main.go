package main

import "fmt"

type AnimalCategory struct {
	kingdom string
	phylum  string
	class   string
	order   string
	family  string
	genus   string
	species string
}

type Animal struct {
	scientificName string
	AnimalCategory
}

type Cat struct {
	name string
	Animal
}

func (cat Cat) String() string {
	return fmt.Sprintf("%s (category: %s, name: %q)", cat.scientificName, cat.Animal.AnimalCategory, cat.name)
}

func (ac AnimalCategory) string() string {
	return fmt.Sprintf("%s%s%s%s%s%s%s", ac.kingdom, ac.phylum, ac.class, ac.order, ac.family, ac.genus, ac.species)
}

func (a Animal) string() string {
	return fmt.Sprintf("%s (category: %s)", a.scientificName, a.AnimalCategory)
}

func (a Animal) Category() string {
	return a.AnimalCategory.string()
}

func main() {
	category := AnimalCategory{species: "cat"}
	fmt.Printf("The animal category: %s\n", category)

	animal := Animal{
		scientificName: "American Shorthair",
		AnimalCategory: category,
	}
	fmt.Printf("The animal: %s\n", animal)
}
