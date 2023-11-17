package core

import "sort"

type CatDataWrapper struct {
	Data []Cat `json:"data"`
}

type Cat struct {
	Breed   string `json:"breed"`
	Country string `json:"country"`
	Origin  string `json:"origin"`
	Coat    string `json:"coat"`
	Pattern string `json:"pattern"`
}

type GroupedCats map[string][]Cat

func GroupCatsByCountry(cats []Cat) GroupedCats {
	groupedCats := make(GroupedCats)
	for _, cat := range cats {
		groupedCats[cat.Country] = append(groupedCats[cat.Country], cat)
	}

	return groupedCats
}

func SortCatsByBreedName(groupedCats GroupedCats) {
	for _, countryCats := range groupedCats {
		sort.Slice(
			countryCats,
			func(i, j int) bool {
				return len(countryCats[i].Breed) < len(countryCats[j].Breed)
			},
		)
	}
}
